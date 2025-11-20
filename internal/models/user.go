package models

type User struct {
	id       int
	username string
	passHash string
	isAdmin  bool
}
