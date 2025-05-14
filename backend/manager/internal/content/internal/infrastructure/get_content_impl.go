package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sky0621/my-test-project/backend/manager/internal/content/internal/application/query"
	"github.com/sky0621/my-test-project/backend/manager/internal/content/internal/infrastructure/rdb"
	"github.com/sky0621/my-test-project/backend/shared/service"
)

var _ query.GetContent = (*getContentImpl)(nil)

func NewGetContent(db *sql.DB) query.GetContent {
	return &getContentImpl{db: db}
}

type getContentImpl struct {
	db *sql.DB
}

func (g getContentImpl) Exec(ctx context.Context, id string) (query.GetContentReadModel, error) {
	q := rdb.New(g.db)
	contentWithPrograms, err := q.GetContentWithProgramsById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return query.GetContentReadModel{}, nil
		}
		return query.GetContentReadModel{}, err
	}

	result := query.GetContentReadModel{}

	for i, row := range contentWithPrograms {
		if result.IsEmpty() {
			cID, err := service.ToID(row.ContentID)
			if err != nil {
				return query.GetContentReadModel{}, err
			}
			result.ID = cID
			result.Name = row.ContentName
			result.Programs = make([]query.ProgramReadModel, len(contentWithPrograms))
		}
		if row.ProgramID != nil && row.Question.Valid && row.Answer.Valid {
			pID, err := service.ToID(row.ProgramID)
			if err != nil {
				return query.GetContentReadModel{}, err
			}
			result.Programs[i] = query.ProgramReadModel{
				ID:       pID,
				Question: row.Question.String,
				Answer:   row.Answer.String,
			}
		}
	}

	return result, nil
}
