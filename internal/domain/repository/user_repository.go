package repository

import (
	"context"

	"github.com/zeusro/go-template/internal/domain/entity"
)

// UserRepository defines the interface for user data access (Repository pattern)
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *entity.User) error

	// GetByID retrieves a user by ID
	GetByID(ctx context.Context, id uint) (*entity.User, error)

	// GetByEmail retrieves a user by email
	GetByEmail(ctx context.Context, email string) (*entity.User, error)

	// GetByUsername retrieves a user by username
	GetByUsername(ctx context.Context, username string) (*entity.User, error)

	// Update updates an existing user
	Update(ctx context.Context, user *entity.User) error

	// Delete soft deletes a user
	Delete(ctx context.Context, id uint) error

	// List retrieves a list of users with pagination
	List(ctx context.Context, limit, offset int) ([]*entity.User, int64, error)
}
