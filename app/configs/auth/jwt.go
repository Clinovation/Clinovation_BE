package auth

import (
	"github.com/Clinovation/Clinovation_BE/app/middlewares/auth"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func SetupJwt() auth.ConfigJWT {
	_ = godotenv.Load()

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExp := os.Getenv("JWT_EXPIRE")

	Exp, _ := strconv.Atoi(jwtExp)
	configJWT := auth.ConfigJWT{
		SecretJWT:   jwtSecret,
		ExpDuration: Exp,
	}

	return configJWT
}