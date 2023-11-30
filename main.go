package main

import (
	"2/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // Panicが発生してもリカバーしてくれる

	taskController := controller.TaskController{}

	e.GET("/tasks", taskController.Get)
	e.POST("/tasks", taskController.Create)

	e.Start(":8080")
}
