package repository

import (
	"context"

	"github.com/arvinpaundra/cent/content/domain/content/entity"
)

type ContentWriter interface {
	Save(ctx context.Context, content *entity.Content) error
}
