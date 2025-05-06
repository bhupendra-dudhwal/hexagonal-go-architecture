package handler

import (
	"project_structure/internal/core/ports/middleware"
	"project_structure/internal/core/ports/user"

	"github.com/gorilla/mux"
)

func NewHandlerObj(middleware middleware.IMiddleware, userService user.IUser) *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	r.Use(middleware.LoggingMiddleware)

	userGroup := r.PathPrefix("/user").Subrouter()
	userGroup.HandleFunc("", userService.Add).Methods("POST")
	userGroup.HandleFunc("", userService.GetAll).Methods("GET")
	userGroup.HandleFunc("/{:id}", userService.Get).Methods("GET")
	userGroup.HandleFunc("/{:id}", userService.Delete).Methods("DELETE")
	userGroup.HandleFunc("/{:id}", userService.Update).Methods("PUT")
	return r
}
