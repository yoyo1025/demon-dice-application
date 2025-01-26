package main

import (
	"demon-dice-application/database"
	"demon-dice-application/domain/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.GetDB()
	app := echo.New()

	// BodyDumpミドルウェアを追加
	app.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("Request Body: %s\n", string(reqBody))
		fmt.Printf("Response Body: %s\n", string(resBody))
	}))

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, world!!")
	})

	app.GET("/users", func(c echo.Context) error {
		db := database.GetDB()
		var users []model.User
		db.Find(&users)
		return c.JSON(http.StatusOK, users)
	})

	app.POST("/users", func(c echo.Context) error {
		db := database.GetDB()
		user := new(model.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		db.Create(user)
		return c.JSON(http.StatusOK, user)
	})

	app.Logger.Fatal(app.Start(":3000"))
}
