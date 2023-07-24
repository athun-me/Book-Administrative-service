package interfaces

import (
	"github.com/golang-jwt/jwt"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/domain"
)

type JwtUseCase interface {
	GenerateAccessToken(adminid int, email string, role string) (string, error)
	VerifyToken(token string) (bool, *domain.JwtClaims)
	GetTokenFromString(signedToken string, claims *domain.JwtClaims) (*jwt.Token, error)
}
