package service

import (
	"context"
	"encoding/json"
	"log/slog"
	"math"
	"strings"
	"time"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/geocoder"
	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/routing"
)

type MatchEvaluation struct {
	Compatibility  float64
	DetourKM       float64
	AdditionalMins int
	UsableKM       float64
	Utilization    float64
}

func EvaluateMatch(candidate domain.MatchCandidate, baselineKM, composedKM, additionalMinutes, usableKM float64) (MatchEvaluation, bool) {
	switch candidate.TripStatus {
	case "PROGRAMADO", "ACTIVO", "EN_CURSO":
	default:
		return MatchEvaluation{}, false
	}
	if candidate.LoadStatus != "PENDIENTE" {
		return MatchEvaluation{}, false
	}
	if candidate.LoadWeight <= 0 || candidate.LoadVolume <= 0 ||
		candidate.TripWeight < candidate.LoadWeight || candidate.TripVolume < candidate.LoadVolume {
		return MatchEvaluation{}, false
	}
	if !vehiclesCompatible(candidate.VehicleName, candidate.LoadRequiredVehicle) {
		return MatchEvaluation{}, false
	}
	dateGap := calendarDayGap(candidate.TripDeparture, candidate.LoadPickup)
	if dateGap > 2 {
		return MatchEvaluation{}, false
	}
	if baselineKM <= 0 || usableKM < 0 {
		return MatchEvaluation{}, false
	}
	detour := math.Max(0, composedKM-baselineKM)
	if detour > 300 || detour > baselineKM*0.60 {
		return MatchEvaluation{}, false
	}
	weightRatio := candidate.LoadWeight / candidate.TripWeight
	volumeRatio := candidate.LoadVolume / candidate.TripVolume
	utilization := math.Max(weightRatio, volumeRatio) * 100
	capacityScore := capacityPoints(utilization / 100)
	dateScore := 10.0
	switch dateGap {
	case 1:
		dateScore = 6
	case 2:
		dateScore = 3
	}
	routeScore := math.Max(0, 55*(1-(detour/math.Max(usableKM, 1)/0.50)))
	score := math.Max(0, math.Min(99, routeScore+capacityScore+dateScore+10))
	return MatchEvaluation{
		Compatibility:  score,
		DetourKM:       detour,
		AdditionalMins: int(math.Max(0, math.Round(additionalMinutes))),
		UsableKM:       usableKM,
		Utilization:    utilization,
	}, true
}

func capacityPoints(ratio float64) float64 {
	if ratio < 0.4 {
		return 20 * ratio / 0.4
	}
	if ratio >= 0.9 {
		return 25
	}
	return 20 + (ratio-0.4)/0.5*5
}

func vehiclesCompatible(vehicle, required string) bool {
	required = normalizeMatchText(required)
	if required == "" || required == "cualquiera" || required == "indistinto" {
		return true
	}
	return normalizeMatchText(vehicle) == required
}

func normalizeMatchText(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	replacer := strings.NewReplacer(
		"á", "a", "é", "e", "í", "i", "ó", "o", "ú", "u", "ü", "u", "ñ", "n",
	)
	return replacer.Replace(value)
}

func calendarDayGap(a, b time.Time) int {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	da := time.Date(ay, am, ad, 0, 0, 0, 0, time.UTC)
	db := time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	gap := int(math.Abs(da.Sub(db).Hours() / 24))
	return gap
}

func (s *Service) RecalculateMatches(ctx context.Context, identity domain.Identity) (int, error) {
	if identity.Role != domain.RoleAdmin && identity.Role != domain.RoleOperator {
		return 0, ErrForbidden
	}
	candidates, err := s.repository.MatchCandidates(ctx, identity.CompanyID)
	if err != nil {
		return 0, err
	}
	if err := s.repository.DeleteProposedMatches(ctx, identity.CompanyID); err != nil {
		return 0, err
	}
	baselineRoutes := make(map[string]routing.Route)
	cargoRoutes := make(map[string]routing.Route)
	stored := 0
	for _, candidate := range candidates {
		tripOrigin, okOrigin := candidatePoint(candidate.TripOrigin, candidate.TripOriginLat, candidate.TripOriginLng)
		tripDestination, okTripDestination := candidatePoint(candidate.TripDestination, candidate.TripDestinationLat, candidate.TripDestinationLng)
		loadOrigin, okLoadOrigin := candidatePoint(candidate.LoadOrigin, candidate.LoadOriginLat, candidate.LoadOriginLng)
		loadDestination, okLoadDestination := candidatePoint(candidate.LoadDestination, candidate.LoadDestinationLat, candidate.LoadDestinationLng)
		if !okOrigin || !okTripDestination || !okLoadOrigin || !okLoadDestination {
			slog.Warn("skipping match candidate with unresolved location",
				"trip_id", candidate.TripID, "load_id", candidate.LoadID)
			continue
		}
		baseline, exists := baselineRoutes[candidate.TripID]
		if !exists {
			baseline = s.router.Route(ctx, []routing.Point{tripOrigin, tripDestination})
			baselineRoutes[candidate.TripID] = baseline
			if err := s.repository.SetTripBaseDistance(ctx, candidate.TripID, baseline.DistanceKM); err != nil {
				slog.Error("failed to store trip base distance",
					"trip_id", candidate.TripID, "error", err)
			}
		}
		cargo, exists := cargoRoutes[candidate.LoadID]
		if !exists {
			cargo = s.router.Route(ctx, []routing.Point{loadOrigin, loadDestination})
			cargoRoutes[candidate.LoadID] = cargo
		}
		composed := s.router.Route(ctx, []routing.Point{tripOrigin, loadOrigin, loadDestination, tripDestination})
		usable := cargo.DistanceKM
		evaluation, ok := EvaluateMatch(
			candidate, baseline.DistanceKM, composed.DistanceKM,
			(composed.Duration - baseline.Duration).Minutes(), usable,
		)
		if !ok || evaluation.Compatibility < 50 {
			continue
		}
		waypoints := []domain.MatchRouteWaypoint{
			routeWaypoint("origen_viaje", "Origen del viaje", tripOrigin),
			routeWaypoint("origen_carga", "Origen de la carga", loadOrigin),
			routeWaypoint("destino_carga", "Destino de la carga", loadDestination),
			routeWaypoint("destino_viaje", "Fin del viaje", tripDestination),
		}
		routeGeometry := struct {
			Coordinates [][]float64                 `json:"coordinates"`
			Waypoints   []domain.MatchRouteWaypoint `json:"waypoints"`
			Source      string                      `json:"source"`
			Planned     [][]float64                 `json:"planned"`
			Utilized    [][]float64                 `json:"utilized"`
			Detour      [][]float64                 `json:"detour"`
		}{
			Coordinates: composed.Coordinates, Waypoints: waypoints,
			Source:  composed.Source,
			Planned: baseline.Coordinates, Utilized: composed.Coordinates, Detour: composed.Coordinates,
		}
		encoded, err := json.Marshal(routeGeometry)
		if err != nil {
			return stored, err
		}
		_, err = s.repository.UpsertMatch(ctx, domain.MatchRecord{
			TripID: candidate.TripID, LoadID: candidate.LoadID,
			Compatibility: evaluation.Compatibility, DetourKM: evaluation.DetourKM,
			AdditionalMins: evaluation.AdditionalMins, UsableKM: evaluation.UsableKM,
			Utilization: evaluation.Utilization, GeometryJSON: encoded,
		})
		if err != nil {
			return stored, err
		}
		stored++
	}
	return stored, nil
}

func candidatePoint(city string, lat, lng *float64) (routing.Point, bool) {
	if lat != nil && lng != nil {
		return routing.Point{Latitude: *lat, Longitude: *lng}, true
	}
	location, ok := geocoder.Resolve(city)
	if !ok {
		return routing.Point{}, false
	}
	return routing.Point{Latitude: location.Latitude, Longitude: location.Longitude}, true
}

func routeWaypoint(name, label string, point routing.Point) domain.MatchRouteWaypoint {
	return domain.MatchRouteWaypoint{
		Name:       name,
		Label:      label,
		Coordinate: []float64{point.Longitude, point.Latitude},
	}
}

func (s *Service) ListMatches(ctx context.Context, identity domain.Identity) ([]domain.MatchItem, error) {
	if identity.Role != domain.RoleAdmin && identity.Role != domain.RoleOperator {
		return nil, ErrForbidden
	}
	items, err := s.repository.ListMatches(ctx, identity.CompanyID)
	return items, translateRepositoryError(err)
}

func (s *Service) GetMatch(ctx context.Context, identity domain.Identity, matchID string) (domain.MatchItem, error) {
	if identity.Role != domain.RoleAdmin && identity.Role != domain.RoleOperator {
		return domain.MatchItem{}, ErrForbidden
	}
	item, err := s.repository.GetMatch(ctx, identity.CompanyID, matchID)
	return item, translateRepositoryError(err)
}

func (s *Service) ListTripMatches(ctx context.Context, identity domain.Identity, tripID string) ([]domain.MatchItem, error) {
	if identity.Role != domain.RoleAdmin && identity.Role != domain.RoleOperator {
		return nil, ErrForbidden
	}
	items, err := s.repository.ListTripMatches(ctx, identity.CompanyID, tripID)
	return items, translateRepositoryError(err)
}
