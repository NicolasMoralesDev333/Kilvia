package httpapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) listMatches(w http.ResponseWriter, r *http.Request) {
	items, err := a.service.ListMatches(r.Context(), mustIdentity(r))
	respond(w, items, err, http.StatusOK)
}

func (a *API) getMatch(w http.ResponseWriter, r *http.Request) {
	item, err := a.service.GetMatch(r.Context(), mustIdentity(r), chi.URLParam(r, "id"))
	respond(w, item, err, http.StatusOK)
}

func (a *API) listTripMatches(w http.ResponseWriter, r *http.Request) {
	items, err := a.service.ListTripMatches(r.Context(), mustIdentity(r), chi.URLParam(r, "tripId"))
	respond(w, items, err, http.StatusOK)
}

func (a *API) recalculateMatches(w http.ResponseWriter, r *http.Request) {
	count, err := a.service.RecalculateMatches(r.Context(), mustIdentity(r))
	respond(w, map[string]int{"recalculated": count}, err, http.StatusOK)
}
