package service

import (
	"github.com/google/uuid"
	"github.com/sky0621/my-test-project/backend/shared/model"
)

func MustCreateNewID() model.ID {
	u, e := CreateNewID()
	if e != nil {
		panic(e)
	}
	return u
}

func CreateNewID() (model.ID, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return model.ID{}, err
	}
	return model.NewID(uuidV7), nil
}

func ParseID(s string) (model.ID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return model.ID{}, err
	}
	return model.NewID(id), nil
}

func ToID(id []byte) (model.ID, error) {
	uuidID, err := uuid.FromBytes(id)
	if err != nil {
		return model.ID{}, err
	}
	return model.NewID(uuidID), nil
}
