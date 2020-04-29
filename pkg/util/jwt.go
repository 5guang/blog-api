package util

import (
	"blog/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	//request.ReqLogin
	Username string
	Password string
	jwt.StandardClaims

}

func GenerateToken(Username,Password string) (string, error)  {
	var (
		nowTime = time.Now()
		expireTime = nowTime.Add(3 * time.Hour)
	)

	claims := Claims{
		Username,
		Password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(token string) (*Claims, error)  {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}