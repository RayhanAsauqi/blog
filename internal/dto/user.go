package dto



type RegisterRequest struct {
	Username     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Role string `json:"role" validate:"required, oneof=user admin"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role" validate:"required, oneof=user admin"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Role string `json:"role"`
}