package gostr

import "github.com/google/uuid"

func Guid() string {
	return uuid.New().String()
}
