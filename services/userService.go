package services

import (
	"context"
	"grpc-gateway/models"

	pb "grpc-gateway/pb"
)

type UserService struct {
	Client pb.UserServiceClient
	Ctx    context.Context
}

func InitUserService(client pb.UserServiceClient) *UserService {
	return &UserService{
		Client: client,
	}
}

func (u *UserService) CreateUser(user *models.UserModel) (*models.UserModel, error) {

	response, err := u.Client.CreateUser(u.Ctx, &pb.CreateUserRequest{Name: user.Name, Email: user.Email})
	if err != nil {
		return nil, err
	}
	id := response.GetId()
	user.Id = id
	return user, nil

}
