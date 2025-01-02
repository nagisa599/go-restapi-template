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

// Todo defines model for todo.
type Todo struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

// User defines model for user.
type User struct {
	Name string `json:"name"`
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
	// todoリストの一覧を取得
	// (GET /todos/{userId})
	GetTodos(ctx echo.Context, userId int64) error
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

// GetTodos converts echo context to params.
func (w *ServerInterfaceWrapper) GetTodos(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId int64

	err = runtime.BindStyledParameterWithOptions("simple", "userId", ctx.Param("userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTodos(ctx, userId)
	return err
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

	router.GET(baseURL+"/todos/:userId", wrapper.GetTodos)
	router.GET(baseURL+"/users", wrapper.GetUserList)
	router.GET(baseURL+"/users/:userId", wrapper.GetUser)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xWUWvcRhD+K2baR9k6x5cS9FjaFEPJQ9o+mcOsT+O7DZJW2V0ZH4egkiikdSGhJTFp",
	"KWmDcZ0U0pbSNG7c/JmNHedflN2908k5nc92TR+C/WCkvdnvm2/mm131oc3CmEUYSQFeHziKmEUCzcv7",
	"xL+ONxMU8vpgWa+2WSQxkvqRxHFA20RSFrk3BIv0Gq6TMA4swArxl7lF0K9rJEgGED6C12w0HAhRCNJB",
	"8DTbzDA4TVMHRLuLIdHx73JcBQ/ecUfJuvZX4SLnjNsNPoo2p7FOBzxQxWOV/6ryHZXvquKWfigeqWIP",
	"UgeuMr5CfR+jMwpbHe6vlbVQlVVSnY+o/KER9Vzluwc7j17fv3NE12IkkUck+AT5GvIPNcoZFdIB0rIw",
	"UMs2o3G1l482cZjAjN02UxHyn5X/qYo9VdzR/6uarzF5lSWRf0ahEZPLq3p/bSubVXHXmJyxoefmz3+M",
	"ml2VfX24vaGyLZVtqPwrlb1Q2fcq/xZGPCbZsgsxZzFySa0Gm62xZUgkeLp7C5fAAdmL0b5iB7muVqmm",
	"P/xRSE6jjmHS40c5+uAtWcxRfKsEYys3sC01lmQ+q0umrP0bDA5IKoMTcNswp4SqI08E1lQiIuEJ8E3U",
	"OKgOo9Gq0TTIFFiM0SyJ6ayIsQ0OrCEXtn/zcw2dhw4gMQUPFuYacw1wICaya5JxdYGE29epLvqpXuqg",
	"KcxRJ+gw4wZzTmVPXj77/HD7Z5V/s3/73v6LTZXdV/kGGCpujLzogwcfofxU4xtGTkKUyAV4S/0xn20b",
	"kz1Vxd7iB6AlgmeSBGdQL7AZQrVKkidY9XjVW+81a7yVtpyj18elRuMEszhioBJDMW2cjOfSkp1wTnr1",
	"0zVSfVB8sf/j75V6bpoB21TZAw3VtHnWsZZ63Jq70GxdmL51/LIxO5vTd44dbakDl0+S7XH3gDlSsJ1w",
	"KnvgLbUcEEkYEt6b7kTdddIRZkR1H1oay9XmERPNXW2EhRtrxyR7fyaQf0yFnO7w7wzFs1d3H+/f/mto",
	"8ZsJ8t7I47E+xiY5uva0TJ03meZLKpXlKntgzuonhz/tvNr6++Xzpwd3f5tAHtCQylOy/y/zZI7RU87T",
	"hDZeTNWEqZpUu7GhMs2oDNX0q+PYY+64uXprbo3p5r64HM7fxlMMbFA0urVWwgPwoCtl7LluwNok6DIh",
	"vSuNKw13bR60BwYItZ9G2S+v7z20djZf/n+o4ktV/KDyLfN8a+TIwdfBcWY+BZq1T8uB9VkhWRzQTtcY",
	"kuo5Wh/8QZr+GwAA//8Ybox2RQ8AAA==",
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
