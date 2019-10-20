package auth

import (
	"basic_blog_go/models"
	"fmt"
	"net/http"
	"time"

	"basic_blog_go/utils"

	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
)

// GetToken returns a valid token for each user
func GetToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Local().Add(time.Minute * time.Duration(60)),
	})

	return token.SignedString([]byte(utils.GetAppConfig("api_secret")))
}

// ValidateToken returns if authorization header is valid or not
func ValidateToken(ctx *context.Context) {
	requestToken := ctx.Request.Header.Get("Authorization")

	token, _ := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.GetAppConfig("api_secret")), nil
	})

	if _, valid := token.Claims.(jwt.MapClaims); valid && token.Valid {
		return
	}

	ctx.Output.SetStatus(http.StatusUnauthorized)
	ctx.Output.JSON(map[string]string{"status": "invalid token"}, true, true)
	return
}
