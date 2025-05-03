package query

import (
	"context"
	"github.com/sky0621/my-test-project/app/internal/shared/model"
)

type SearchContents interface {
	Exec(ctx context.Context, partialName *string) ([]SearchContentsReadModel, error)
}

type SearchContentsReadModel struct {
	ID       model.ID
	Name     string
	Programs []ProgramReadModel
}
