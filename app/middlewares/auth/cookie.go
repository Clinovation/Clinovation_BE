package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func LogoutCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie("is-login")
	if err != nil {
		return nil, err
	}
	fmt.Println(cookie)
	expire := time.Now().Add(-5 * 24 * time.Hour)
	cookieEx := http.Cookie{
		Name:    "is-login",
		Expires: expire,
	}

	c.SetCookie(&cookieEx)
	return cookie, nil
}