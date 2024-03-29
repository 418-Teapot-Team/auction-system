package utils

import (
	"errors"
	"time"

	"auction-system/pkg/models"

	"github.com/golang-jwt/jwt"
)

type JwtManager struct {
	secretKey           string
	issuer              string
	AccessTokenDuration time.Duration
}

type jwtClaims struct {
	jwt.StandardClaims
}

func NewJwtManager(secretKey, issuer string, accessTokenDuration time.Duration) *JwtManager {
	return &JwtManager{
		secretKey:           secretKey,
		issuer:              issuer,
		AccessTokenDuration: accessTokenDuration,
	}
}

func (m *JwtManager) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(m.secretKey), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}

func (m *JwtManager) GenerateToken(user models.User) (signedToken string, err error) {
	claims := &jwtClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Id.String(),
			ExpiresAt: time.Now().Local().Add(m.AccessTokenDuration).Unix(),
			Issuer:    m.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(m.secretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
