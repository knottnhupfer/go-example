package users

import (
	"fmt"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"name\": \"David\"}")
}
