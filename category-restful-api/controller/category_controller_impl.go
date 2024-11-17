package controller

import (
	"medomeckz/category-restful-api/helper"
	"medomeckz/category-restful-api/model/web"
	"medomeckz/category-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Service: service,
	}
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)
	category := c.Service.Create(r.Context(), categoryCreateRequest)

	response := web.WebResponse{
		Code:   201,
		Status: "success",
		Data:   category,
	}

	helper.WriteToResponseBody(w, response)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryUpdateRquest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRquest)

	categoryId := p.ByName("categoryId")
	atoi, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRquest.Id = atoi
	category := c.Service.Update(r.Context(), categoryUpdateRquest)

	response := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   category,
	}

	helper.WriteToResponseBody(w, response)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId")
	atoi, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	c.Service.Delete(r.Context(), atoi)

	response := web.WebResponse{
		Code:   200,
		Status: "success",
	}

	helper.WriteToResponseBody(w, response)
}

func (c *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId")
	atoi, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	category := c.Service.FindById(r.Context(), atoi)

	response := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   category,
	}

	helper.WriteToResponseBody(w, response)

}

func (c *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categories := c.Service.FindAll(r.Context())

	response := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   categories,
	}

	helper.WriteToResponseBody(w, response)
}
