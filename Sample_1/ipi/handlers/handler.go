package handlers

import (
	"Sample_1/ipi/repositories"
	"Sample_1/ipi/responses"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/copier"
)

type user struct {
	userRepository repositories.IUserRepository
}

func NewUser(userRepository repositories.IUserRepository) (*user, error) {
	return &user{
		userRepository: userRepository,
	}, nil
}

func (m *user) GetInforUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := ctx.Value("keyID")
	Id := id.(int64)
	result, err := m.userRepository.GetUserbyID(Id)

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
	resp, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	w.Write(resp)
}
