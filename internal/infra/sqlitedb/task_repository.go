package sqlitedb

import (
	"context"
	"database/sql"

	"github.com/mecitsemerci/go-todo-app/internal/domain/auth"
	"github.com/mecitsemerci/go-todo-app/internal/domain/todo"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) Insert(ctx context.Context, task *todo.Task) error {
	query := `INSERT INTO tasks (id, title, description, priority, created_at, updated_at, user_id) VALUES (?,?,?,?,?,?,?)`

	_, err := t.db.Exec(query,
		task.ID,
		task.Title,
		task.Description,
		task.Priority.Value(),
		task.CreatedAt,
		task.UpdatedAt,
		task.CreatedBy.String(),
	)

	return err
}

func (t *TaskRepository) GetByID(ctx context.Context, taskID todo.TaskID) (*todo.Task, error) {
	query := `SELECT id, title, description, completed, priority, deleted, created_at, updated_at, user_id FROM tasks WHERE id =?`

	var task todo.Task
	err := t.db.QueryRow(query, taskID.String()).
		Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Priority,
			&task.Deleted,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.CreatedBy,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

func (t *TaskRepository) GetAll(ctx context.Context) ([]*todo.Task, error) {
	query := `SELECT id, title, description, completed, priority, deleted, created_at, updated_at, user_id FROM tasks WHERE deleted=0`

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*todo.Task
	for rows.Next() {
		var task todo.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Priority,
			&task.Deleted,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *TaskRepository) GetAllByUser(ctx context.Context, userID auth.UserID) ([]*todo.Task, error) {
	query := `SELECT id, title, description, completed, priority, deleted, created_at, updated_at, user_id FROM tasks WHERE user_id=? AND deleted=0`

	rows, err := t.db.Query(query, userID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*todo.Task
	for rows.Next() {
		var task todo.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Priority,
			&task.Deleted,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *TaskRepository) Update(ctx context.Context, task *todo.Task) error {
	query := `UPDATE tasks SET title=?, description=?, priority=?, updated_at=?, completed=? WHERE id=?`

	_, err := t.db.Exec(query,
		task.Title,
		task.Description,
		task.Priority.Value(),
		task.UpdatedAt,
		task.Completed,
		task.ID.String(),
	)

	return err
}

func (t *TaskRepository) Delete(ctx context.Context, taskID todo.TaskID) error {
	query := `UPDATE tasks SET deleted=1 FROM tasks WHERE id=?`

	_, err := t.db.Exec(query, taskID.String())

	return err
}
