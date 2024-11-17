package exception

import (
	"medomeckz/category-restful-api/helper"
	"medomeckz/category-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if validationError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(w, response)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NotFound",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, response)
		return true
	} else {
		return false
	}
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, response)
		return true
	} else {
		return false
	}
}
