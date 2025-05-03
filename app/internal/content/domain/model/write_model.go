package model

import (
	"github.com/sky0621/my-test-project/app/internal/shared/model"
)

type SaveContentWriteModel struct {
	ID       model.ID
	Name     string
	Programs []ProgramWriteModel
}

type ProgramWriteModel struct {
	ID       model.ID
	Question string
	Answer   string
}
