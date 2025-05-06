package app

import (
	"net/http"
	"project_structure/internal/core/ports/middleware"
	"project_structure/internal/core/ports/user"
	"time"
)

type App struct {
	Middleware  middleware.IMiddleware
	UserService user.IUser
	Timezone    *time.Location
	Server      *http.Server
}
