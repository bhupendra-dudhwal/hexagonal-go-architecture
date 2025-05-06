package handler

import (
	"project_structure/internal/core/ports"

	"github.com/gorilla/mux"
)

type handler struct {
	route *mux.Router
}

func NewHandlerObj(middleware ports.IMiddleware) (ports.IHandler, *mux.Router) {
	route := mux.NewRouter().StrictSlash(true)
	route.Use(middleware.RequestIDMiddleware)

	return &handler{
		route: route,
	}, route
}

func (h *handler) SetCommonRouting(healthServices ports.IHealthServices) {
	commonGroup := h.route.PathPrefix("/healthz").Subrouter()
	{
		commonGroup.HandleFunc("/liveness", healthServices.Liveness).Methods("GET")
		commonGroup.HandleFunc("/readiness", healthServices.Readiness).Methods("GET")
	}
}

func (h *handler) SetUserRouting(userService ports.IUserService) {
	userGroup := h.route.PathPrefix("/users").Subrouter()
	{
		userGroup.HandleFunc("", userService.Add).Methods("POST")
		userGroup.HandleFunc("", userService.List).Methods("GET")
		userGroup.HandleFunc("{id:[0-9]+}", userService.Info).Methods("GET")
		userGroup.HandleFunc("{id:[0-9]+}", userService.Update).Methods("PUT")
		userGroup.HandleFunc("{id:[0-9]+}", userService.Delete).Methods("DELETE")
	}
}

func (h *handler) SetAuthRouting(authService ports.IAuthService) {
	authGroup := h.route.PathPrefix("/auth").Subrouter()
	{
		authGroup.HandleFunc("/login", authService.Login).Methods("POST")
		authGroup.HandleFunc("/register", authService.Register).Methods("POST")
	}
}
