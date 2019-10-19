package viewmodels

// UserRequest defines the structure for requests related to users such as
// login and creating
type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// LoginRequest defines the structure for login into the app request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse defines the structure for login response
type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	JWT   string `json:"jwt"`
}
