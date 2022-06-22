package jwts

import (
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

func ValidateToken(requestToken string) (*jwt.Token, *userCredential, error) {
	user := &userCredential{}
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, user, err
}
