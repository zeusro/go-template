package grpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/zeusro/go-template/api/grpc/pb"
	"github.com/zeusro/go-template/internal/core/config"
	"github.com/zeusro/go-template/internal/core/logprovider"
	domainentity "github.com/zeusro/go-template/internal/domain/entity"
	"github.com/zeusro/go-template/internal/domain/service"
)

// Server wraps gRPC server
type Server struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
	logger       logprovider.Logger
}

// NewServer creates a new gRPC server
func NewServer(userService service.UserService, log logprovider.Logger) *Server {
	return &Server{
		userService: userService,
		logger:      log,
	}
}

// Start starts the gRPC server
func (s *Server) Start(cfg config.Config) error {
	port := cfg.GRPC.Port
	if port == 0 {
		port = 9090
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, s)

	// Enable reflection for testing
	if cfg.Debug {
		reflection.Register(grpcServer)
	}

	s.logger.Infof("gRPC server starting on port %d", port)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

// GetUser implements UserService.GetUser
func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.userService.GetUser(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: toProtoUser(user),
	}, nil
}

// CreateUser implements UserService.CreateUser
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.userService.CreateUser(ctx, req.Email, req.Username, req.Name)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: toProtoUser(user),
	}, nil
}

// ListUsers implements UserService.ListUsers
func (s *Server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	limit := int(req.Limit)
	if limit == 0 {
		limit = 10
	}
	offset := int(req.Offset)

	users, total, err := s.userService.ListUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	protoUsers := make([]*pb.User, len(users))
	for i, u := range users {
		protoUsers[i] = toProtoUser(u)
	}

	return &pb.ListUsersResponse{
		Users: protoUsers,
		Total: total,
	}, nil
}

// toProtoUser converts domain user to proto user
func toProtoUser(user *domainentity.User) *pb.User {
	if user == nil {
		return nil
	}
	return &pb.User{
		Id:        uint32(user.ID),
		Email:     user.Email,
		Username:  user.Username,
		Name:      user.Name,
		Active:    user.Active,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}
