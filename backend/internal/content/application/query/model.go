package query

import "github.com/sky0621/my-test-project/backend/internal/shared/model"

type ProgramReadModel struct {
	ID       model.ID
	Question string
	Answer   string
}
