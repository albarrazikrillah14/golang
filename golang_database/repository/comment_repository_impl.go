package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{
		DB: db,
	}
}

func (repo commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repo.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)

	return comment, nil
}

func (repo commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, query, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		return comment, err
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repo commentRepositoryImpl) findAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []entity.Comment{}

	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
