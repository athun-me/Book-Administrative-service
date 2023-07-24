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
	pb.AuthServiceClient
}

func NewAdminHandler(useCase interfaces.AdminUseCase, jwtUseCase interfaces.JwtUseCase) *UseHandler {
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
