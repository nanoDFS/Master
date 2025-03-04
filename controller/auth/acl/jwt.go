package acl

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nanoDFS/Master/utils/secrets"
)

type Claims struct {
	UserId string
	FileId string
	Access ACL
	Size   int64
	jwt.RegisteredClaims
}

type JWT struct {
}

func NewJWT() *JWT {
	return &JWT{}
}

func (t *JWT) Generate(claim *Claims) ([]byte, error) {
	secretKeyString := secrets.Get("JWT_SECRETE_KEY")
	secretKey := []byte(secretKeyString)

	claim.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    secrets.Get("JWT_ISSUER"),
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
