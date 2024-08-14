package auth

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type Service interface {
	GenerateToken(userID int) (string, error)
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
