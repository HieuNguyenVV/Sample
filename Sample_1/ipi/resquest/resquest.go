package resquest

type UserResquest struct {
	Id          int64         `json:"id"`
	Fullname    string        `json:"fullname"`
	Username    string        `json:"username"`
	Gender      string        `json:"gender"`
	Birthday    string        `json:"birthday"`
	Auth_tokens []*Auth_token `json:"authtokens,omitempty"`
}
type Auth_token struct {
	Id     int64  `json:"id"`
	Tocken string `json:"tocken"`
}
