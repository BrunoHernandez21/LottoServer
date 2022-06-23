package jwts

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("TestForFasthttpWithJWT")

type userCredential struct {
	ID uint32 `bson:"_id" json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id uint32) (string, time.Time) {
	expireAt := time.Now().Add(24 * time.Hour)
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
		},
	})
	tokenString, _ := newToken.SignedString(jwtKey)
	return tokenString, expireAt
}

func ValidateToken(tok string) (*jwt.Token, *userCredential, error) {
	user := &userCredential{}
	var temp string
	if strings.Contains(strings.ToLower(tok), "bearer ") {
		temp = tok[7:]
	} else {
		temp = tok
	}

	token, err := jwt.ParseWithClaims(temp, user, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, user, err
}
