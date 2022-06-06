package Common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID int64
	jwt.StandardClaims
}

func GenerateToken(userid int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3600 * time.Second)
	issuer := "dy_lzp"
	claims := Claims{
		ID: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

//只检查token是否过期
func CheckToken(token string) bool {
	userClaims, err := ParseToken(token)
	if err != nil {
		return false
	}
	if time.Now().After(time.Unix(userClaims.ExpiresAt, 0)) {
		return false
	}

	return true
}
