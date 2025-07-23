package model

type User struct {
	ID string `json:"user_id"`
}

type LoginRequest struct {
}
type LoginResponse struct {
}
type RegisterRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}
type RegisterResponse struct {
	SessionToken string `json:"sessionToken"`
}
