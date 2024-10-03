package service

import (
	"context"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/mecitsemerci/go-todo-app/internal/domain/auth"
	"github.com/mecitsemerci/go-todo-app/internal/domain/priority"
	"github.com/mecitsemerci/go-todo-app/internal/domain/todo"
	dto "github.com/mecitsemerci/go-todo-app/internal/rest/handlers/todo"
	"github.com/mecitsemerci/go-todo-app/pkg/identity"
)

type Todo struct {
	taskRepo todo.TaskRepository
	logger   *zap.Logger
}

func NewTodo(taskRepo todo.TaskRepository, logger *zap.Logger) *Todo {
	return &Todo{
		taskRepo: taskRepo,
		logger:   logger,
	}
}

func (t *Todo) GetAll(ctx context.Context) (*dto.GetTasksOutput, error) {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "TodoService.GetAll")
	defer span.End()

	currentUser := identity.GetCurrentUserFromContext(ctx)

	var tasks []*todo.Task
	var err error

	if auth.Role(currentUser.Role()).IsAdmin() {
		tasks, err = t.taskRepo.GetAll(spanCtx)
	} else {
		tasks, err = t.taskRepo.GetAllByUser(spanCtx, auth.UserID(currentUser.UserID()))
	}

	if err != nil {
		return nil, errors.Wrap(err, "get all failed")
	}

	tasksOutput := make([]*dto.TaskOutput, 0)

	for _, task := range tasks {
		tasksOutput = append(tasksOutput, &dto.TaskOutput{
			TaskID:      task.ID.String(),
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			Priority:    task.Priority.Value(),
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return &dto.GetTasksOutput{
		Tasks:      tasksOutput,
		TotalCount: len(tasks),
		Page:       1, // TODO: Implement pagination
	}, nil
}

func (t *Todo) Find(ctx context.Context, input dto.GetTaskInput) (*dto.TaskOutput, error) {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "TodoService.Find")
	defer span.End()

	task, err := t.taskRepo.GetByID(spanCtx, todo.TaskID(input.TaskID))

	if err != nil {
		return nil, errors.Wrap(err, "get by id failed")
	}

	if task == nil {
		return nil, errors.Wrap(todo.ErrNoItemFound, "item not found")
	}

	return &dto.TaskOutput{
		TaskID:      task.ID.String(),
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		Priority:    task.Priority.Value(),
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

func (t *Todo) Create(ctx context.Context, input dto.CreateTaskInput) (*dto.TaskOutput, error) {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "TodoService.Create")
	defer span.End()

	currentUser := identity.GetCurrentUserFromContext(ctx)

	newTask := todo.NewTask(input.Title, input.Description, priority.Level(input.Priority), auth.UserID(currentUser.UserID()))

	err := t.taskRepo.Insert(spanCtx, newTask)

	if err != nil {
		return nil, errors.Wrap(err, "insert failed")
	}

	return &dto.TaskOutput{
		TaskID:      newTask.ID.String(),
		Title:       newTask.Title,
		Description: newTask.Description,
		Completed:   newTask.Completed,
		Priority:    newTask.Priority.Value(),
		CreatedAt:   newTask.CreatedAt,
		UpdatedAt:   newTask.UpdatedAt,
	}, nil
}

func (t *Todo) Update(ctx context.Context, input dto.UpdateTaskInput) error {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "TodoService.Update")
	defer span.End()

	task, err := t.taskRepo.GetByID(spanCtx, todo.TaskID(input.TaskID))

	if err != nil {
		return errors.Wrap(err, "get by id failed")
	}

	if task == nil {
		return errors.Wrap(todo.ErrNoItemFound, "item not found")
	}

	task.UpdateTitle(input.Title).
		UpdateDescription(input.Description).
		UpdatePriority(priority.Level(input.Priority)).
		UpdateCompleted(input.Completed)

	err = t.taskRepo.Update(spanCtx, task)

	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	return nil
}

func (t *Todo) Delete(ctx context.Context, input dto.DeleteTaskInput) error {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "TodoService.Delete")
	defer span.End()

	err := t.taskRepo.Delete(spanCtx, todo.TaskID(input.TaskID))

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	return nil
}
