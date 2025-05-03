package model

import "github.com/google/uuid"

func NewID(id uuid.UUID) ID {
	return ID{id: id}
}

type ID struct {
	id uuid.UUID
}

func (i ID) Value() uuid.UUID {
	return i.id
}

func (i ID) String() string {
	return i.Value().String()
}

func (i ID) MarshalBinary() ([]byte, error) {
	return i.Value().MarshalBinary()
}

func (i ID) IsEmpty() bool {
	return i.Value() == uuid.Nil
}
