package ports

import "project_structure/internal/core/model"

type IBuilder interface {
	SetConfig() IBuilder
	SetTimezone() IBuilder
	SetLogger() IBuilder
	SetMiddleware() IBuilder
	SetUserService() IBuilder
	SetServer() IBuilder
	SetResponse() IBuilder
	SetAuthService() IBuilder
	SetHealthService() IBuilder
	SetHandlers() IBuilder
	Build() *model.App
}
