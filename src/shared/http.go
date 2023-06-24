package shared

import (
	"encoding/json"
	"net/http"
)

func WriteJsonOkResponse(w http.ResponseWriter, response any) {
	WriteJsonResponse(w, 200, response)
}

func WriteJsonResponse(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(MustMarshalJson(response))
}

func MustMarshalJson(response any) []byte {
	jsonResp, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return jsonResp
}
