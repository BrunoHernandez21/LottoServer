package jwts

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtSignKey = []byte("TestForFasthttpWithJWT")

type userCredential struct {
	ID string `bson:"_id" json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id primitive.ObjectID) (string, time.Time) {

	expireAt := time.Now().Add(48 * time.Hour)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		ID: id.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
		},
	})

	tokenString, _ := newToken.SignedString(jwtSignKey)

	return tokenString, expireAt
}

func ValidateToken(requestToken string) (*jwt.Token, *userCredential, error) {
	user := &userCredential{}
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})

	return token, user, err
}
