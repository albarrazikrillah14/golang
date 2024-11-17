package helper

import (
	"medomeckz/category-restful-api/model/domain"
	"medomeckz/category-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) *web.CategoryResponse {
	return &web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoriesResponse(categories []domain.Category) *[]web.CategoryResponse {
	response := []web.CategoryResponse{}

	for _, value := range categories {
		response = append(response, *ToCategoryResponse(value))
	}

	return &response
}
