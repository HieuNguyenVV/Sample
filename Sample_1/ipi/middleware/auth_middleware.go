package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/titpetric/factory/resputil"
)

type JWT struct {
	tokenClaim string
}

func New() *JWT {
	jwt := &JWT{
		tokenClaim: "Bearer AKcqHRCTHaBLnznmH3fw6bRSMBSZpa9tAngkKnGydBmST5XFGpxzgsGMuT3z7QsZ",
	}
	return jwt
}

func (jwt *JWT) Authenticate(r *http.Request) (string, error) {
	auth_jwt := r.Header.Get("Authorization")
	fmt.Println(auth_jwt)
	if auth_jwt == "" {
		return "", errors.New("Empty JWT")
	}
	if auth_jwt != jwt.tokenClaim {
		return "", errors.New("invalid JWT")
	}
	return auth_jwt, nil
}
func (jwt *JWT) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := jwt.Authenticate(r)
			if err != nil {
				resputil.JSON(w, err)
				return
			}
			//json.NewEncoder(w).Encode(auth_jwt)
			next.ServeHTTP(w, r)
		})
	}
}
