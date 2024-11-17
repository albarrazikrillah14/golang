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
	category := web.CategoyCreateRequest{}
	helper.ReadFromRequestBody(r, category)

	data := c.Service.Create(r.Context(), category)

	response := &web.WebResponse{
		Code:   201,
		Status: "success",
		Data:   data,
	}

	helper.WriteResponse(w, response)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	category := &web.CategoryUpdateRequset{}
	helper.ReadFromRequestBody(r, category)

	id := p.ByName("categoryId")
	categoryId, err := strconv.Atoi(id)
	helper.PanicIfError(err)
	category.Id = categoryId

	data := c.Service.Update(r.Context(), *category)

	response := &web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   data,
	}

	helper.WriteResponse(w, response)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("categoryId")
	categoryId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	c.Service.Delete(r.Context(), categoryId)

	response := &web.WebResponse{
		Code:   200,
		Status: "success",
	}

	helper.WriteResponse(w, response)
}

func (c *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("categoryId")
	categoryId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	data := c.Service.FindById(r.Context(), categoryId)

	response := &web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   data,
	}

	helper.WriteResponse(w, response)
}

func (c *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data := c.Service.FindAll(r.Context())

	response := &web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   data,
	}

	helper.WriteResponse(w, response)
}
