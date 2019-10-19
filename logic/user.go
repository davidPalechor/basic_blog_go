package logic

import (
	"basic_blog_go/utils"
	"fmt"

	"basic_blog_go/auth"
	"basic_blog_go/models"
	"basic_blog_go/viewmodels"

	"golang.org/x/crypto/bcrypt"
)

// UserLogic defines a struct that will be associated to all operations
// related with users
type UserLogic struct {
	DB utils.DBOrmer
}

// NewUserLogic returns an istance of UserLogic struct
func NewUserLogic() *UserLogic {
	return &UserLogic{
		DB: utils.NewDBWrapper(),
	}
}

// CreateUser inserts a new registry into 'user' table
func (u *UserLogic) CreateUser(obj viewmodels.UserRequest) error {
	passwordHash, err := HashPassword(obj.Password)
	if err != nil {
		return fmt.Errorf("Failed to hass password")
	}

	user := models.User{
		Name:     obj.Name,
		Password: passwordHash,
		Email:    obj.Email,
		Active:   true,
	}
	if _, err := u.DB.Insert(&user); err != nil {
		return fmt.Errorf("Failed to create post")
	}
	return nil
}

// LoginUser returns a response object with a JWT value
func (u *UserLogic) LoginUser(obj viewmodels.LoginRequest) *viewmodels.LoginResponse {
	var user models.User
	query := u.DB.GetQueryTable(models.UserTableName).Filter("email", obj.Email)

	if err := u.DB.One(query, &user); err == nil {
		if match := CheckPasswordHash(obj.Password, user.Password); match {
			token, _ := auth.GetToken(&user)
			response := viewmodels.LoginResponse{
				Name:  user.Name,
				Email: user.Email,
				JWT:   token,
			}
			return &response
		}
		return nil
	}
	return nil
}

// HashPassword turns a plain text password into a hash
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("Empty password")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash returns wether a password is correct or not
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
