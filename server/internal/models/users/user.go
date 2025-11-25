package models

type User struct {
	Id       int
	Username string
	PassHash string
	IsAdmin  bool
}
