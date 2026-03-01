package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeusro/go-template/internal/domain/entity"
	"github.com/zeusro/go-template/internal/domain/repository"
)

// UserService defines the business logic for user operations
type UserService interface {
	// CreateUser creates a new user with validation
	CreateUser(ctx context.Context, email, username, name string) (*entity.User, error)

	// GetUser retrieves a user by ID
	GetUser(ctx context.Context, id uint) (*entity.User, error)

	// GetUserByEmail retrieves a user by email
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)

	// UpdateUser updates user information
	UpdateUser(ctx context.Context, id uint, name *string, active *bool) (*entity.User, error)

	// DeleteUser soft deletes a user
	DeleteUser(ctx context.Context, id uint) error

	// ListUsers retrieves a paginated list of users
	ListUsers(ctx context.Context, limit, offset int) ([]*entity.User, int64, error)
}

// userService implements UserService
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, email, username, name string) (*entity.User, error) {
	// Validate email format (basic validation)
	if email == "" {
		return nil, errors.New("email is required")
	}
	if username == "" {
		return nil, errors.New("username is required")
	}

	// Check if user already exists
	existing, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existing != nil {
		return nil, errors.New("user with this email already exists")
	}

	existing, err = s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existing != nil {
		return nil, errors.New("user with this username already exists")
	}

	user := &entity.User{
		Email:    email,
		Username: username,
		Name:     name,
		Active:   true,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id uint, name *string, active *bool) (*entity.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if name != nil {
		user.Name = *name
	}
	if active != nil {
		user.Active = *active
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return errors.New("user not found")
	}

	if err := s.userRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (s *userService) ListUsers(ctx context.Context, limit, offset int) ([]*entity.User, int64, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	users, total, err := s.userRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list users: %w", err)
	}

	return users, total, nil
}
