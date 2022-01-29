package common

import (
	"coupon/conf"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Secret string `json:"secret"`
	UserID uint32 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(secret string, userID uint32) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(conf.AppConfig.JWT.ExpireTime) * time.Second)
	claims := Claims{
		Secret: MD5(secret),
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    conf.AppConfig.JWT.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func GetJWTSecret() []byte {
	return []byte(conf.AppConfig.JWT.Secret)
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if err != nil || tokenClaims == nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Parse token result error! type mismatch!")
}
