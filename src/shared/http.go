package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MustReadJsonBody[T any](r *http.Request, target *T) {
	err := json.NewDecoder(r.Body).Decode(target)
	//TODO: handle properly
	if err != nil {
		panic(err)
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

func MustMarshalJson(response any) []byte {
	jsonResp, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println("Json", string(jsonResp))
	return jsonResp
}
