package httpapi

import (
	"net/http"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/service"
	"github.com/go-chi/chi/v5"
)

var demoMatches = []map[string]any{
	{
		"id": "match-ba-sl-001", "compatibility": 94, "level": "Alta compatibilidad",
		"trip": map[string]any{
			"id": "VJ-024", "origin": map[string]string{"city": "Buenos Aires", "region": "CABA"},
			"destination": map[string]string{"city": "Mendoza", "region": "Mendoza"},
			"date": "12 sep 2026", "dateShort": "12 sep", "timeWindow": "06:30–08:00",
			"vehicle": "Semirremolque", "capacityFreeTons": 14.1, "matchCount": 3, "status": "matched",
		},
		"load": map[string]any{
			"id": "CG-1042", "origin": map[string]string{"city": "Buenos Aires", "region": "CABA"},
			"destination": map[string]string{"city": "San Luis", "region": "Capital"},
			"pickupDate": "12 sep 2026", "pickupTime": "07:30", "cargoType": "Carga general seca",
			"weightTons": 8.2, "volumeM3": 38, "requiredVehicle": "Semirremolque",
			"restrictions": "Sin refrigeración",
		},
		"capacity": map[string]any{"freeTons": 14.1, "loadTons": 8.2, "utilizationPercent": 58},
		"metrics": map[string]any{"detourKm": 38, "additionalMinutes": 35, "usableKm": 620, "profitability": "A calcular según tarifa"},
		"criteria": []map[string]string{
			{"label": "Ruta", "value": "Muy compatible", "tone": "positive"},
			{"label": "Fecha", "value": "Compatible", "tone": "positive"},
			{"label": "Vehículo", "value": "Compatible", "tone": "positive"},
			{"label": "Capacidad", "value": "Compatible", "tone": "positive"},
			{"label": "Desvío", "value": "Bajo", "tone": "warning"},
		},
		"geometry": map[string]any{
			"planned":  [][]float64{{-58.3816, -34.6037}, {-63.24, -33.15}, {-66.3356, -33.3017}, {-68.8458, -32.8895}},
			"utilized": [][]float64{{-58.3816, -34.6037}, {-63.24, -33.15}, {-66.3356, -33.3017}},
			"detour":   [][]float64{{-66.05, -33.22}, {-66.3356, -33.3017}},
		},
	},
}

func (a *API) listMatches(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, demoMatches)
}

func (a *API) getMatch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, match := range demoMatches {
		if match["id"] == id {
			writeJSON(w, http.StatusOK, match)
			return
		}
	}
	writeError(w, service.ErrNotFound)
}

func (a *API) listTripMatches(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, demoMatches)
}
