package util

import (
	"strings"

	"github.com/google/uuid"
)

func GetUniqueID() string {
	uuidWithHyphen := uuid.New()
	return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
}
