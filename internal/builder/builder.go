package builder

import (
	"fmt"
	"net/http"
	"os"
	"project_structure/internal/core/model"
	"project_structure/internal/core/ports"
	"project_structure/internal/core/service"
	"project_structure/internal/inbound/handler"
	"project_structure/internal/inbound/middleware"
	"project_structure/internal/inbound/response"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type appBuilder struct {
	config   any
	logger   zerolog.Logger
	timezone *time.Location

	healthService ports.IHealthServices
	userService   ports.IUserService
	authService   ports.IAuthService

	middleware ports.IMiddleware
	server     *http.Server
	response   ports.IResponse

	router *mux.Router
}

func NewAppBuilder() ports.IBuilder {
	return &appBuilder{}
}

func (a *appBuilder) SetConfig() ports.IBuilder {
	a.config = a.loadConfig()
	return a
}

func (a *appBuilder) SetTimezone() ports.IBuilder {
	loc, err := time.LoadLocation("Africa/Lagos") // Asia/Kolkata
	if err != nil {
		a.logger.Fatal().Err(fmt.Errorf("failed to initialize timezone: %w", err)).Msg("startup failure")
	}
	a.timezone = loc
	return a
}

func (a *appBuilder) SetLogger() ports.IBuilder {
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(a.timezone)
	}
	a.logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	return a
}

func (a *appBuilder) SetMiddleware() ports.IBuilder {
	a.middleware = middleware.NewMiddlewareObj()
	return a
}

func (a *appBuilder) SetResponse() ports.IBuilder {
	a.response = response.NewResponsepObj()
	return a
}

func (a *appBuilder) SetHealthService() ports.IBuilder {
	a.healthService = service.NewHealthServiceObj(a.response)
	return a
}

func (a *appBuilder) SetUserService() ports.IBuilder {
	a.userService = service.NewUserServiceObj(a.response)
	return a
}

func (a *appBuilder) SetAuthService() ports.IBuilder {
	a.authService = service.NewAuthServiceObj(a.response)
	return a
}

func (a *appBuilder) SetHandlers() ports.IBuilder {
	handlerObj, router := handler.NewHandlerObj(a.middleware)
	a.router = router

	handlerObj.SetCommonRouting(a.healthService)

	handlerObj.SetAuthRouting(a.authService)

	handlerObj.SetUserRouting(a.userService)

	return a
}

func (a *appBuilder) SetServer() ports.IBuilder {
	srvc := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}
	a.server = srvc

	return a
}

func (a *appBuilder) Build() *model.App {
	return &model.App{
		Timezone: a.timezone,
		Server:   a.server,
		Log:      a.logger,
	}
}
