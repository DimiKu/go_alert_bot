package handlers

import (
	"net/http"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world"))
}
