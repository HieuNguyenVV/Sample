package middleware1

import (
	"Sample_1/ipi/responses"
	"encoding/json"
	"errors"
	"net/http"
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
		return "", errors.New("Empty Tocken")
	}
	if tocken != t.tokenClaim {
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
				fail, _ := json.Marshal(failReq)
				w.WriteHeader(http.StatusForbidden)
				w.Write(fail)
				// resputil.JSON(w, err)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
