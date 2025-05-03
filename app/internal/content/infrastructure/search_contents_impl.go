package infrastructure

import (
	"context"
	"database/sql"
	"github.com/sky0621/my-test-project/app/internal/content/application/query"
	"github.com/sky0621/my-test-project/app/internal/content/infrastructure/rdb"
	"github.com/sky0621/my-test-project/app/internal/shared/converter"
	"github.com/sky0621/my-test-project/app/internal/shared/service"
)

var _ query.SearchContents = (*searchContentImpl)(nil)

func NewSearchContents(db *sql.DB) query.SearchContents {
	return &searchContentImpl{db: db}
}

type searchContentImpl struct {
	db *sql.DB
}

func (s searchContentImpl) Exec(ctx context.Context, partialName *string) ([]query.SearchContentsReadModel, error) {
	q := rdb.New(s.db)
	if partialName == nil {
		return list(ctx, q)
	} else {
		return search(ctx, q, *partialName)
	}
}

func list(ctx context.Context, q *rdb.Queries) ([]query.SearchContentsReadModel, error) {
	contentWithPrograms, err := q.ListContentsWithPrograms(ctx)
	if err != nil {
		return nil, err
	}

	// content_id をキーとするマップ
	resultMap := make(map[string]*query.SearchContentsReadModel)

	var results []*query.SearchContentsReadModel

	for _, content := range contentWithPrograms {
		if _, ok := resultMap[string(content.ContentID)]; !ok {
			cID, err := service.ToID(content.ContentID)
			if err != nil {
				return nil, err
			}
			resultMap[string(content.ContentID)] = &query.SearchContentsReadModel{
				ID:       cID,
				Name:     content.ContentName,
				Programs: []query.ProgramReadModel{},
			}
			results = append(results, resultMap[string(content.ContentID)])
		}

		if content.ProgramID != nil && content.Question.Valid && content.Answer.Valid {
			pID, err := service.ToID(content.ProgramID)
			if err != nil {
				return nil, err
			}
			resultMap[string(content.ContentID)].Programs = append(resultMap[string(content.ContentID)].Programs, query.ProgramReadModel{
				ID:       pID,
				Question: content.Question.String,
				Answer:   content.Answer.String,
			})
		}
	}

	return converter.ToVals(results), nil
}

func search(ctx context.Context, q *rdb.Queries, partialName string) ([]query.SearchContentsReadModel, error) {
	contentWithPrograms, err := q.SearchContentsWithPrograms(ctx, partialName)
	if err != nil {
		return nil, err
	}

	// content_id をキーとするマップ
	resultMap := make(map[string]*query.SearchContentsReadModel)

	var results []*query.SearchContentsReadModel

	for _, content := range contentWithPrograms {
		if _, ok := resultMap[string(content.ContentID)]; !ok {
			cID, err := service.ToID(content.ContentID)
			if err != nil {
				return nil, err
			}
			resultMap[string(content.ContentID)] = &query.SearchContentsReadModel{
				ID:       cID,
				Name:     content.ContentName,
				Programs: []query.ProgramReadModel{},
			}
			results = append(results, resultMap[string(content.ContentID)])
		}

		if content.ProgramID != nil && content.Question.Valid && content.Answer.Valid {
			pID, err := service.ToID(content.ProgramID)
			if err != nil {
				return nil, err
			}
			resultMap[string(content.ContentID)].Programs = append(resultMap[string(content.ContentID)].Programs, query.ProgramReadModel{
				ID:       pID,
				Question: content.Question.String,
				Answer:   content.Answer.String,
			})
		}
	}

	return converter.ToVals(results), nil
}
