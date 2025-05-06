package middleware

import (
	"context"
	"log"
	"net/http"
	"project_structure/internal/core/constants"
	"project_structure/internal/core/ports"
	"project_structure/internal/utils"

	"github.com/google/uuid"
)

type middle struct{}

func NewMiddlewareObj() ports.IMiddleware {
	return &middle{}
}

func (m *middle) RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), constants.REQUEST_KEY, uuid.NewString())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *middle) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Check header and auth
		next.ServeHTTP(w, r)
	})
}

func (m *middle) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := utils.GetReqID(r.Context())

		log.Printf("Incoming request | Method: %s | URL: %s | Request-ID: %s", r.Method, r.URL.Path, reqID)
		next.ServeHTTP(w, r)
	})
}
