package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // Panicが発生してもリカバーしてくれる

	e.GET("/tasks", func(c echo.Context) error {
		fmt.Println("get all tasks")
		return c.String(200, "get all tasks")
	})
	e.POST("/tasks", func(c echo.Context) error {
		fmt.Println("create tasks")
		return c.String(200, "create task")
	})

	e.Start(":8080")
}
