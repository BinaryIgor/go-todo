package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const UNHANDLED_ERROR = "UNHANDLED_ERROR"
const UNKNOWN_ERROR = "UNKNOWN_ERROR"

func MustReadJsonBody[T any](r *http.Request, target *T) {
	err := json.NewDecoder(r.Body).Decode(target)
	//TODO: handle properly
	if err != nil {
		ThrowValidationError("INVALID_BODY", fmt.Sprintln("Expected json, but got error", err))
	}
}

func WriteJsonOkResponse(w http.ResponseWriter, response any) {
	WriteJsonResponse(w, 200, response)
}

func WriteJsonResponse(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(MustMarshalJson(response))
}

func WriteJsonErrorResponse(w http.ResponseWriter, err any) {
	var response any
	var status int
	if v, ok := err.(AppError); ok {
		response = NewApiError([]string{v.Code}, v.Message)
		status = statusFromError(v)
	} else if v, ok := err.(error); ok {
		response = NewApiError([]string{UNHANDLED_ERROR}, v.Error())
		status = 500
	} else {
		response = NewApiError([]string{UNKNOWN_ERROR}, "")
		status = 500
	}
	fmt.Println("Writing response...", response)
	WriteJsonResponse(w, status, response)
}

func statusFromError(err any) int {
	switch err.(type) {
	case ValidationError:
		return 400
	case NotFoundError:
		return 404
	default:
		return 400
	}
}

func MustMarshalJson(response any) []byte {
	jsonResp, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println("Json", string(jsonResp))
	return jsonResp
}
