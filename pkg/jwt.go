package pkg

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

type Claims struct {
	Id   int
	Role string
	jwt.RegisteredClaims
}

func NewClaims(id int, role string) *Claims {
	return &Claims{
		Id:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
}

func (c *Claims) GenerateToken() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("Secret not provide")
	}
	// buat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// tanda tangan token
	return token.SignedString([]byte(jwtSecret))
}

func (c *Claims) VerifyToken(token string) error {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return errors.New("Secret not provide")
	}
	parsedToken, err := jwt.ParseWithClaims(token, c, func(t *jwt.Token) (interface{}, error) {
		// fungsi callback yang digunakan oleh parseWithClaims untuk mengambil secret
		return []byte(jwtSecret), nil
	})
	log.Println(parsedToken.Valid)

	// buat token
	if err != nil {
		return err
	}
	if !parsedToken.Valid {
		return errors.New("Expired token")
	}
	return nil
}
