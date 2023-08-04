package handler

import (
	"context"
	"errors"
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

func (u *AdminHandler) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	admin := domain.Admin{
		Id:       uint(req.Id),
		Password: req.Password,
	}
	err := u.UseCase.ChangePassword(admin)
	if err != nil {
		return &pb.ChangePasswordResponse{
			Status: http.StatusNotFound,
			Error:  "Error in changing the password",
		}, err
	}
	return &pb.ChangePasswordResponse{
		Status: http.StatusOK,
	}, nil
}

func (u *AdminHandler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	adminData := domain.Admin{}
	ok, claims := u.jwtUseCase.VerifyToken(req.Accesstoken)
	if !ok {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Error:  "Token Verification Failed",
		}, errors.New("Token failed")
	}
	adminData, err := u.UseCase.ValidateJwtAdmin(claims.Adminid)
	if err != nil {
		return &pb.ValidateResponse{
			Status:  http.StatusUnauthorized,
			Adminid: int64(adminData.Id),
			Error:   "Admin not found with essesntial token credential",
			Source:  claims.Source,
		}, err
	}
	return &pb.ValidateResponse{
		Status:  http.StatusOK,
		Adminid: int64(adminData.Id),
		Source:  claims.Source,
	}, nil
}
