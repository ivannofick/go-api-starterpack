package libs

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Code          int    `json:"code"`
	MessageClient string `json:"message_client"`
}

type ApiResponse struct {
	Data   any            `json:"data"`
	Meta   any            `json:"meta,omitempty"`
	Status StatusResponse `json:"status"`
}

func customMessageFourHundredAndEighteen(message string) string {
	return "[I am a teapot]: " + message
}

func ResponseAPI(
	w http.ResponseWriter,
	data any,
	meta any,
	message string,
	status int,
) {
	if status == 418 {
		message = customMessageFourHundredAndEighteen(message)
		status = 400
	}

	response := ApiResponse{
		Data: data,
		Meta: meta,
		Status: StatusResponse{
			Code:          status,
			MessageClient: message,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
