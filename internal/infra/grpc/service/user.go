package service

import (
	"context"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	"github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

type UserService struct {
	usecase *usecase.User
	generated.UnimplementedUserServiceServer
}

func NewUserService(usecase *usecase.User) *UserService {
	return &UserService{
		usecase: usecase,
	}
}

func (s *UserService) GetByID(ctx context.Context, req *generated.GetByIDRequest) (*generated.GetByIDResponse, error) {
	user, err := s.usecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

func convertUser(user *entity.User) *generated.GetByIDResponse {
	if user == nil {
		return nil
	}

	return &generated.GetByIDResponse{
		Id:   user.ID,
		Name: user.Name,
	}
}
