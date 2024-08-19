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
	user, err := ctr.logic.Auth(logic.AuthParams{
		Token:        payload.Token,
		XLinkService: payload.XLink,
	})

	ok, err := ctr.logic.AuthRead(user)
	if err != nil {
		return &proto.GetAllUsersResponse{
			Data:    nil,
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	if !ok {
		return &proto.GetAllUsersResponse{
			Data:    nil,
			Status:  false,
			Message: "Unauthorized",
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
	user, err := ctr.logic.Auth(logic.AuthParams{
		Token:        payload.Token,
		XLinkService: payload.XLink,
	})

	ok, err := ctr.logic.AuthCreate(user)
	if err != nil {
		return &proto.CreateUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	if !ok {
		return &proto.CreateUserResponse{

			Status:  false,
			Message: "Unauthorized",
		}, nil
	}

	err = ctr.logic.CreateUser(payload.User)
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
	user, err := ctr.logic.Auth(logic.AuthParams{
		Token:        payload.Token,
		XLinkService: payload.XLink,
	})

	ok, err := ctr.logic.AuthUpdate(user)
	if err != nil {
		return &proto.UpdateUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	if !ok {
		return &proto.UpdateUserResponse{
			Status:  false,
			Message: "Unauthorized",
		}, nil
	}

	err = ctr.logic.UpdateUserById(int(payload.Id), payload.User)
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
	user, err := ctr.logic.Auth(logic.AuthParams{
		Token:        payload.Token,
		XLinkService: payload.XLink,
	})

	ok, err := ctr.logic.AuthDelete(user)
	if err != nil {
		return &proto.DeleteUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	if !ok {
		return &proto.DeleteUserResponse{
			Status:  false,
			Message: "Unauthorized",
		}, nil
	}

	err = ctr.logic.DeleteUserById(int(payload.UserId))
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
