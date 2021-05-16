/**
 * @Author: jiangbo
 * @Description:
 * @File:  jwt
 * @Version: 1.0.0
 * @Date: 2021/05/16 9:34 下午
 */

package data

import (
	"github.com/dgrijalva/jwt-go"
	"jiang.geek/work04/model/db_model"
	"time"
)

var jwtKey = []byte("test_jiang")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user db_model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "jiang",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims:= &Claims{}
	token,err:= jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}

