package dto

import "time"

type UserCreateResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateRequest struct {
	Email    string `json:"email" valid:"email~Invalid format email"`
	Username string `json:"username"`
}

type UserUpdateResponse struct {
	Id        uint       `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	UpdatedAt *time.Time `json:"updated_at"`
}
