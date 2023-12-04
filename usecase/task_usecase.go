package usecase

import (
	"2/repository"
	"fmt"
)

type TaskUsecase interface {
	CreateTask(title string) error
	GetTask(id int) (*repository.Task, error)
	UpdateTask(id int, title, description string) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	r repository.TaskRepository
}

func NewTaskUsecase(r repository.TaskRepository) TaskUsecase {
	return &taskUsecase{r: r}
}

func (u *taskUsecase) CreateTask(title string) error {
	task := repository.Task{Title: title}
	id, err := u.r.Create(&task)
	fmt.Println(id)
	return err
}

func (u *taskUsecase) GetTask(id int) (*repository.Task, error) {
	t, err := u.r.Read(id)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (u *taskUsecase) UpdateTask(id int, title, description string) error {
	task := repository.Task{ID: id, Title: title}
	err := u.r.Update(&task)
	return err
}

func (u *taskUsecase) DeleteTask(id int) error {
	err := u.r.Delete(id)
	return err
}
