package ports

import "net/http"

type IMiddleware interface {
	RequestIDMiddleware(next http.Handler) http.Handler
	AuthMiddleware(next http.Handler) http.Handler
	LoggingMiddleware(next http.Handler) http.Handler
}
