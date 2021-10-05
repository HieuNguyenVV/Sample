package handlers

import (
	"Sample_1/ipi/repositories"
	"Sample_1/ipi/responses"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/copier"
)

type user struct {
	userRepository repositories.UserRepositories
}

func NewUser(userRepository repositories.UserRepositories) (*user, error) {
	return &user{
		userRepository: userRepository,
	}, nil
}

func (m *user) GetInforUser(w http.ResponseWriter, r *http.Request) {
	tocken := r.Header.Get("Authorization")
	fmt.Println(tocken)
	result, err := m.userRepository.GetUserbyToken(tocken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error happened in reading data to db. Err: %s", err)
		return
	}
	user := responses.UserResponse{}
	err = copier.Copy(&user, &result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error happened in mapping. Err: %s", err)
		return
	}
	resp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	w.Write(resp)
}
