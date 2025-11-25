package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
	IsAdmin  bool
}

type Session struct {
	User User
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

	fmt.Println(resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&req)
	if err != nil {
		panic(err)
	}

	session.User.Id = req.ID
	session.User.IsAdmin = req.IsAdmin

	fmt.Println("ID: ", req.ID)
	fmt.Println("Nome: ", req.Username)
	fmt.Println("Senha: ", req.PassHash)
	fmt.Println("Admin: ", req.IsAdmin)

	return session
}
