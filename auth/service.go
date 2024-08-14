package auth

import (
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token,error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	SECREAT_KEY := []byte(viper.GetString("SECRET_KEY"))

	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	SignToken, err := token.SignedString(SECREAT_KEY)
	if err != nil {
		return SignToken, err
	}

	return SignToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
    parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid token signing method")
        }

        // Konfigurasi Viper untuk membaca file .env
        viper.SetConfigName(".env")
        viper.SetConfigType("env")
        err := viper.ReadInConfig()
        if err != nil {
            log.Fatalf("Error while reading config file: %v", err)
            return nil, err
        }

        // Mendapatkan SECRET_KEY dari file konfigurasi
        SECRET_KEY := []byte(viper.GetString("SECRET_KEY"))

        // Mengembalikan kunci untuk memvalidasi token
        return SECRET_KEY, nil
    })

    if err != nil {
        return nil, err
    }

    return parsedToken, nil
}
