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

type SessionData struct {
	Status  int
	User    User
	IsAdmin bool
}

func Login(username string, password string) *SessionData {
	data := &SessionData{}
	var req Request

	data.User.Username = username
	data.User.Password = password

	payload := map[string]interface{}{
		"username": data.User.Username,
		"password": data.User.Password,
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://server:8080/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data.User.Id = -1

	if resp.StatusCode == http.StatusNotFound {
		data.Status = UserNotFound
		return data
	}

	if resp.StatusCode == http.StatusBadRequest {
		data.Status = WrongPassword
		return data

	}

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&req)
		if err != nil {
			panic(err)
		}

		data.User.Id = req.ID
		data.IsAdmin = req.IsAdmin
		data.Status = Online

		return data
	}

	data.Status = Unknown

	return data
}

func CreateUser(username string, password string) *SessionData {
	data := &SessionData{}
	var req Request

	data.User.Username = username
	data.User.Password = password

	payload := map[string]interface{}{
		"username": data.User.Username,
		"password": data.User.Password,
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://server:8080/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data.User.Id = -1

	if resp.StatusCode == http.StatusConflict {
		data.Status = Conflict
		return data

	}

	if resp.StatusCode == http.StatusCreated {
		err = json.NewDecoder(resp.Body).Decode(&req)
		if err != nil {
			panic(err)
		}

		data.User.Id = req.ID
		data.IsAdmin = req.IsAdmin
		data.Status = Online

		return data
	}

	data.Status = Unknown

	return data
}
