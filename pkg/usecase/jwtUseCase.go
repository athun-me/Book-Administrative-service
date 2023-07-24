package usecas

import (
	"os"

	"github.com/golang-jwt/jwt"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/domain"
	interfaces "githum.com/athunlal/bookNowAdmin-svc/pkg/usecase/interface"
)

type jwtUseCase struct {
	SecretKey string
}

// GenerateAccessToken implements interfaces.JwtUseCase.
func (*jwtUseCase) GenerateAccessToken(adminid int, email string, role string) (string, error) {
	panic("unimplemented")
}

// GetTokenFromString implements interfaces.JwtUseCase.
func (*jwtUseCase) GetTokenFromString(signedToken string, claims *domain.JwtClaims) (*jwt.Token, error) {
	panic("unimplemented")
}

// VerifyToken implements interfaces.JwtUseCase.
func (*jwtUseCase) VerifyToken(token string) (bool, *domain.JwtClaims) {
	panic("unimplemented")
}

func NewJWTuseCase() interfaces.JwtUseCase {
	return &jwtUseCase{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
