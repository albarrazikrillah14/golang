package main

import (
	"medomeckz/category-restful-api/app"
	"medomeckz/category-restful-api/controller"
	"medomeckz/category-restful-api/middleware"
	"medomeckz/category-restful-api/repository"
	"medomeckz/category-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	repository := repository.NewCategoryRepository()
	db := app.NewDB()
	validate := validator.New()
	service := service.NewCategoryService(repository, db, *validate)
	controller := controller.NewCategoryController(service)

	router := httprouter.New()
	router.GET("/api/categories", controller.FindAll)
	router.GET("/api/categories/:categoryId", controller.FindById)
	router.POST("/api/categories", controller.Create)
	router.DELETE("/api/categories/:categoryId", controller.Delete)
	router.PUT("/api/categories/:categoryId", controller.Update)

	router.PanicHandler = middleware.ErrorMiddleware

	server := http.Server{
		Addr:    "localhost:3001",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
