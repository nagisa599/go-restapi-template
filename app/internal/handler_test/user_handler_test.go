package handler_test

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/nasunagisa/restapi/app/config/Error"
	"github.com/sebdah/goldie/v2"
)

func TestGetUser(t *testing.T) {
	// Echoインスタンスの作成
	e := echo.New()

	// Goldieの設定
	g := goldie.New(
		t,
		goldie.WithFixtureDir("testdata"),
		goldie.WithNameSuffix(".golden.json"),
		goldie.WithTestNameForDir(true),
		goldie.WithSubTestNameForDir(true),
	)

	// テストケースの配列
	tests := []struct {
		name     string
		in       int64
		wantErr  bool
		errType  *Error.CustomerError
	}{
		{
			name:    "正常系",
			in:      1,
			wantErr: false,
		},
		{
			name:    "NotFoundエラー",
			in:      9999, // 存在しないID
			wantErr: true,
			errType: Error.NewNotFoundError(),
		},
	}

	// テストケースの実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			// ハンドラ関数を呼び出し
			err := userHandler.GetUser(ctx, tt.in)
			if tt.wantErr {
				if err == nil || err.Error() != tt.errType.Error() {
					t.Errorf("GetUser() error = %v, wantErr %v", err, tt.errType.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				// レスポンスの検証
				g.Assert(t, tt.name, rec.Body.Bytes())
			}
		})
	}
}
