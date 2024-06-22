package user_api

import (
	"context"

	"github.com/amosehiguese/dc/modules/dc/store"
	"github.com/google/uuid"
)

type User interface {
	Create(ctx context.Context, body *store.UserModel) (*store.UserModel, error)
	Update(ctx context.Context, id uuid.UUID, body *store.UserModel) (*store.UserModel, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetAllUsers(ctx context.Context, filter map[string]any, sortBy []string, limit int, page int) ([]*store.UserModel, error)
	GetByID(ctx context.Context, id uuid.UUID) (*store.UserModel, error)
	GetByEmail(ctx context.Context, email string) (*store.UserModel, error)
}

type UserRepo struct {
	allowedSortByFields   []string
	allowedFilterByFields []string
}

func NewUserRepo() User {
	var user User = &UserRepo{
		allowedSortByFields: []string{
			"Name",
			"CreatedAt",
			"UpdatedAt",
		},
		allowedFilterByFields: []string{
			"Name",
			"Role",
			"IsVerified",
		},
	}
	return user
}

// Ensures that UserRepo implements User
var _ User = (*UserRepo)(nil)

func (u *UserRepo) Create(ctx context.Context, body *store.UserModel) (*store.UserModel, error) {
	c := store.GetDBClient()
	result := c.Client.WithContext(ctx).Create(&body)
	if result.Error != nil {
		return nil, result.Error
	}
	return body, nil
}

func (u *UserRepo) Update(ctx context.Context, id uuid.UUID, body *store.UserModel) (*store.UserModel, error) {
	c := store.GetDBClient()
	c.Client.WithContext(ctx).Model(&store.UserModel{}).Where("ID=?", id).Updates(body)
	return body, nil
}

func (u *UserRepo) Delete(ctx context.Context, id uuid.UUID) error {
	c := store.GetDBClient()
	c.Client.Delete(&store.UserModel{}, id)
	return nil
}

func (u *UserRepo) GetAllUsers(ctx context.Context, filter map[string]any, sortBy []string, limit int, page int) ([]*store.UserModel, error) {
	return []*store.UserModel{}, nil
}

func (u *UserRepo) GetByID(ctx context.Context, id uuid.UUID) (*store.UserModel, error) {
	return u.by(ctx, "id", id)
}

func (u *UserRepo) GetByEmail(ctx context.Context, email string) (*store.UserModel, error) {
	return u.by(ctx, "email", email)
}

func (u *UserRepo) by(ctx context.Context, key string, value any) (*store.UserModel, error) {
	c := store.GetDBClient()
	var user store.UserModel
	if row := c.Client.WithContext(ctx).Where(key+"=?", value).First(&user); row.Error != nil {
		return nil, row.Error
	}

	return &user, nil
}
