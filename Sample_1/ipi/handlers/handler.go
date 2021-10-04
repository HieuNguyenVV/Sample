package handlers

import (
	"Sample_1/ipi/responses"
	"encoding/json"
	"net/http"
)

type InforUser interface {
	GetInforUser(w http.ResponseWriter, r *http.Request)
}
type user struct {
}

func NewUser() InforUser {
	return &user{}
}
func (m *user) GetInforUser(w http.ResponseWriter, r *http.Request) {
	res := responses.UserResponse{
		Fullname: "Nguyen Van Hieu",
		Username: "HieuNV38",
		Gender:   "Male",
		Birthday: "16/02/1998",
	}

	json.NewEncoder(w).Encode(res)
}
