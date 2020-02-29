package libutil

import (
	"github.com/google/uuid"
)

func Uuid() string {
	uuidObj := uuid.New()
	return uuidObj.String()
}
