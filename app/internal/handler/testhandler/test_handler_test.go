package test_handler

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/nasunagisa/restapi/app/test"
	"github.com/sebdah/goldie/v2"
)

func TestGetUser(t *testing.T) {
	// ハンドラ関数の定義
	userHandler := test.InitTest()	
	e := echo.New()

	g:=goldie.New(
		t,
		goldie.WithFixtureDir("testdata"),
		goldie.WithNameSuffix(".golden.json"),
		goldie.WithTestNameForDir(true),
		goldie.WithSubTestNameForDir(true),
	)
	test := []struct {
		name string
		in  int64
	}{
		{
			name: "正常系",
			in: 1,
		
		
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			// Call the handler
			if err := userHandler.GetUser(ctx, tt.in); err != nil {
				t.Errorf("error: %v", err)
			}
			// Assert the response using the goldie library
			g.Assert(t, tt.name, rec.Body.Bytes())
		})
	}
}

		
