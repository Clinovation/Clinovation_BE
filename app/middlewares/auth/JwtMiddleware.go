package auth

import (
	"github.com/Clinovation/Clinovation_BE/helpers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	Uuid string `json:"uuid"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT   string
	ExpDuration int
}

func (cj *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(cj.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return c.JSON(http.StatusForbidden,
				helpers.BuildErrorResponse("failed to init token",
					e, helpers.EmptyObj{}))
		}),
	}
}

// GenerateToken jwt
func (cj *ConfigJWT) GenerateToken(userID string, UserRole string) string {
	claims := JwtCustomClaims{
		userID,
		UserRole,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cj.ExpDuration))).Unix(),
		},
	}
	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(cj.SecretJWT))

	return token
}

//get Doctor
func GetDoctor(c echo.Context) *JwtCustomClaims {
	doctor := c.Get("user").(*jwt.Token)
	claims := doctor.Claims.(*JwtCustomClaims)
	return claims
}

//get Doctor
func GetNurse(c echo.Context) *JwtCustomClaims {
	doctor := c.Get("user").(*jwt.Token)
	claims := doctor.Claims.(*JwtCustomClaims)
	return claims
}

//get Medical Staff
func GetMedicalStaff(c echo.Context) *JwtCustomClaims {
	medicalStaff := c.Get("user").(*jwt.Token)
	claims := medicalStaff.Claims.(*JwtCustomClaims)
	return claims
}

//get user
func GetUser(c echo.Context) *JwtCustomClaims {
	medicalStaff := c.Get("user").(*jwt.Token)
	claims := medicalStaff.Claims.(*JwtCustomClaims)
	return claims
}