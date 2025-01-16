package models

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Country      string `json:"country"`
	ProfileImage string `json:"profileimage"`
	Password     string `json:"password"`
	Verified     bool   `json:"verified"`
	Role         string `json:"role"`
}
