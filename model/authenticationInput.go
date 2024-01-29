package model

type authenticationInput struct{
	Username string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}