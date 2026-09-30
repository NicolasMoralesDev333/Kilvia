package httpapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
)

func TestRoleMiddlewareRejectsDriverFromVehicleWrite(t *testing.T) {
	protected := requireRoles(domain.RoleAdmin, domain.RoleOperator)(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	request := httptest.NewRequest(http.MethodPost, "/api/v1/vehicles", nil)
	ctx := context.WithValue(request.Context(), identityContextKey{}, domain.Identity{Role: domain.RoleDriver})
	response := httptest.NewRecorder()

	protected.ServeHTTP(response, request.WithContext(ctx))

	if response.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", response.Code)
	}
}

func TestCORSRejectsUnknownPreflightOrigin(t *testing.T) {
	handler := cors("http://localhost:4200")(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	request := httptest.NewRequest(http.MethodOptions, "/api/v1/trips", nil)
	request.Header.Set("Origin", "https://example.invalid")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", response.Code)
	}
}
