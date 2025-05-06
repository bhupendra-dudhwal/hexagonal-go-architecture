package ports

type IHandler interface {
	SetCommonRouting(healthServices IHealthServices)
	SetUserRouting(userService IUserService)
	SetAuthRouting(authService IAuthService)
}
