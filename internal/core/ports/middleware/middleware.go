package middleware

import "net/http"

type IMiddleware interface {
	AuthMiddleware(next http.Handler) http.Handler
	LoggingMiddleware(next http.Handler) http.Handler
}
