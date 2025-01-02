package router

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	openapi "github.com/nasunagisa/restapi/app/gen"
	"github.com/nasunagisa/restapi/app/internal/handler"
)
type serverImpl struct {
    uh handler.IUserHandler
    th handler.ITodoHandler
}
// Implement the ServerInterface methods
func (s *serverImpl) GetTodos(ctx echo.Context, userId int64) error {
    return s.th.GetTodos(ctx, userId)
}

func (s *serverImpl) GetUserList(ctx echo.Context, params openapi.GetUserListParams) error {
    return s.uh.GetUserList(ctx, params)
}

func (s *serverImpl) GetUser(ctx echo.Context, userId int64) error {
    return s.uh.GetUser(ctx, userId)
}

func NesRouter(uh handler.IUserHandler,th handler.ITodoHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	server := &serverImpl{
		uh: uh,
		th: th,
	}
	
	openapi.RegisterHandlers(e, server)
	return e
}


