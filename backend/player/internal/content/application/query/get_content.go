package query

import (
	"context"
	"github.com/sky0621/my-test-project/backend/shared/model"
)

type GetContent interface {
	Exec(ctx context.Context, id string) (GetContentReadModel, error)
}

type GetContentReadModel struct {
	ID       model.ID
	Name     string
	Programs []ProgramReadModel
}

func (m GetContentReadModel) IsEmpty() bool {
	return m.ID.IsEmpty() && m.Name == ""
}
