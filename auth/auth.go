package auth

import (
	"basic_blog_go/models"
	"fmt"

	"basic_blog_go/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.Claims
}

// GetToken returns a valid token for each user
func GetToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.Name,
		"email": user.Email,
	})

	return token.SignedString([]byte(utils.GetAppConfig("api_secret")))
}

func ValidateToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.GetAppConfig("api_secret")), nil
	})

	if _, valid := token.Claims.(jwt.MapClaims); valid && token.Valid {
		return nil
	}

	return err
}
