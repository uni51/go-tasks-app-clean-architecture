package repository

import (
	"2/model"
	"database/sql"
)

// Task 構造体はタスクの情報を表します。
type Task struct {
	ID    int    `json:"id"`    // タスクの一意の識別子
	Title string `json:"title"` // タスクのタイトル
}

// TaskRepository インターフェースはタスクに対するデータベース操作を定義します。
type TaskRepository interface {
	Create(task *model.Task) (int, error) // タスクを作成し、作成されたタスクのIDを返します。
	Read(id int) (*model.Task, error)     // 指定されたIDのタスクを取得します。
	Update(task *model.Task) error        // タスクを更新します。
	Delete(id int) error                  // 指定されたIDのタスクを削除します。
}

// taskRepositoryImpl 構造体は TaskRepository インターフェースの実装を提供します。
type taskRepositoryImpl struct {
	db *sql.DB // データベース接続
}

// NewTaskRepository は新しい TaskRepository インスタンスを作成します。
func NewTaskRepository(db *sql.DB) *taskRepositoryImpl {
	return &taskRepositoryImpl{db: db}
}

// Create はデータベースに新しいタスクを作成します。
func (r *taskRepositoryImpl) Create(task *model.Task) (int, error) {
	stmt := `INSERT INTO tasks (title) VALUES (?) RETURNING id`
	err := r.db.QueryRow(stmt, task.Title).Scan(&task.ID)

	return task.ID, err
}

// Read は指定されたIDのタスクをデータベースから取得します。
func (r *taskRepositoryImpl) Read(id int) (*model.Task, error) {
	// SQL クエリ文
	stmt := `SELECT id, title FROM tasks WHERE id = ?`
	// タスク構造体の初期化
	task := model.Task{}
	// データベースクエリを実行し、結果をタスク構造体にマッピング
	err := r.db.QueryRow(stmt, id).Scan(&task.ID, &task.Title)
	// エラーとともにタスク構造体へのポインタを返す
	return &task, err
}

// Update は指定されたタスクの情報をデータベースで更新します。
func (r *taskRepositoryImpl) Update(task *model.Task) error {
	// SQL クエリ文
	stmt := `UPDATE tasks SET title = ? WHERE id = ?`
	// タスク情報をデータベースで更新
	rows, err := r.db.Exec(stmt, task.Title, task.ID)
	if err != nil {
		return err
	}
	// 更新された行数を取得
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	// 更新された行がない場合はエラーを返す
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	// エラー情報を返す
	return err
}

// Delete は指定されたIDのタスクをデータベースから削除します。
func (r *taskRepositoryImpl) Delete(id int) error {
	// SQL クエリ文
	stmt := `DELETE FROM tasks WHERE id = ?`
	// データベースクエリを実行し、削除された行数を取得
	rows, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	// 削除された行数を取得
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	// 削除された行がない場合はエラーを返す
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	// エラー情報を返す
	return err
}
