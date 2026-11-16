package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/geocoder"
)

func (r *Postgres) MatchCandidates(ctx context.Context, companyID string) ([]domain.MatchCandidate, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT v.id_viaje, v.codigo_viaje, v.origen, v.destino,
		       v.fecha_salida, v.fecha_llegada_estimada, v.capacidad_disponible,
		       v.volumen_disponible, v.estado, tv.nombre,
		       concat(ve.patente, ' · ', tv.nombre),
		       v.lat_origen, v.lng_origen, v.lat_destino, v.lng_destino,
		       c.id_carga, c.codigo_carga, c.origen, c.destino, c.fecha,
		       c.peso, c.volumen, c.tipo_carga, c.vehiculo_requerido,
		       c.restricciones, c.estado,
		       c.lat_origen, c.lng_origen, c.lat_destino, c.lng_destino
		FROM viaje v
		JOIN usuario operador ON operador.id_usuario = v.id_usuario_operador
		JOIN vehiculo ve ON ve.id_vehiculo = v.id_vehiculo
		JOIN tipo_vehiculo tv ON tv.id_tipo_vehiculo = ve.id_tipo_vehiculo
		JOIN carga c ON c.estado = 'PENDIENTE'
		JOIN usuario solicitante ON solicitante.id_usuario = c.id_usuario
		WHERE operador.id_empresa = $1
		  AND solicitante.id_empresa = $1
		  AND v.estado IN ('PROGRAMADO', 'ACTIVO', 'EN_CURSO')
		ORDER BY v.fecha_salida, c.fecha
	`, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := make([]domain.MatchCandidate, 0)
	for rows.Next() {
		var item domain.MatchCandidate
		if err := rows.Scan(
			&item.TripID, &item.TripCode, &item.TripOrigin, &item.TripDestination,
			&item.TripDeparture, &item.TripArrival, &item.TripWeight, &item.TripVolume,
			&item.TripStatus, &item.VehicleName, &item.VehicleLabel,
			&item.TripOriginLat, &item.TripOriginLng, &item.TripDestinationLat, &item.TripDestinationLng,
			&item.LoadID, &item.LoadCode, &item.LoadOrigin, &item.LoadDestination, &item.LoadPickup,
			&item.LoadWeight, &item.LoadVolume, &item.LoadCargoType, &item.LoadRequiredVehicle,
			&item.LoadRestrictions, &item.LoadStatus,
			&item.LoadOriginLat, &item.LoadOriginLng, &item.LoadDestinationLat, &item.LoadDestinationLng,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *Postgres) UpsertMatch(ctx context.Context, item domain.MatchRecord) (string, error) {
	var id string
	err := r.pool.QueryRow(ctx, `
		INSERT INTO match_logistico (
			id_viaje, id_carga, compatibilidad, desvio_km, tiempo_adicional,
			km_aprovechables, capacidad_utilizada, route_geometry
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8::jsonb)
		ON CONFLICT (id_viaje, id_carga) DO UPDATE SET
			compatibilidad = EXCLUDED.compatibilidad,
			desvio_km = EXCLUDED.desvio_km,
			tiempo_adicional = EXCLUDED.tiempo_adicional,
			km_aprovechables = EXCLUDED.km_aprovechables,
			capacidad_utilizada = EXCLUDED.capacidad_utilizada,
			route_geometry = EXCLUDED.route_geometry,
			estado = 'PROPUESTO',
			fecha_creacion = now()
		RETURNING id_match
	`, item.TripID, item.LoadID, item.Compatibility, item.DetourKM, item.AdditionalMins,
		item.UsableKM, item.Utilization, string(item.GeometryJSON)).Scan(&id)
	return id, normalizeError(err)
}

func (r *Postgres) DeleteProposedMatches(ctx context.Context, companyID string) error {
	_, err := r.pool.Exec(ctx, `
		DELETE FROM match_logistico m
		USING viaje v, usuario operador, carga c, usuario solicitante
		WHERE m.id_viaje = v.id_viaje
		  AND operador.id_usuario = v.id_usuario_operador
		  AND m.id_carga = c.id_carga
		  AND solicitante.id_usuario = c.id_usuario
		  AND operador.id_empresa = $1
		  AND solicitante.id_empresa = $1
		  AND m.estado = 'PROPUESTO'
		  AND NOT EXISTS (
		      SELECT 1 FROM oferta o WHERE o.id_match = m.id_match
		  )
	`, companyID)
	return err
}

func (r *Postgres) SetTripBaseDistance(ctx context.Context, tripID string, distance float64) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE viaje SET distance_base_km = $2 WHERE id_viaje = $1
	`, tripID, distance)
	return err
}

const matchSelect = `
	SELECT m.id_match, m.compatibilidad, m.desvio_km, m.tiempo_adicional,
	       m.km_aprovechables, m.capacidad_utilizada, m.route_geometry,
	       v.codigo_viaje, v.origen, v.destino, v.fecha_salida,
	       v.fecha_llegada_estimada, v.capacidad_disponible, tv.nombre,
	       c.codigo_carga, c.origen, c.destino, c.fecha, c.tipo_carga,
	       c.peso, c.volumen, c.vehiculo_requerido, c.restricciones,
	       count(*) OVER (PARTITION BY v.id_viaje)
	FROM match_logistico m
	JOIN viaje v ON v.id_viaje = m.id_viaje
	JOIN usuario operador ON operador.id_usuario = v.id_usuario_operador
	JOIN vehiculo ve ON ve.id_vehiculo = v.id_vehiculo
	JOIN tipo_vehiculo tv ON tv.id_tipo_vehiculo = ve.id_tipo_vehiculo
	JOIN carga c ON c.id_carga = m.id_carga
	JOIN usuario solicitante ON solicitante.id_usuario = c.id_usuario
`

func (r *Postgres) ListMatches(ctx context.Context, companyID string) ([]domain.MatchItem, error) {
	rows, err := r.pool.Query(ctx, matchSelect+`
		WHERE operador.id_empresa = $1 AND solicitante.id_empresa = $1
		ORDER BY m.compatibilidad DESC, m.fecha_creacion DESC
	`, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanMatchItems(rows)
}

func (r *Postgres) GetMatch(ctx context.Context, companyID, matchID string) (domain.MatchItem, error) {
	rows, err := r.pool.Query(ctx, matchSelect+`
		WHERE operador.id_empresa = $1 AND solicitante.id_empresa = $1 AND m.id_match = $2
	`, companyID, matchID)
	if err != nil {
		return domain.MatchItem{}, err
	}
	defer rows.Close()
	items, err := scanMatchItems(rows)
	if err != nil {
		return domain.MatchItem{}, err
	}
	if len(items) == 0 {
		return domain.MatchItem{}, ErrNotFound
	}
	return items[0], nil
}

func (r *Postgres) ListTripMatches(ctx context.Context, companyID, tripID string) ([]domain.MatchItem, error) {
	rows, err := r.pool.Query(ctx, matchSelect+`
		WHERE operador.id_empresa = $1 AND solicitante.id_empresa = $1 AND v.id_viaje = $2
		ORDER BY m.compatibilidad DESC, m.fecha_creacion DESC
	`, companyID, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanMatchItems(rows)
}

type matchRows interface {
	Next() bool
	Scan(...any) error
	Err() error
}

func scanMatchItems(rows matchRows) ([]domain.MatchItem, error) {
	items := make([]domain.MatchItem, 0)
	for rows.Next() {
		var id, tripCode, tripOrigin, tripDestination string
		var loadCode, loadOrigin, loadDestination, vehicle, cargoType, requiredVehicle, restrictions string
		var score, detour, usable, utilization float64
		var additional, matchCount int
		var departure, arrival, pickup time.Time
		var freeWeight, loadWeight, loadVolume float64
		var geometryJSON []byte
		if err := rows.Scan(&id, &score, &detour, &additional, &usable, &utilization, &geometryJSON,
			&tripCode, &tripOrigin, &tripDestination, &departure, &arrival, &freeWeight, &vehicle,
			&loadCode, &loadOrigin, &loadDestination, &pickup, &cargoType, &loadWeight, &loadVolume,
			&requiredVehicle, &restrictions, &matchCount); err != nil {
			return nil, err
		}
		var geometry RouteGeometryData
		if len(geometryJSON) > 0 {
			if err := json.Unmarshal(geometryJSON, &geometry); err != nil {
				return nil, fmt.Errorf("decode match route geometry: %w", err)
			}
		}
		compatibility := int(score + 0.5)
		level := "Compatibilidad moderada"
		if compatibility >= 80 {
			level = "Alta compatibilidad"
		} else if compatibility >= 65 {
			level = "Buena compatibilidad"
		}
		dateGap := int(mathAbsDateDays(departure, pickup))
		item := domain.MatchItem{
			ID: id, Compatibility: compatibility, Level: level,
			Trip: domain.MatchTrip{
				ID: tripCode, Origin: location(tripOrigin), Destination: location(tripDestination),
				Date: spanishDate(departure, true), DateShort: spanishDate(departure, false),
				TimeWindow: departure.Format("15:04") + "–" + arrival.Format("15:04"),
				Vehicle:    vehicle, CapacityFreeTons: freeWeight, MatchCount: matchCount, Status: "matched",
			},
			Load: domain.MatchLoad{
				ID: loadCode, Origin: location(loadOrigin), Destination: location(loadDestination),
				PickupDate: spanishDate(pickup, true), PickupTime: pickup.Format("15:04"),
				CargoType: cargoType, WeightTons: loadWeight, VolumeM3: loadVolume,
				RequiredVehicle: requiredVehicle, Restrictions: restrictions,
			},
			Capacity: domain.MatchCapacity{
				FreeTons: freeWeight, LoadTons: loadWeight, UtilizationPercent: utilization,
			},
			Metrics: domain.MatchMetrics{
				DetourKM: detour, AdditionalMinutes: additional, UsableKM: usable,
				Profitability: "A calcular según tarifa",
			},
			Criteria: criteria(utilization, dateGap, detour),
			Geometry: domain.MatchGeometry{
				Planned: geometry.Planned, Utilized: geometry.Utilized, Detour: geometry.Detour,
			},
			RouteGeometry: domain.RouteGeometry{
				Coordinates: geometry.Coordinates,
				Waypoints:   geometry.Waypoints,
				Source:      geometry.Source,
			},
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

type RouteGeometryData struct {
	Coordinates [][]float64                 `json:"coordinates"`
	Waypoints   []domain.MatchRouteWaypoint `json:"waypoints"`
	Source      string                      `json:"source"`
	Planned     [][]float64                 `json:"planned"`
	Utilized    [][]float64                 `json:"utilized"`
	Detour      [][]float64                 `json:"detour"`
}

func location(city string) domain.MatchLocation {
	result := domain.MatchLocation{City: city}
	if len(city) == 0 {
		return result
	}
	if resolved, ok := resolveCity(city); ok {
		result.Region = resolved
	}
	return result
}

func resolveCity(city string) (string, bool) {
	resolved, ok := geocoder.Resolve(city)
	return resolved.Region, ok
}

var spanishMonths = [...]string{
	"ene", "feb", "mar", "abr", "may", "jun",
	"jul", "ago", "sep", "oct", "nov", "dic",
}

func spanishDate(value time.Time, includeYear bool) string {
	month := spanishMonths[int(value.Month())-1]
	if includeYear {
		return fmt.Sprintf("%d %s %d", value.Day(), month, value.Year())
	}
	return fmt.Sprintf("%d %s", value.Day(), month)
}

func mathAbsDateDays(a, b time.Time) float64 {
	yearA, monthA, dayA := a.Date()
	yearB, monthB, dayB := b.Date()
	dateA := time.Date(yearA, monthA, dayA, 0, 0, 0, 0, time.UTC)
	dateB := time.Date(yearB, monthB, dayB, 0, 0, 0, 0, time.UTC)
	days := dateA.Sub(dateB).Hours() / 24
	if days < 0 {
		return -days
	}
	return math.Abs(days)
}

func criteria(utilization float64, dateGap int, detour float64) []domain.MatchCriterion {
	routeValue, routeTone := "Compatible", "positive"
	if detour > 100 {
		routeValue, routeTone = "Desvío moderado", "warning"
	}
	dateValue := "Compatible"
	if dateGap == 1 {
		dateValue = "1 día de diferencia"
	} else if dateGap == 2 {
		dateValue = "2 días de diferencia"
	}
	dateTone := "positive"
	if dateGap > 0 {
		dateTone = "warning"
	}
	capacityValue, capacityTone := "Compatible", "positive"
	if utilization > 80 {
		capacityValue, capacityTone = "Capacidad ajustada", "warning"
	}
	return []domain.MatchCriterion{
		{Label: "Ruta", Value: routeValue, Tone: routeTone},
		{Label: "Fecha", Value: dateValue, Tone: dateTone},
		{Label: "Vehículo", Value: "Compatible", Tone: "positive"},
		{Label: "Capacidad", Value: capacityValue, Tone: capacityTone},
		{Label: "Desvío", Value: fmt.Sprintf("%.0f km", detour), Tone: routeTone},
	}
}
