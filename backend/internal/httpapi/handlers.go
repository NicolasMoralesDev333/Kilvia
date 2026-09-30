package httpapi

import (
	"net/http"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
	"github.com/go-chi/chi/v5"
)

func (a *API) register(w http.ResponseWriter, r *http.Request) {
	var input domain.RegisterCompanyInput
	if !decodeJSON(w, r, &input) {
		return
	}
	result, err := a.service.Register(r.Context(), input)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, result)
}

func (a *API) login(w http.ResponseWriter, r *http.Request) {
	var input domain.LoginInput
	if !decodeJSON(w, r, &input) {
		return
	}
	result, err := a.service.Login(r.Context(), input)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (a *API) me(w http.ResponseWriter, r *http.Request) {
	item, err := a.service.Me(r.Context(), mustIdentity(r))
	respond(w, item, err, http.StatusOK)
}

func (a *API) getCompany(w http.ResponseWriter, r *http.Request) {
	item, err := a.service.GetCompany(r.Context(), mustIdentity(r))
	respond(w, item, err, http.StatusOK)
}

func (a *API) updateCompany(w http.ResponseWriter, r *http.Request) {
	var input domain.Company
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.UpdateCompany(r.Context(), mustIdentity(r), input)
	respond(w, item, err, http.StatusOK)
}

func (a *API) listUsers(w http.ResponseWriter, r *http.Request) {
	items, err := a.service.ListUsers(r.Context(), mustIdentity(r))
	respond(w, items, err, http.StatusOK)
}

func (a *API) getUser(w http.ResponseWriter, r *http.Request) {
	item, err := a.service.GetUser(r.Context(), mustIdentity(r), chi.URLParam(r, "id"))
	respond(w, item, err, http.StatusOK)
}

func (a *API) createUser(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateUserInput
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.CreateUser(r.Context(), mustIdentity(r), input)
	respond(w, item, err, http.StatusCreated)
}

func (a *API) updateUser(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdateUserInput
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.UpdateUser(r.Context(), mustIdentity(r), chi.URLParam(r, "id"), input)
	respond(w, item, err, http.StatusOK)
}

func (a *API) listDrivers(w http.ResponseWriter, r *http.Request) {
	items, err := a.service.ListDrivers(r.Context(), mustIdentity(r))
	respond(w, items, err, http.StatusOK)
}

func (a *API) getDriver(w http.ResponseWriter, r *http.Request) {
	item, err := a.service.GetDriver(r.Context(), mustIdentity(r), chi.URLParam(r, "id"))
	respond(w, item, err, http.StatusOK)
}

func (a *API) createDriver(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateDriverInput
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.CreateDriver(r.Context(), mustIdentity(r), input)
	respond(w, item, err, http.StatusCreated)
}

func (a *API) updateDriver(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdateDriverInput
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.UpdateDriver(r.Context(), mustIdentity(r), chi.URLParam(r, "id"), input)
	respond(w, item, err, http.StatusOK)
}

func (a *API) listVehicleTypes(w http.ResponseWriter, r *http.Request) {
	items, err := a.service.ListVehicleTypes(r.Context())
	respond(w, items, err, http.StatusOK)
}

func (a *API) listVehicles(w http.ResponseWriter, r *http.Request) {
	items, err := a.service.ListVehicles(r.Context(), mustIdentity(r))
	respond(w, items, err, http.StatusOK)
}

func (a *API) getVehicle(w http.ResponseWriter, r *http.Request) {
	item, err := a.service.GetVehicle(r.Context(), mustIdentity(r), chi.URLParam(r, "id"))
	respond(w, item, err, http.StatusOK)
}

func (a *API) createVehicle(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateVehicleInput
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.CreateVehicle(r.Context(), mustIdentity(r), input)
	respond(w, item, err, http.StatusCreated)
}

func (a *API) updateVehicle(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdateVehicleInput
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.UpdateVehicle(r.Context(), mustIdentity(r), chi.URLParam(r, "id"), input)
	respond(w, item, err, http.StatusOK)
}

func (a *API) listTrips(w http.ResponseWriter, r *http.Request) {
	items, err := a.service.ListTrips(r.Context(), mustIdentity(r))
	respond(w, items, err, http.StatusOK)
}

func (a *API) getTrip(w http.ResponseWriter, r *http.Request) {
	item, err := a.service.GetTrip(r.Context(), mustIdentity(r), chi.URLParam(r, "id"))
	respond(w, item, err, http.StatusOK)
}

func (a *API) createTrip(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateTripInput
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.CreateTrip(r.Context(), mustIdentity(r), input)
	respond(w, item, err, http.StatusCreated)
}

func (a *API) updateTrip(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdateTripInput
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.UpdateTrip(r.Context(), mustIdentity(r), chi.URLParam(r, "id"), input)
	respond(w, item, err, http.StatusOK)
}

func (a *API) listLoads(w http.ResponseWriter, r *http.Request) {
	items, err := a.service.ListLoads(r.Context(), mustIdentity(r))
	respond(w, items, err, http.StatusOK)
}

func (a *API) getLoad(w http.ResponseWriter, r *http.Request) {
	item, err := a.service.GetLoad(r.Context(), mustIdentity(r), chi.URLParam(r, "id"))
	respond(w, item, err, http.StatusOK)
}

func (a *API) createLoad(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateLoadInput
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.CreateLoad(r.Context(), mustIdentity(r), input)
	respond(w, item, err, http.StatusCreated)
}

func (a *API) updateLoad(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdateLoadInput
	if !decodeJSON(w, r, &input) {
		return
	}
	item, err := a.service.UpdateLoad(r.Context(), mustIdentity(r), chi.URLParam(r, "id"), input)
	respond(w, item, err, http.StatusOK)
}

func mustIdentity(r *http.Request) domain.Identity {
	identity, _ := identityFromContext(r.Context())
	return identity
}

func respond(w http.ResponseWriter, value any, err error, status int) {
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, status, value)
}
