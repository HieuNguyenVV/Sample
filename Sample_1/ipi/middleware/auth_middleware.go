package middleware1

import (
	"Sample_1/ipi/repositories"
	"Sample_1/ipi/responses"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Tocken struct {
	userRepository repositories.IUserRepository
}

func New(userRepository repositories.IUserRepository) *Tocken {
	return &Tocken{
		userRepository: userRepository,
	}
}

func (t *Tocken) Authenticate(r *http.Request) (int64, error) {
	tocken := r.Header.Get("Authorization")
	fmt.Println(tocken)
	if tocken == "" {
		return 0, errors.New("Empty Tocken")
	}
	result, err := t.userRepository.GetTocken(tocken)
	if err != nil {
		return 0, errors.New("invalid Tocken")
	}
	return result.UserID, nil
}
func (t *Tocken) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := t.Authenticate(r)
			if err != nil {
				failReq := responses.FailedRequest{false, err.Error()}
				fail, err := json.Marshal(failReq)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Printf("Error happened in JSON marshal. Err: %s", err)
					return
				}
				w.WriteHeader(http.StatusForbidden)
				w.Write(fail)
				return
			}
			ctx := context.WithValue(r.Context(), "keyID", id)
			next.ServeHTTP(w, r.WithContext(ctx))
			//next.ServeHTTP(w, r)
		})
	}
}
