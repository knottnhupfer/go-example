package handlers

import (
	"net/http"
	"service/handlers/users"
)

func RegisterHandlers() {

	http.HandleFunc("/users", users.Get)
}
