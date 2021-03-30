package handler

import (
	"fmt"
	"net/http"
)

type Handler struct {
	
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello golang")
}