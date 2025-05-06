package builder

import (
	"fmt"
	"net/http"
	"project_structure/internal/core/model/app"
	"project_structure/internal/core/ports/builder"
	middlewarePorts "project_structure/internal/core/ports/middleware"
	responsePort "project_structure/internal/core/ports/response"
	userPorts "project_structure/internal/core/ports/user"
	"project_structure/internal/core/service/user"
	"project_structure/internal/inbound/handler"
	"project_structure/internal/inbound/middleware"
	"project_structure/internal/inbound/response"
	"time"
)

type appBuilder struct {
	config      any
	logger      any
	timezone    *time.Location
	userService userPorts.IUser
	middleware  middlewarePorts.IMiddleware
	server      *http.Server
	response    responsePort.IResponse
}

func NewAppBuilder() builder.IBuilder {
	return &appBuilder{}
}

func (a *appBuilder) SetConfig() builder.IBuilder {
	a.config = a.loadConfig()
	return a
}

func (a *appBuilder) SetTimezone() builder.IBuilder {
	loc, err := time.LoadLocation("Africa/Lagos") // Asia/Kolkata
	if err != nil {
		panic("failed to load timezone: " + err.Error())
	}
	a.timezone = loc
	return a
}

func (a *appBuilder) SetLogger() builder.IBuilder {
	a.logger = nil
	return a
}

func (a *appBuilder) SetMiddleware() builder.IBuilder {
	a.middleware = middleware.NewMiddlewareObj()
	return a
}

func (a *appBuilder) SetResponse() builder.IBuilder {
	a.response = response.NewRespObj()
	return a
}

func (a *appBuilder) SetUserService() builder.IBuilder {
	a.userService = user.NewUserServiceObj(a.response)
	return a
}

func (a *appBuilder) SetServer() builder.IBuilder {
	handlers := handler.NewHandlerObj(a.middleware, a.userService)
	fmt.Println("handlers:: ", handlers)
	srvc := &http.Server{
		Addr:    ":8080",
		Handler: handlers,
	}
	a.server = srvc

	return a
}

func (a *appBuilder) Build() *app.App {
	return &app.App{
		Middleware:  a.middleware,
		UserService: a.userService,
		Timezone:    a.timezone,
		Server:      a.server,
	}
}
