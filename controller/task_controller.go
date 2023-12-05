package controller

import (
	"2/model"
	"2/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// TaskControllerはタスクに関連するHTTPリクエストを処理するためのインターフェースです。
type TaskController interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
}

// NewTaskControllerはTaskControllerの新しいインスタンスを作成します。
func NewTaskController(u usecase.TaskUsecase) TaskController {
	return &taskController{u}
}

// taskControllerはTaskControllerの実装です。
type taskController struct {
	u usecase.TaskUsecase
}

// Getは指定されたタスクの詳細情報を取得します。
func (t *taskController) Get(c echo.Context) error {
	// パラメータからタスクのIDを取得
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Errorf("parse error: %v", err.Error())
		return c.JSON(http.StatusBadRequest, msg.Error()) // エラーがあれば400 Bad Requestを返す
	}

	// タスクの取得
	task, err := t.u.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err) // エラーがあれば500 Internal Server Errorを返す
	}

	return c.JSON(http.StatusOK, task) // タスクが正常に取得された場合は200 OKを返す
}

// Createは新しいタスクを作成します。
func (t *taskController) Create(c echo.Context) error {
	var task model.Task

	// リクエストボディからタスク情報をバインド
	if err := c.Bind(&task); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, nil) // バインドエラーがあれば400 Bad Requestを返す
	}

	// タスクの作成
	createdId, err := t.u.CreateTask(task.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err) // エラーがあれば500 Internal Server Errorを返す
	}

	return c.JSON(http.StatusOK, createdId) // タスクが正常に作成された場合は200 OKを返す
}
