package models

type User struct {
	UserUUID     string `json:"uuid"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
