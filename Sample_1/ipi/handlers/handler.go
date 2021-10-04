package handlers

import (
	"Sample_1/ipi/responses"
	"encoding/json"
	"log"
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
	resp, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//resputil.JSON(w, err)
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	w.Write(resp)

}
