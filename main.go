package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	// BodyDumpミドルウェアを追加
	app.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("Request Body: %s\n", string(reqBody))
		fmt.Printf("Response Body: %s\n", string(resBody))
	}))

	// ルートハンドラ
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, world!!")
	})

	// サーバーをポート30000で起動
	app.Logger.Fatal(app.Start(":30000"))
}
