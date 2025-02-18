package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	//UserService *service.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name string `json:"name"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println("request: ")
	fmt.Println(req.Name)

	//user, err := h.UserService.CreateUser(r.Context(), req.Name)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
	//	return
	//}

	response, err := json.Marshal("success")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(response)
}
