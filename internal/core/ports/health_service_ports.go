package ports

import "net/http"

type IHealthServices interface {
	Readiness(w http.ResponseWriter, r *http.Request)
	Liveness(w http.ResponseWriter, r *http.Request)
}
