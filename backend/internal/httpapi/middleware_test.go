package httpapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/repository"
	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/service"
	"github.com/golang-jwt/jwt/v5"
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

type routeTestRepository struct {
	repository.Repository
}

func TestDriverCannotAccessMatchRoutes(t *testing.T) {
	const secret = "a-secret-that-is-longer-than-thirty-two-characters"
	svc := service.New(&routeTestRepository{}, secret, time.Hour)
	handler := New(svc, "http://localhost:4200")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &service.Claims{
		UserID: "driver-1", CompanyID: "company-1", Role: domain.RoleDriver,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	raw, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range []string{
		"/api/v1/trips/trip-1/matches",
		"/api/v1/matches",
		"/api/v1/matches/recalculate",
	} {
		method := http.MethodGet
		if path == "/api/v1/matches/recalculate" {
			method = http.MethodPost
		}
		request := httptest.NewRequest(method, path, nil)
		request.Header.Set("Authorization", "Bearer "+raw)
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		if response.Code != http.StatusForbidden {
			t.Errorf("%s %s: expected driver to be rejected with 403, got %d", method, path, response.Code)
		}
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
