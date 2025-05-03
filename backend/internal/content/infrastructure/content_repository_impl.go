package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sky0621/my-test-project/backend/internal/content/domain/model"
	"github.com/sky0621/my-test-project/backend/internal/content/domain/repository"
	"github.com/sky0621/my-test-project/backend/internal/content/infrastructure/rdb"
	"github.com/sky0621/my-test-project/backend/internal/shared/service"
)

var _ repository.Content = (*contentRepositoryImpl)(nil)

func NewContentRepository(db *sql.DB) repository.Content {
	return &contentRepositoryImpl{db: db}
}

type contentRepositoryImpl struct {
	db *sql.DB
}

func (c contentRepositoryImpl) Save(ctx context.Context, content model.ContentAggregate) error {
	if err := service.WithTransaction(ctx, c.db, func(tx *sql.Tx) error {
		q := rdb.New(tx)
		newContentID, err := content.GetID().MarshalBinary()
		if err != nil {
			return err
		}
		req := rdb.CreateContentsParams{
			ID:   newContentID,
			Name: content.GetName(),
		}
		if err := q.CreateContents(ctx, req); err != nil {
			return err
		}

		programs := content.GetPrograms()
		switch len(programs) {
		case 1:
			newProgramID1, err := programs[0].ID.MarshalBinary()
			if err != nil {
				return err
			}
			req1 := rdb.CreateProgramsBatch01Params{
				ID:        newProgramID1,
				Question:  programs[0].Question,
				Answer:    programs[0].Answer,
				ContentID: newContentID,
			}
			batch01, err := q.CreateProgramsBatch01(ctx, req1)
			if err != nil {
				return err
			}
			affected, err := batch01.RowsAffected()
			if err != nil {
				return err
			}
			if affected != 1 {
				return fmt.Errorf("expected 1 row affected, got %d", affected)
			}
		case 2:
			newProgramID1, err := programs[0].ID.MarshalBinary()
			if err != nil {
				return err
			}
			newProgramID2, err := programs[1].ID.MarshalBinary()
			if err != nil {
				return err
			}
			req2 := rdb.CreateProgramsBatch02Params{
				ID:          newProgramID1,
				Question:    programs[0].Question,
				Answer:      programs[0].Answer,
				ContentID:   newContentID,
				ID_2:        newProgramID2,
				Question_2:  programs[1].Question,
				Answer_2:    programs[1].Answer,
				ContentID_2: newContentID,
			}
			batch02, err := q.CreateProgramsBatch02(ctx, req2)
			if err != nil {
				return err
			}
			affected, err := batch02.RowsAffected()
			if err != nil {
				return err
			}
			if affected != 2 {
				return fmt.Errorf("expected 2 row affected, got %d", affected)
			}

			// TODO: 3, 4, 5
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
