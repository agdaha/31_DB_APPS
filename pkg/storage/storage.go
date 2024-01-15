package storage

import (
	"skillfactory/31_DB_APPS/pkg/storage/postgres"
)

type Data interface {
	NewTask(postgres.Task) (int, error)
	Tasks() ([]postgres.Task, error)
	TasksWithAuthor(int) ([]postgres.Task, error)
	TasksWithLabel(int) ([]postgres.Task, error)
	UpdateTask(int, postgres.Task) error
	DeleteTask(int) error
}
