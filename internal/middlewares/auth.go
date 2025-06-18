package middlewares

import (
	"net/http"

	"server.go/internal/repository"
)

type AuthenticationMiddleware struct {
	SessionRepo repository.SessionRepository
}

func (amw *AuthenticationMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
		if token == "" {
			http.Error(w, "Missing session token", http.StatusUnauthorized)
			return
		}

		session, err := amw.SessionRepo.GetSession(token)
		if err != nil || !session.IsSessionValid {
			http.Error(w, "Invalid or expired session", http.StatusForbidden)
			return
		}

		// Optional: attach session to context for downstream access
		// ctx := context.WithValue(r.Context(), "session", session)
		// next.ServeHTTP(w, r.WithContext(ctx))

		next.ServeHTTP(w, r)
	})
}
