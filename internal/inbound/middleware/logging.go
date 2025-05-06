package middleware

import (
	"context"
	"log"
	"net/http"
	"project_structure/internal/core/constants"

	"github.com/google/uuid"
)

func (m *middle) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.Method, r.Header)
		ctx := context.WithValue(r.Context(), constants.REQUEST_KEY, uuid.NewString())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
