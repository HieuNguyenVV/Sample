package handlers

import (
	"Sample_1/ipi/repositories"
	"Sample_1/ipi/responses"
	"Sample_1/ipi/utils"
	"encoding/json"
	"log"
	"net/http"

	"errors"

	"github.com/jinzhu/copier"
)

type user struct {
	userRepository repositories.IUserRepository
}

func NewUser(userRepository repositories.IUserRepository) *user {
	return &user{
		userRepository: userRepository,
	}
}

func (m *user) GetInforUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := ctx.Value("keyID")
	//Id := id.(int64)
	result, err := m.userRepository.GetUserbyID(id.(int64))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		failReq := utils.HandleHTTPError(err)
		fail, err := json.Marshal(failReq)
		w.Write(fail)
		log.Printf("Error happened in reading data to db. Err: %s", err)
		return
	}
	user := responses.UserResponse{}
	err = copier.Copy(&user, &result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		failReq := utils.HandleHTTPError(err)
		fail, err := json.Marshal(failReq)
		w.Write(fail)
		log.Printf("Error happened in mapping. Err: %s", err)
		return
	}
	resp, err := json.Marshal(user)
	//fmt.Println(resp)
	err = errors.New("erroooooooooo")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		failReq := utils.HandleHTTPError(err)
		fail, err := json.Marshal(failReq)
		w.Write(fail)
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	//json.NewEncoder(w).Encode(user)
	w.Write(resp)
}
