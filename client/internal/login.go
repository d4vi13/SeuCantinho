package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	Online = iota
	UserNotFound
	WrongPassword
	Conflict
	Unknown
)

type Request struct {
	ID       int    `json:"Id"`
	Username string `json:"Username"`
	PassHash string `json:"PassHash"`
	IsAdmin  bool   `json:"IsAdmin"`
}

type User struct {
	Id       int
	Username string
	Password string
}

type Session struct {
	Status  int
	User    User
	IsAdmin bool
}

func Login(username string, password string) *Session {
	session := &Session{}
	var req Request

	session.User.Username = username
	session.User.Password = password

	payload := map[string]interface{}{
		"username": session.User.Username,
		"password": session.User.Password,
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://server:8080/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	session.User.Id = -1

	if resp.StatusCode == http.StatusNotFound {
		session.Status = UserNotFound
		return session
	}

	if resp.StatusCode == http.StatusBadRequest {
		session.Status = WrongPassword
		return session

	}

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&req)
		if err != nil {
			panic(err)
		}

		session.User.Id = req.ID
		session.IsAdmin = req.IsAdmin
		session.Status = Online

		return session
	}

	session.Status = Unknown

	return session
}

func CreateUser(username string, password string) *Session {
	session := &Session{}
	var req Request

	session.User.Username = username
	session.User.Password = password

	payload := map[string]interface{}{
		"username": session.User.Username,
		"password": session.User.Password,
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://server:8080/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	session.User.Id = -1

	if resp.StatusCode == http.StatusConflict {
		session.Status = Conflict
		return session

	}

	if resp.StatusCode == http.StatusCreated {
		err = json.NewDecoder(resp.Body).Decode(&req)
		if err != nil {
			panic(err)
		}

		session.User.Id = req.ID
		session.IsAdmin = req.IsAdmin
		session.Status = Online

		return session
	}

	session.Status = Unknown

	return session
}
