package jwts

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtPayload struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     int8   `json:"role"` //1:user, 2:admin
}

type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims
}

func GenToken(payload JwtPayload, accessSecret string, expires int) (string, error) {
	claim := CustomClaims{
		JwtPayload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expires) * time.Hour)),
		}, //设置过期时间等
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) //生成header和payload

	return token.SignedString([]byte(accessSecret)) //生成signature
}

func ParseToken(tokenString string, accessSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
		//将tokenString解析为jwt.Token，其中的payload解析为CustomClaims
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		//断言，将token.Claims转换为CustomClaims
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
