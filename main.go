package main

import (
	"2/controller"
	"2/repository"
	"2/usecase"
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	return db, err
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // Panicが発生してもリカバーしてくれる

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	e.GET("/tasks/:id", taskController.Get)
	e.POST("/tasks", taskController.Create)
	e.Start(":8080")
}
