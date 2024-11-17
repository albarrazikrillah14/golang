package service

import (
	"context"
	"medomeckz/category-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, r web.CategoyCreateRequest) web.CategoryResponse
	Update(ctx context.Context, r web.CategoryUpdateRequset) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
