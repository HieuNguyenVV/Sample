package middleware

import (
	"errors"
	"net/http"

	"github.com/titpetric/factory/resputil"
)

type Tocken struct {
	tokenClaim string
}

func New() *Tocken {
	t := &Tocken{
		tokenClaim: "Bearer AKcqHRCTHaBLnznmH3fw6bRSMBSZpa9tAngkKnGydBmST5XFGpxzgsGMuT3z7QsZ",
	}
	return t
}

func (t *Tocken) Authenticate(r *http.Request) (string, error) {
	tocken := r.Header.Get("Authorization")
	if tocken == "" {
		return "", errors.New("Empty JWT")
	}
	if tocken != t.tokenClaim {
		return "", errors.New("invalid JWT")
	}
	return tocken, nil
}
func (t *Tocken) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := t.Authenticate(r)
			if err != nil {
				resputil.JSON(w, err)
				http.Error(w, "", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
