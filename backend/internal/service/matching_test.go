package service

import (
	"testing"
	"time"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
)

func validCandidate() domain.MatchCandidate {
	departure := time.Date(2026, time.September, 12, 7, 0, 0, 0, time.UTC)
	return domain.MatchCandidate{
		TripStatus: "PROGRAMADO", LoadStatus: "PENDIENTE",
		TripDeparture: departure, LoadPickup: departure,
		TripWeight: 20, TripVolume: 60, LoadWeight: 8, LoadVolume: 30,
		VehicleName: "Semirremolque", LoadRequiredVehicle: "semirremolque",
	}
}

func TestEvaluateMatchScoresCompatibleCandidate(t *testing.T) {
	result, ok := EvaluateMatch(validCandidate(), 600, 650, 45, 600)
	if !ok {
		t.Fatal("expected a compatible match")
	}
	if result.Compatibility < 85 || result.Compatibility > 99 {
		t.Fatalf("unexpected score: %.2f", result.Compatibility)
	}
	if result.DetourKM != 50 || result.UsableKM != 600 {
		t.Fatalf("unexpected route metrics: %+v", result)
	}
}

func TestEvaluateMatchRejectsHardFilters(t *testing.T) {
	tests := []struct {
		name   string
		change func(*domain.MatchCandidate)
		base   float64
		route  float64
	}{
		{"trip status", func(c *domain.MatchCandidate) { c.TripStatus = "COMPLETADO" }, 600, 650},
		{"load status", func(c *domain.MatchCandidate) { c.LoadStatus = "ASIGNADA" }, 600, 650},
		{"weight capacity", func(c *domain.MatchCandidate) { c.TripWeight = 7 }, 600, 650},
		{"volume capacity", func(c *domain.MatchCandidate) { c.TripVolume = 20 }, 600, 650},
		{"vehicle", func(c *domain.MatchCandidate) { c.LoadRequiredVehicle = "Furgón" }, 600, 650},
		{"date", func(c *domain.MatchCandidate) { c.LoadPickup = c.LoadPickup.AddDate(0, 0, 3) }, 600, 650},
		{"detour absolute", func(*domain.MatchCandidate) {}, 600, 901},
		{"detour ratio", func(*domain.MatchCandidate) {}, 100, 161},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			candidate := validCandidate()
			test.change(&candidate)
			if _, ok := EvaluateMatch(candidate, test.base, test.route, 10, 600); ok {
				t.Fatal("expected hard filter to reject candidate")
			}
		})
	}
}
