// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Error defines model for error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// User defines model for user.
type User struct {
	Name   string `json:"name"`
	UserId int64  `json:"userId"`
}

// BadRequestResponse defines model for BadRequestResponse.
type BadRequestResponse = Error

// ForbiddenResponse defines model for ForbiddenResponse.
type ForbiddenResponse = Error

// InternalServerErrorResponse defines model for InternalServerErrorResponse.
type InternalServerErrorResponse = Error

// NotFoundResponse defines model for NotFoundResponse.
type NotFoundResponse = Error

// GetUserListParams defines parameters for GetUserList.
type GetUserListParams struct {
	// Page ページ番号
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// Limit 1ページあたりの表示件数
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// ユーザー一覧情報取得
	// (GET /users)
	GetUserList(ctx echo.Context, params GetUserListParams) error
	// ユーザー情報取得
	// (GET /users/{userId})
	GetUser(ctx echo.Context, userId int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUserList converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserList(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserListParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserList(ctx, params)
	return err
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId int64

	err = runtime.BindStyledParameterWithOptions("simple", "userId", ctx.Param("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUser(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/users", wrapper.GetUserList)
	router.GET(baseURL+"/users/:userId", wrapper.GetUser)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xWXWvcRhT9K+a2j7Ilx04JeixtiqHkIaVPZjFj6Xp3gqSZzIyMFyOoJAppXUhoSUxa",
	"StpgXCeFflCaxo2bPzOx4/yLMjP75V3Fazt5jB/MSLr3nHvmnjuzmxCxlLMMMyUh3ASBkrNMon34kMTX",
	"8WaOUl3vvTZvI5YpzJRZEs4TGhFFWebfkCwz73CDpDxxAKskXhEOwTyukyTvQcQI4WIQeJCilKSNEBq2",
	"mX5wURQeyKiDKTHx7wtcgxDe84fF+u6r9FEIJlxCjDISlJtyIARdP9bV77ra09W+rm+ZRf1I1wdQeHCV",
	"iVUax5hdUNhaP79R1sKorAHV2xFVPbSinulq/2jv0av7d07oWsoUiowkn6FYR/GxQbmgQtpDWpEWasVV",
	"NKn28skm9guYcWkzI0LeWPnfuj7Q9R3zf1TzNaausjyLLyg0Y2plzeQ3tnJxVNw1pmZc6Fvz539Wzb4u",
	"vzne3dLlji63dPW1Lp/r8gddfQdDHlvsoAtcMI5CUafBVWttmRIFoenewiXwQHU5ukdsozC7NVCz2f8o",
	"laBZ2zKZ8aMCYwiXHeYwvjUAY6s3MFIGK5fYUExG0iZ4F74Uj9f5wWJDnWO19DI9hz1Ziomn2RqztFQl",
	"9hvHbJZwOis5RuDBOgrpNn5+LjDlmADCKYSwMBfMBeABJ6pjJfiG0K7aaE003rld27Ynuj548fSL491f",
	"juovD3/6U1ffHt6+d/h8W5f3dbUFlkNY6xnd8AmqzyWKT6lUlk2QFJUlWp6k+N5SPH159/Hh7X/A6IMQ",
	"buYouv19CIGbzowacaoBCm+caX5ApctKlw+s/X47/nnv5c6/L549Obr7x2vIE5pSdU72lnfyhrkUBGcY",
	"1yEDVZjKaRNnbVkM2IkQpNs8gFPbuG0ncVuXDwzgoqu2iXugym+4NG3qwvTUyVvJZi5Oz5w4AwsPLp+l",
	"2tMuDHv2YJQLqroQLrc8kHmaEtE9Ze/cxpnWk7bsjy60DJQbKn/TDXNxpuk611xNn6kh8NJHfVeboR+a",
	"enDSDI8fJXJ8ncubD683dfl0c59u5nc2voiNpxjYohh0Z61cJBBCRyke+n7CIpJ0mFThleBK4K/Pg/FA",
	"D2HchorFTJe/vrr30NnZ/pj5S9df6fpHXe3Y9a2hI004TB7bo5WfA83Zp+XBxqxUjCe03bGGpGaONnp/",
	"UBT/BwAA//81Fpg6GAwAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
