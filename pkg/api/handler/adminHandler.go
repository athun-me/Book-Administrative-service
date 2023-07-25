package handler

import (
	"context"
	"net/http"

	"githum.com/athunlal/bookNowAdmin-svc/pkg/domain"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/pb"
	interfaces "githum.com/athunlal/bookNowAdmin-svc/pkg/usecase/interface"
)

type AdminHandler struct {
	UseCase    interfaces.AdminUseCase
	jwtUseCase interfaces.JwtUseCase
	pb.AdminServiceServer
}

func NewAdminHandler(useCase interfaces.AdminUseCase, jwtUseCase interfaces.JwtUseCase) *AdminHandler {
	return &AdminHandler{
		UseCase:    useCase,
		jwtUseCase: jwtUseCase,
	}
}

func (h *AdminHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	admin := domain.Admin{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	adminDetail, err := h.UseCase.Login(admin)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "Error in logging the user",
		}, err
	}

	accessToken, err := h.jwtUseCase.GenerateAccessToken(int(adminDetail.Id), adminDetail.Email, "admin")
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error in Generating JWT token",
		}, err
	}
	return &pb.LoginResponse{
		Status:      http.StatusOK,
		Accesstoken: accessToken,
	}, nil
}

func (u *AdminHandler) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePassworResponse, error) {
	admin := domain.Admin{
		Id:       uint(req.Id),
		Password: req.Password,
	}
	err := u.UseCase.ChangePassword(admin)
	if err != nil {
		return &pb.ChangePassworResponse{
			Status: http.StatusNotFound,
			Error:  "Error in changing the password",
		}, err
	}
	return &pb.ChangePassworResponse{
		Status: http.StatusOK,
	}, nil
}
