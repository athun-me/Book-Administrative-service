package usecas

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/domain"
	interfaces "githum.com/athunlal/bookNowAdmin-svc/pkg/usecase/interface"
)

type jwtUseCase struct {
	SecretKey string
}

// GenerateAccessToken implements interfaces.JwtUseCase.
func (u *jwtUseCase) GenerateAccessToken(adminid int, email string, role string) (string, error) {
	claims := domain.JwtClaims{
		Adminid: uint(adminid),
		Email:   email,
		Source:  "AccessToken",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(500)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(u.SecretKey))

	return accessToken, err
}

// GetTokenFromString implements interfaces.JwtUseCase.
func (u *jwtUseCase) GetTokenFromString(signedToken string, claims *domain.JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(u.SecretKey), nil
	})
}

// VerifyToken implements interfaces.JwtUseCase.
func (u *jwtUseCase) VerifyToken(token string) (bool, *domain.JwtClaims) {
	claims := &domain.JwtClaims{}

	tkn, err := u.GetTokenFromString(token, claims)
	if err != nil {
		return false, claims
	}
	if tkn.Valid {
		if err := claims.Valid(); err != nil {
			return false, claims
		}
	}
	return true, claims

}

func NewJWTuseCase() interfaces.JwtUseCase {
	return &jwtUseCase{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
