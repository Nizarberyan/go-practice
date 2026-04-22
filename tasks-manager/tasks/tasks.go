// Package tasks: the common tasks library
package tasks

import (
	"context"

	"practice/task-manager/db/db"
)

type TaskService struct {
	queries *db.Queries
}

func NewTaskService(q *db.Queries) *TaskService {
	return &TaskService{
		queries: q,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task db.CreateTaskParams) (db.Task, error) {
	return s.queries.CreateTask(ctx, task)
}

func (s *TaskService) GetTasks(ctx context.Context) ([]db.Task, error) {
	return s.queries.ListTasks(ctx)
}

func (s *TaskService) GetTask(ctx context.Context, id int64) (db.Task, error) {
	return s.queries.GetTask(ctx, id)
}

func (s *TaskService) DeleteTask(ctx context.Context, id int64) error {
	return s.queries.DeleteTask(ctx, id)
}
