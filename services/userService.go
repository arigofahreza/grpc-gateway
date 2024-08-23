package services

import (
	"context"
	"fmt"
	"grpc-gateway/models"

	pb "grpc-gateway/pb"
)

type UserService struct {
	Client pb.UserServiceClient
	Ctx    context.Context
}

func InitUserService(client pb.UserServiceClient, ctx context.Context) *UserService {
	return &UserService{
		Client: client,
		Ctx:    ctx,
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

func (u *UserService) GetUser(id string) (*models.UserModel, error) {
	fmt.Println(id)
	response, err := u.Client.GetUser(u.Ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	user := &models.UserModel{
		Id:    response.GetId(),
		Name:  response.GetName(),
		Email: response.GetEmail(),
	}
	return user, nil
}

func (u *UserService) UpdateUser(user *models.UserModel) (*models.UserModel, error) {
	response, err := u.Client.UpdateUser(u.Ctx, &pb.UpdateUserRequest{Id: user.Id, Name: user.Name, Email: user.Email})
	if err != nil {
		return nil, err
	}
	user.Name = response.GetName()
	user.Email = response.GetEmail()
	return user, nil
}

func (u *UserService) DeleteUser(id string) (*string, error) {
	_, err := u.Client.DeleteUser(u.Ctx, &pb.DeleteUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &id, nil
}
