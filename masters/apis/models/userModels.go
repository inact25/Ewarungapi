package models

type UserModels struct {
	UserID       string `json:"id"`
	UserName     string `json:"userName"`
	UserEmail    string `json:"email"`
	UserPassword string `json:"password"`
	UserLevel    string `json:"level"`
}
type OAuth struct {
	UserID    string `json:"id"`
	UserLevel string `json:"level"`
}
