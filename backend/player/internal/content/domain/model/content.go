package model

import (
	"errors"
	"github.com/sky0621/my-test-project/backend/player/internal/content/domain/model/entity"
	"github.com/sky0621/my-test-project/backend/shared/model"
)

func NewContentAggregate(m SaveContentWriteModel) (ContentAggregate, error) {
	if len(m.Programs) > 5 {
		return ContentAggregate{}, errors.New("too many programs")
	}

	programs := make([]entity.Program, len(m.Programs))
	for i, program := range m.Programs {
		programs[i] = entity.Program{
			ID:       program.ID,
			Question: program.Question,
			Answer:   program.Answer,
		}
	}

	return ContentAggregate{
		content: entity.Content{
			ID:   m.ID,
			Name: m.Name,
		},
		programs: programs,
	}, nil
}

type ContentAggregate struct {
	content  entity.Content
	programs []entity.Program
}

func (a ContentAggregate) GetID() model.ID {
	return a.content.ID
}

func (a ContentAggregate) GetName() string {
	return a.content.Name
}

func (a ContentAggregate) GetPrograms() []entity.Program {
	return a.programs
}
