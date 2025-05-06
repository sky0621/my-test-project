package entity

import "github.com/sky0621/my-test-project/backend/shared/model"

type Program struct {
	ID       model.ID
	Question string
	Answer   string
}
