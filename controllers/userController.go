package controllers

import (
	"context"
	"grpc-gateway/models"
	pb "grpc-gateway/pb"
	"grpc-gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Client pb.UserServiceClient
	Ctx    context.Context
}

func InitUserController(client pb.UserServiceClient, ctx context.Context) *UserController {
	return &UserController{
		Client: client,
		Ctx:    ctx,
	}
}

func (u *UserController) CreateUserController(c *gin.Context) {
	body := models.UserModel{}
	resp := models.ResponseModel{}
	err := c.BindJSON(&body)
	if err != nil {
		resp.Metadata = models.Metadata{
			Status:  "error",
			Code:    "400",
			Message: "Invalid request body",
		}
		resp.Data = nil
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}
	createUserService := services.InitUserService(u.Client, u.Ctx)
	result, err := createUserService.CreateUser(&body)
	if err != nil {
		resp.Metadata = models.Metadata{
			Status:  "error",
			Code:    "500",
			Message: err.Error(),
		}
		resp.Data = nil
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Metadata = models.Metadata{
		Status:  "OK",
		Code:    "200",
		Message: "Create user success",
	}
	resp.Data = result
	c.JSON(http.StatusOK, resp)

}

func (u *UserController) GetUserController(c *gin.Context) {
	id := c.Param("id")
	resp := models.ResponseModel{}
	getUserService := services.InitUserService(u.Client, u.Ctx)
	result, err := getUserService.GetUser(id)
	if err != nil {
		resp.Metadata = models.Metadata{
			Status:  "error",
			Code:    "500",
			Message: err.Error(),
		}
		resp.Data = nil
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Metadata = models.Metadata{
		Status:  "OK",
		Code:    "200",
		Message: "Get user success",
	}
	resp.Data = result
	c.JSON(http.StatusOK, resp)
}

func (u *UserController) UpdateUserController(c *gin.Context) {
	body := models.UserModel{}
	resp := models.ResponseModel{}
	err := c.BindJSON(&body)
	if err != nil {
		resp.Metadata = models.Metadata{
			Status:  "error",
			Code:    "400",
			Message: "Invalid request body",
		}
		resp.Data = nil
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}
	updateUserService := services.InitUserService(u.Client, u.Ctx)
	result, err := updateUserService.UpdateUser(&body)
	if err != nil {
		resp.Metadata = models.Metadata{
			Status:  "error",
			Code:    "500",
			Message: err.Error(),
		}
		resp.Data = nil
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Metadata = models.Metadata{
		Status:  "OK",
		Code:    "200",
		Message: "Update user success",
	}
	resp.Data = result
	c.JSON(http.StatusOK, resp)
}

func (u *UserController) DeleteUserController(c *gin.Context) {
	id := c.Param("id")
	resp := models.ResponseModel{}
	deleteUserService := services.InitUserService(u.Client, u.Ctx)
	result, err := deleteUserService.DeleteUser(id)
	if err != nil {
		resp.Metadata = models.Metadata{
			Status:  "error",
			Code:    "500",
			Message: err.Error(),
		}
		resp.Data = nil
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Metadata = models.Metadata{
		Status:  "OK",
		Code:    "200",
		Message: "Delete user success",
	}
	resp.Data = result
	c.JSON(http.StatusOK, resp)
}
