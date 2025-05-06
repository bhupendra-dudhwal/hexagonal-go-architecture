package builder

import "project_structure/internal/core/model/app"

type IBuilder interface {
	SetConfig() IBuilder
	SetTimezone() IBuilder
	SetLogger() IBuilder
	SetMiddleware() IBuilder
	SetUserService() IBuilder
	SetServer() IBuilder
	SetResponse() IBuilder
	Build() *app.App
}
