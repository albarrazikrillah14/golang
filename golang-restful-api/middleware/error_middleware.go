package middleware

import (
	"medomeckz/category-restful-api/common/exceptions"
	"medomeckz/category-restful-api/helper"
	"medomeckz/category-restful-api/model/web"
	"net/http"
)

func ErrorMiddleware(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	exception, ok := err.(exceptions.Error)
	if ok {
		w.WriteHeader(exception.StatusCode)
		response := &web.WebResponse{
			Code:   exception.StatusCode,
			Status: "fail",
			Data:   exception.Message,
		}

		helper.WriteResponse(w, response)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		response := &web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "fail",
			Data:   err,
		}

		helper.WriteResponse(w, response)
	}
}
