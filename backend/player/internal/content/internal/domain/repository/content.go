package repository

import (
	"context"
	"github.com/sky0621/my-test-project/backend/player/internal/content/internal/domain/model"
)

type Content interface {
	Save(ctx context.Context, content model.ContentAggregate) error
}
