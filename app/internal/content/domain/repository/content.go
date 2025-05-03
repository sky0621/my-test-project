package repository

import (
	"context"
	"github.com/sky0621/my-test-project/app/internal/content/domain/model"
)

type Content interface {
	Save(ctx context.Context, content model.ContentAggregate) error
}
