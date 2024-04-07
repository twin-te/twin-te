package authv3

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/twin-te/twin-te/back/handler/common/middleware"
	authmodule "github.com/twin-te/twin-te/back/module/auth"
)

// impl handles the requests beggining with the following paths.
//   - "/:provider"
//   - "/:provider/callback"
//   - "/logout"
//   - "/google/idToken"
type impl struct {
	accessController authmodule.AccessController
	authUseCase      authmodule.UseCase
}

func New(
	accessController authmodule.AccessController,
	authUseCase authmodule.UseCase,
) *echo.Echo {
	h := &impl{
		accessController: accessController,
		authUseCase:      authUseCase,
	}

	e := echo.New()

	e.Use(
		echomiddleware.Recover(),
		echomiddleware.Logger(),
		middleware.NewEchoErrorHandler(),
		middleware.NewEchoWithActor(accessController),
	)

	e.GET("/:provider", h.handleOAuth2)
	e.GET("/:provider/callback", h.handleOAuth2Callback)
	e.GET("/logout", h.handleLogout)
	e.GET("/google/idToken", h.handleIDTokenGoogle)

	return e
}
