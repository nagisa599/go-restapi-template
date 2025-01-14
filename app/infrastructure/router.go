package infrastructure

import (
	"fmt"
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
	fmt.Println("ff")
    return s.th.GetTodos(ctx, userId)
}

func (s *serverImpl) GetUserList(ctx echo.Context, params openapi.GetUserListParams) error {
    return s.uh.GetUserList(ctx, params)
}

func (s *serverImpl) GetUser(ctx echo.Context, userId int64) error {
	fmt.Println("ff")
    return s.uh.GetUser(ctx, userId)
}

func NesRouter(uh handler.IUserHandler,th handler.ITodoHandler,eh handler.IErrorHandler) *echo.Echo {
	// echoのインスタンスを生成
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}),middleware.Logger())
	server := &serverImpl{
		uh: uh,
		th: th,
	}

	e.HTTPErrorHandler = eh.CustomHTTPErrorHandler
	
	openapi.RegisterHandlersWithBaseURL(e, server,"/v1")
	return e
}


