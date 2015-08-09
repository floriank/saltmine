package main

import (
	"net/http"
)

type Project struct {
	Identifier  string
	Title       string
	Description string
}

func ProjectIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, charset=UTF8")
	w.Write([]byte("foo"))
}

func ProjectCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
}
