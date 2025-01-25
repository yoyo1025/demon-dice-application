package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	Age  int
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// DSN(Data Source Name)
	dsn := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s",
    dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// GORMでPostgreSQLに接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	// マイグレーション
	db.AutoMigrate(&User{})

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
		var users []User
		db.Find(&users)
		return c.JSON(http.StatusOK, users)
	})

	app.POST("/users", func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		db.Create(user)
		return c.JSON(http.StatusOK, user)
	})

	app.Logger.Fatal(app.Start(":3000"))
}
