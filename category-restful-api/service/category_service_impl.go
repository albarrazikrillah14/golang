package service

import (
	"context"
	"database/sql"
	"medomeckz/category-restful-api/exception"
	"medomeckz/category-restful-api/helper"
	"medomeckz/category-restful-api/model/domain"
	"medomeckz/category-restful-api/model/web"
	"medomeckz/category-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}
	category = s.Repository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.Repository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name
	s.Repository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.Repository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	s.Repository.Delete(ctx, tx, category)
}

func (s *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.Repository.FindById(ctx, tx, categoryId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := s.Repository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
