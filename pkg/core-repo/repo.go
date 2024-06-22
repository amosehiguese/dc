package corerepo

import "github.com/google/uuid"

type Repository[T any, U any] interface {
	GetAll(filter map[string]any, sortBy []string, limit int, page int) ([]U, error)
	GetByID(id uuid.UUID, options map[string]any) (U, error)
	DynamicGet(key any) (U, error)
	Create(body U) (U, error)
	Update(id uuid.UUID, body U) (U, error)
	Delete(id uuid.UUID) error
}
