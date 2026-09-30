package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type API struct {
	service *service.Service
}

func New(svc *service.Service, frontendOrigin string) http.Handler {
	api := &API{service: svc}
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(15 * time.Second))
	router.Use(cors(frontendOrigin))

	router.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	router.Route("/api/v1", func(v1 chi.Router) {
		v1.Post("/auth/register", api.register)
		v1.Post("/auth/login", api.login)

		v1.Group(func(protected chi.Router) {
			protected.Use(api.authenticate)
			protected.Get("/auth/me", api.me)

			protected.Get("/company", api.getCompany)
			protected.With(requireRoles(domain.RoleAdmin)).Patch("/company", api.updateCompany)

			protected.Route("/users", func(users chi.Router) {
				users.Use(requireRoles(domain.RoleAdmin))
				users.Get("/", api.listUsers)
				users.Post("/", api.createUser)
				users.Get("/{id}", api.getUser)
				users.Patch("/{id}", api.updateUser)
			})

			protected.Route("/drivers", func(drivers chi.Router) {
				drivers.Use(requireRoles(domain.RoleAdmin, domain.RoleOperator))
				drivers.Get("/", api.listDrivers)
				drivers.Get("/{id}", api.getDriver)
				drivers.With(requireRoles(domain.RoleAdmin)).Post("/", api.createDriver)
				drivers.With(requireRoles(domain.RoleAdmin)).Patch("/{id}", api.updateDriver)
			})

			protected.With(requireRoles(domain.RoleAdmin, domain.RoleOperator)).Get("/vehicle-types", api.listVehicleTypes)
			protected.Route("/vehicles", func(vehicles chi.Router) {
				vehicles.Use(requireRoles(domain.RoleAdmin, domain.RoleOperator))
				vehicles.Get("/", api.listVehicles)
				vehicles.Post("/", api.createVehicle)
				vehicles.Get("/{id}", api.getVehicle)
				vehicles.Patch("/{id}", api.updateVehicle)
			})

			protected.Route("/trips", func(trips chi.Router) {
				trips.Get("/", api.listTrips)
				trips.Get("/{id}", api.getTrip)
				trips.With(requireRoles(domain.RoleAdmin, domain.RoleOperator)).Post("/", api.createTrip)
				trips.With(requireRoles(domain.RoleAdmin, domain.RoleOperator)).Patch("/{id}", api.updateTrip)
				trips.With(requireRoles(domain.RoleAdmin, domain.RoleOperator)).Get("/{tripId}/matches", api.listTripMatches)
			})

			protected.Route("/loads", func(loads chi.Router) {
				loads.Use(requireRoles(domain.RoleAdmin, domain.RoleOperator))
				loads.Get("/", api.listLoads)
				loads.Post("/", api.createLoad)
				loads.Get("/{id}", api.getLoad)
				loads.Patch("/{id}", api.updateLoad)
			})

			protected.With(requireRoles(domain.RoleAdmin, domain.RoleOperator)).Get("/matches", api.listMatches)
			protected.With(requireRoles(domain.RoleAdmin, domain.RoleOperator)).Get("/matches/{id}", api.getMatch)
		})
	})
	return router
}

type errorEnvelope struct {
	Error apiError `json:"error"`
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func decodeJSON(w http.ResponseWriter, r *http.Request, target any) bool {
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		writeError(w, service.ErrInvalidInput)
		return false
	}
	return true
}

func writeError(w http.ResponseWriter, err error) {
	status, code, message := http.StatusInternalServerError, "INTERNAL_ERROR", "No pudimos procesar la solicitud"
	switch {
	case errors.Is(err, service.ErrInvalidInput):
		status, code, message = http.StatusBadRequest, "VALIDATION_ERROR", "Los datos enviados no son válidos"
	case errors.Is(err, service.ErrUnauthorized):
		status, code, message = http.StatusUnauthorized, "UNAUTHORIZED", "La sesión no es válida"
	case errors.Is(err, service.ErrForbidden):
		status, code, message = http.StatusForbidden, "FORBIDDEN", "No tenés permisos para realizar esta acción"
	case errors.Is(err, service.ErrNotFound):
		status, code, message = http.StatusNotFound, "NOT_FOUND", "No encontramos el recurso solicitado"
	case errors.Is(err, service.ErrConflict):
		status, code, message = http.StatusConflict, "CONFLICT", "El registro ya existe"
	}
	writeJSON(w, status, errorEnvelope{Error: apiError{Code: code, Message: message}})
}
