package responses

type UserResponse struct {
	Id       int64  `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}
type FailedRequest struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
