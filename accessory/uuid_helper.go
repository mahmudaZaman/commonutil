package accessory

import (
	"github.com/google/uuid"
	"strings"
)

func getUUIDWithoutDash() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
