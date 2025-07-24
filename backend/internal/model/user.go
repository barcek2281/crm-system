package model

type User struct {
	ID        string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Job      Job     `json:"job"`
	Company  Company `json:"company"`
	Password string  `json:"password"` // lol.
	IsAdmin  bool    `json:"is_admin,omitempty"`
	Login    string  `json:"login"`
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
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Job       Job     `json:"job"`
	Company   Company `json:"company"`

	Password string `json:"password" validate:"required,min=8"`
	IsAdmin  bool   `json:"is_admin"`
	Login    string `json:"login"`
}
type RegisterResponse struct {
	ID           string `json:"user_id"`
	Login        string `json:"login"`
	SessionToken string `json:"sessionToken"`
}
