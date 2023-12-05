package controller

import (
	"2/model"
	"2/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskController interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
}

func NewTaskController(u usecase.TaskUsecase) TaskController {
	return &taskController{u}
}

type taskController struct {
	u usecase.TaskUsecase
}

// type Task struct {
// 	ID    int    `json:"id"`
// 	Title string `json:"title"`
// }

func (t *taskController) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Errorf("parse error: %v", err.Error())
		return c.JSON(http.StatusBadRequest, msg.Error())
	}

	task, err := t.u.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, task)
}

func (t *taskController) Create(c echo.Context) error {
	var task model.Task
	if err := c.Bind(&task); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, nil)
	}
	createdId, err := t.u.CreateTask(task.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, createdId)
}
