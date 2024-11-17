package helper

import (
	"encoding/json"
	"medomeckz/category-restful-api/common/exceptions"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil {
		panic(exceptions.InvariantError(err.Error()))
	}
}

func WriteResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}
