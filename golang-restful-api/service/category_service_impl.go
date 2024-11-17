package service

import (
	"context"
	"database/sql"
	"medomeckz/category-restful-api/helper"
	"medomeckz/category-restful-api/model/domain"
	"medomeckz/category-restful-api/model/web"
	"medomeckz/category-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
	DB         *sql.DB
	Validate   validator.Validate
}

func NewCategoryService(repo repository.CategoryRepository, db *sql.DB, validate validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		Repository: repo,
		DB:         db,
		Validate:   validate,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, r web.CategoyCreateRequest) web.CategoryResponse {
	err := s.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	domain := domain.Category{
		Name: r.Name,
	}
	category := s.Repository.Save(ctx, tx, domain)

	return *helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Update(ctx context.Context, r web.CategoryUpdateRequset) web.CategoryResponse {
	err := s.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result := s.FindById(ctx, r.Id)
	domain := domain.Category{
		Id:   result.Id,
		Name: r.Name,
	}

	response := s.Repository.Update(ctx, tx, domain)

	return *helper.ToCategoryResponse(response)
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result := s.FindById(ctx, categoryId)

	domain := domain.Category{
		Id:   result.Id,
		Name: result.Name,
	}
	s.Repository.Delete(ctx, tx, domain)
}

func (s *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := s.Repository.FindById(ctx, tx, categoryId)

	return *helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	response := s.Repository.FindAll(ctx, tx)

	return *helper.ToCategoriesResponse(response)
}
