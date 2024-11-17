package repository

import (
	"context"
	"database/sql"
	"errors"
	"medomeckz/category-restful-api/helper"
	"medomeckz/category-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRespository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO category(name) VALUES($1) RETURNING id"
	result, err := tx.QueryContext(ctx, query, category.Name)

	var id int

	if result.Next() {
		err := result.Scan(&id)
		helper.PanicIfError(err)
	}

	helper.PanicIfError(err)

	category.Id = int(id)
	return category

}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "UPDATE category SET name = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "DELETE FROM category WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicIfError(err)
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	query := "SELECT id, name FROM category WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err)

	defer rows.Close()
	category := domain.Category{}

	if rows.Next() {
		err = rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()
	categories := []domain.Category{}

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
