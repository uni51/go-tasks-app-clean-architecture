package usecase_test

import (
	"2/model"
	"2/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TaskRepositoryMockはTaskRepositoryのモックです
type TaskRepositoryMock struct {
	mock.Mock
}

// CreateはモックのCreateメソッドを実装します
func (m *TaskRepositoryMock) Create(task *model.Task) (int, error) {
	args := m.Called(task)
	return args.Int(0), args.Error(1)
}

// ReadはモックのReadメソッドを実装します
func (m *TaskRepositoryMock) Read(id int) (*model.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Task), args.Error(1)
}

// UpdateはモックのUpdateメソッドを実装します
func (m *TaskRepositoryMock) Update(task *model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

// DeleteはモックのDeleteメソッドを実装します
func (m *TaskRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// TestTaskUsecaseはTaskUsecaseのテストケースです
func TestTaskUsecase(t *testing.T) {
	// モックの作成
	mockRepo := new(TaskRepositoryMock)
	taskUsecase := usecase.NewTaskUsecase(mockRepo)

	task := &model.Task{Title: "test"}

	// Createメソッドのモック設定
	mockRepo.On("Create", task).Return(1, nil)

	// タスク作成のテスト
	id, err := taskUsecase.CreateTask(task.Title)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

// TestTaskUsecase_GetTaskはGetTaskメソッドのテストケースです
func TestTaskUsecase_GetTask(t *testing.T) {
	// モックの作成
	mockRepo := new(TaskRepositoryMock)
	taskUsecase := usecase.NewTaskUsecase(mockRepo)

	task := &model.Task{Title: "test"}

	// Readメソッドのモック設定
	mockRepo.On("Read", 1).Return(task, nil)

	// タスク取得のテスト
	tt, err := taskUsecase.GetTask(1)
	assert.NoError(t, err)
	assert.Equal(t, task, tt)
}

// TestTaskUsecase_UpdateTaskはUpdateTaskメソッドのテストケースです
func TestTaskUsecase_UpdateTask(t *testing.T) {
	// モックの作成
	mockRepo := new(TaskRepositoryMock)
	taskUsecase := usecase.NewTaskUsecase(mockRepo)

	task := &model.Task{ID: 1, Title: "test"}

	// Updateメソッドのモック設定
	mockRepo.On("Update", task).Return(nil)

	// タスク更新のテスト
	err := taskUsecase.UpdateTask(task.ID, task.Title)
	assert.NoError(t, err)
}

// TestTaskUsecase_DeleteTaskはDeleteTaskメソッドのテストケースです
func TestTaskUsecase_DeleteTask(t *testing.T) {
	// モックの作成
	mockRepo := new(TaskRepositoryMock)
	taskUsecase := usecase.NewTaskUsecase(mockRepo)

	// Deleteメソッドのモック設定
	mockRepo.On("Delete", 1).Return(nil)

	// タスク削除のテスト
	err := taskUsecase.DeleteTask(1)
	assert.NoError(t, err)
}
