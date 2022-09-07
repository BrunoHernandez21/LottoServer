package jwts

import (
	"lottomusic/src/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type userCredential struct {
	ID uint32 `bson:"_id" json:"id"`
	jwt.StandardClaims
}
type newUserCredential struct {
	ID uint32 `bson:"_id" json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint32) (string, time.Time) {
	expireAt := time.Now().Add(24 * time.Hour)
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512,
		&newUserCredential{
			ID: id,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{
					Time: expireAt,
				},
			},
		})

	tokenString, _ := newToken.SignedString(config.JwtKey)
	return tokenString, expireAt
}

func GenerateShortToken(id uint32) (string, time.Time) {
	expireAt := time.Now().Add(1 * time.Hour)
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512,
		&newUserCredential{
			ID: id,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{
					Time: expireAt,
				},
			},
		})
	tokenString, _ := newToken.SignedString(config.JwtKey)
	return tokenString, expireAt
}

func GenerateLongToken(id uint32) (string, time.Time) {
	expireAt := time.Now().Add(8760 * time.Hour)
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512,
		&newUserCredential{
			ID: id,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{
					Time: expireAt,
				},
			},
		})
	tokenString, _ := newToken.SignedString(config.JwtKey)
	return tokenString, expireAt
}
func GenerateUnExpiredToken(id uint32) string {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512,
		&newUserCredential{
			ID:               id,
			RegisteredClaims: jwt.RegisteredClaims{},
		})
	tokenString, _ := newToken.SignedString(config.JwtKey)
	return tokenString
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
		return config.JwtKey, nil
	})
	return token, user, err
}
