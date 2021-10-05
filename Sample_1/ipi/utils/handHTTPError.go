package utils

import "Sample_1/ipi/responses"

func HandleHTTPError(err error) *responses.FailedRequest {
	return &responses.FailedRequest{false, err.Error()}
}
