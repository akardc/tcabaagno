package routes

import (
	"encoding/json"
	"net/http"
)

type responseBody struct {
	Payload interface{} `json:"payload"`
}

func RequestHandler(handlerFunc func() (interface{}, error)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := handlerFunc()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-type", "application/json")
		resBody := responseBody{
			Payload: res,
		}
		jsonRes, err := json.Marshal(resBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(jsonRes)
	})
}
