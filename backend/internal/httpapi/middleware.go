package httpapi

import (
	"context"
	"net/http"
	"strings"

	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/domain"
	"github.com/NicolasMoralesDev333/Kilvia/backend/internal/service"
)

type identityContextKey struct{}

func (a *API) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			writeError(w, service.ErrUnauthorized)
			return
		}
		identity, err := a.service.ParseToken(strings.TrimSpace(strings.TrimPrefix(authorization, "Bearer ")))
		if err != nil {
			writeError(w, service.ErrUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), identityContextKey{}, identity)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func requireRoles(roles ...domain.Role) func(http.Handler) http.Handler {
	allowed := make(map[domain.Role]struct{}, len(roles))
	for _, role := range roles {
		allowed[role] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			identity, ok := identityFromContext(r.Context())
			if !ok {
				writeError(w, service.ErrUnauthorized)
				return
			}
			if _, ok := allowed[identity.Role]; !ok {
				writeError(w, service.ErrForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func identityFromContext(ctx context.Context) (domain.Identity, bool) {
	identity, ok := ctx.Value(identityContextKey{}).(domain.Identity)
	return identity, ok
}

func cors(frontendOrigin string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if origin := r.Header.Get("Origin"); origin != "" && origin == frontendOrigin {
				w.Header().Set("Access-Control-Allow-Origin", frontendOrigin)
				w.Header().Set("Vary", "Origin")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
			}
			if r.Method == http.MethodOptions {
				if r.Header.Get("Origin") != frontendOrigin {
					http.Error(w, "forbidden origin", http.StatusForbidden)
					return
				}
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
