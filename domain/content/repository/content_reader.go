package repository

import (
	"context"

	"github.com/arvinpaundra/cent/content/domain/content/entity"
)

type ContentReader interface {
	FindActiveContentByUserId(ctx context.Context, userId int64) (*entity.Content, error)
}
