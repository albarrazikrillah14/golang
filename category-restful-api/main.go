package main

import (
	"medomeckz/category-restful-api/app"
	"medomeckz/category-restful-api/controller"
	"medomeckz/category-restful-api/exception"
	"medomeckz/category-restful-api/helper"
	"medomeckz/category-restful-api/repository"
	"medomeckz/category-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {

	validate := validator.New()
	db := app.NewDB()
	repository := repository.NewCategoryRespository()
	service := service.NewCategoryService(repository, db, validate)
	controller := controller.NewCategoryController(service)

	router := httprouter.New()

	router.GET("/api/categories", controller.FindAll)
	router.GET("/api/categories/:categoryId", controller.FindById)
	router.DELETE("/api/categories/:categoryId", controller.Delete)
	router.PUT("/api/categories/:categoryId", controller.Update)
	router.POST("/api/categories", controller.Create)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
