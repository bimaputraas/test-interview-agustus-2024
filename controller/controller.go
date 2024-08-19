package controller

import (
	"coding-interview-agustus-1/logic"
	"coding-interview-agustus-1/proto"
	"context"
)

type (
	Controller struct {
		proto.UnimplementedUserServiceServer
		logic logic.Logic
		middleware
	}
)

func (ctr *Controller) Login(ctx context.Context, payload *proto.LoginPayload) (*proto.LoginResponse, error) {
	token, err := ctr.logic.Login(payload.Email, payload.Password)
	if err != nil {
		return &proto.LoginResponse{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		}, nil
	}

	return &proto.LoginResponse{
		Status:  true,
		Message: "Successfully",
		Data: &proto.LoginData{
			AccessToken: token,
		},
	}, nil
}
func (ctr *Controller) GetAllUsers(ctx context.Context, payload *proto.GetAllUsersPayload) (*proto.GetAllUsersResponse, error) {
	if err := ctr.Auth(&header{
		token:        payload.Token,
		xLinkService: payload.XLink,
		action:       logic.Read,
	}); err != nil {
		return &proto.GetAllUsersResponse{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		}, nil
	}

	users, err := ctr.logic.GetAllUsers()
	if err != nil {
		return &proto.GetAllUsersResponse{
			Data:    nil,
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &proto.GetAllUsersResponse{
		Data:    users,
		Status:  true,
		Message: "Successfully",
	}, nil
}
func (ctr *Controller) CreateUser(ctx context.Context, payload *proto.CreateUserPayload) (*proto.CreateUserResponse, error) {
	if err := ctr.Auth(&header{
		token:        payload.Token,
		xLinkService: payload.XLink,
		action:       logic.Create,
	}); err != nil {
		return &proto.CreateUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	err := ctr.logic.CreateUser(payload.User)
	if err != nil {
		return &proto.CreateUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &proto.CreateUserResponse{
		Status:  true,
		Message: "Successfully",
	}, nil
}
func (ctr *Controller) UpdateUser(ctx context.Context, payload *proto.UpdateUserPayload) (*proto.UpdateUserResponse, error) {
	if err := ctr.Auth(&header{
		token:        payload.Token,
		xLinkService: payload.XLink,
		action:       logic.Update,
	}); err != nil {
		return &proto.UpdateUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	err := ctr.logic.UpdateUserById(int(payload.Id), payload.User)
	if err != nil {
		return &proto.UpdateUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &proto.UpdateUserResponse{
		Status:  true,
		Message: "Successfully",
	}, nil
}
func (ctr *Controller) DeleteUser(ctx context.Context, payload *proto.DeleteUserPayload) (*proto.DeleteUserResponse, error) {
	if err := ctr.Auth(&header{
		token:        payload.Token,
		xLinkService: payload.XLink,
		action:       logic.Delete,
	}); err != nil {
		return &proto.DeleteUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	err := ctr.logic.DeleteUserById(int(payload.UserId))
	if err != nil {
		return &proto.DeleteUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &proto.DeleteUserResponse{
		Status:  true,
		Message: "Successfully",
	}, nil
}
