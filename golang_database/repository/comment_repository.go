package repository

import (
	"context"
	"golang-database/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	findAll(ctx context.Context) ([]entity.Comment, error)
}
