package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nasunagisa/restapi/app/internal/domain"
	"go.uber.org/zap"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
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