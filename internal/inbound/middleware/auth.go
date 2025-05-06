package middleware

import (
	"net/http"
	"project_structure/internal/core/ports/middleware"
)

type middle struct{}

func NewMiddlewareObj() middleware.IMiddleware {
	return &middle{}
}

func (m *middle) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Check header and auth
		next.ServeHTTP(w, r)
	})
}
