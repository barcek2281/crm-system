package model

type User struct {
	ID          string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	// Email       string `json:"email"`
	// PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	Password    string `json:"password"` // lol.
	IsAdmin     bool   `json:"is_admin,omitempty"`
}

type LoginRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
type LoginResponse struct {
	ID           string `json:"user_id"`
	SessionToken string `json:"sessionToken"`
}
type RegisterRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Role        string `json:"role"`
	// Email       string `json:"email" validate:"required,email"`
	// PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required,min=8"`
	Login       string `json:"passoword" validate:"required,min=4"`
}
type RegisterResponse struct {
	ID           string `json:"user_id"`
	SessionToken string `json:"sessionToken"`
}
