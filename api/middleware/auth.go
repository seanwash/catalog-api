package middleware

import (
	"context"
	"net/http"
)

type User struct {
	Name    string
	IsAdmin bool
}

var userKey = "user"
var adminAPIKey = "ADMIN123"

// Auth receives a request and based on an optional API key header will
// inject a matching user into the context.
func Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			givenAPIKey := r.Header.Get("API-KEY")

			// Pass through unauthenticated requests.
			if givenAPIKey == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Inject admin user.
			if givenAPIKey == adminAPIKey {
				user := &User{Name: "Eric", IsAdmin: true}
				ctx := context.WithValue(r.Context(), userKey, user)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			// Inject non admin user.
			user := &User{Name: "Greg", IsAdmin: false}
			ctx := context.WithValue(r.Context(), userKey, user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// UserFromContext receives a context and returns a User struct if one is present.
func UserFromContext(ctx context.Context) *User {
	user, _ := ctx.Value(userKey).(*User)
	return user
}
