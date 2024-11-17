package repository

import (
	"context"
	"database/sql"
	"medomeckz/category-restful-api/common/exceptions"
	"medomeckz/category-restful-api/helper"
	"medomeckz/category-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO categories(name) VALUES($1) RETURNING id"
	result, err := tx.QueryContext(ctx, query, category.Name)
	helper.PanicIfError(err)

	id := 0
	if result.Next() {
		err = result.Scan(&id)
	}
	helper.PanicIfError(err)

	category.Id = int(id)

	return category
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "UPDATE categories SET name = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	if err != nil {
		panic(exceptions.NotFoundError("gagal menambahkan kategori"))
	}
	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, category.Id)
	if err != nil {
		panic(exceptions.NotFoundError("gagal mengubah kategori"))
	}
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Category {
	query := "SELECT id, name FROM categories WHERE id = $1"
	result, err := tx.QueryContext(ctx, query, categoryId)
	if err != nil {
		panic(exceptions.InvariantError("gagal mendapatkan kategori"))
	}

	category := domain.Category{}
	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category

	} else {
		panic(exceptions.NotFoundError("kategory tidak ditemukan"))
	}

}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT id, name FROM categories"
	result, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	categories := []domain.Category{}
	for result.Next() {
		category := domain.Category{}

		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
