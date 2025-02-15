package infrastructure

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	openapi "github.com/nasunagisa/restapi/app/gen"
	"github.com/nasunagisa/restapi/app/internal/domain"
	"github.com/nasunagisa/restapi/app/internal/handler"
	customMiddleware "github.com/nasunagisa/restapi/app/middleware"
	"go.uber.org/zap"
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
	// echoのインスタンスを生成
	e := echo.New()

	e.HTTPErrorHandler = CustomHTTPErrorHandler
	
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}),customMiddleware.NewApplicationLog().Log())
	server := &serverImpl{
		uh: uh,
		th: th,
	}

	openapi.RegisterHandlersWithBaseURL(e, server,"/v1")
	return e
}



func CustomHTTPErrorHandler(err error, c echo.Context) {
	fmt.Println("CustomHTTPErrorHandler")
	logger, _ := zap.NewProduction()
	if me, ok := err.(*domain.MyError); ok {
		// カスタムエラーの場合
		logger.Error(me.Msg, zap.String("stack", me.StackTrace))
		c.JSON(me.Code, echo.Map{
			"code":    me.Code,
			"message": me.Msg,
		})
	} else {
		// その他のエラーの場合、InternalServerErrorを返す
		logger.Error("対応されてないエラー")
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}
}