package model

import "errors"

type Task struct {
	ID    int    `json:"id"`    // タスクの一意の識別子
	Title string `json:"title"` // タスクのタイトル
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	return nil
}
