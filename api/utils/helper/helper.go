package helper

import uuid "github.com/satori/go.uuid"

func GetIdFromString(idString string) uuid.UUID {
	id, _ := uuid.FromString(idString)
	return id
}
