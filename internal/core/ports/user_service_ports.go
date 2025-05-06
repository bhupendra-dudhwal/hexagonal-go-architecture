package ports

import "net/http"

type IUserService interface {
	Add(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Info(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
