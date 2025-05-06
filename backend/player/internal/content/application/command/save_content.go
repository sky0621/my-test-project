package command

import (
	"context"
	"github.com/sky0621/my-test-project/backend/player/internal/content/domain/model"
	"github.com/sky0621/my-test-project/backend/player/internal/content/domain/repository"
)

func NewSaveContent(r repository.Content) SaveContent {
	return SaveContent{r: r}
}

type SaveContent struct {
	r repository.Content
}

func (s SaveContent) Save(ctx context.Context, m model.SaveContentWriteModel) error {
	content, err := model.NewContentAggregate(m)
	if err != nil {
		return err
	}

	if err := s.r.Save(ctx, content); err != nil {
		return err
	}

	return nil
}
