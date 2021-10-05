package middleware1

import (
	"Sample_1/ipi/repositories"
	"Sample_1/ipi/responses"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Tocken struct {
	//tokenClaim string
	userRepository repositories.UserRepositories
}

func New(userRepository repositories.UserRepositories) *Tocken {
	// t := &Tocken{
	// 	tokenClaim: "Bearer gEnwYRcEbgHAyrzWYtWRYrPSmxCmpJQv9xmAPe4RsDaBdyB47UB2JPbJRCCbmd7WmMKpaEFBQAc3H423",
	// }
	return &Tocken{
		userRepository: userRepository,
	}
}

func (t *Tocken) Authenticate(r *http.Request) (string, error) {
	tocken := r.Header.Get("Authorization")
	fmt.Println(tocken)
	if tocken == "" {
		return "", errors.New("Empty Tocken")
	}
	result, _ := t.userRepository.GetTocken(tocken)
	if result != true {
		return "", errors.New("invalid Tocken")
	}
	return tocken, nil
}
func (t *Tocken) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := t.Authenticate(r)
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
			next.ServeHTTP(w, r)
		})
	}
}
