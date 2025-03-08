package acl

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId string
	FileId string
	Mode   Mode
	Size   int64
	jwt.RegisteredClaims
}

type JWT struct {
}

func NewJWT() *JWT {
	return &JWT{}
}

func (t *JWT) Generate(claim *Claims) ([]byte, error) {
	secretKeyString := os.Getenv("JWT_SECRETE_KEY")
	secretKey := []byte(secretKeyString)

	claim.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    os.Getenv("JWT_ISSUER"),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return []byte(signedToken), nil
}
