package domain

import "github.com/golang-jwt/jwt"

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaims struct {
	jwt.StandardClaims
	Adminid uint
	Email   string
	Source  string
}
