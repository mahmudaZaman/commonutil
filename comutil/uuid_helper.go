package comutil

import (
	"github.com/google/uuid"
	"strings"
)

// GetUUIDWithoutDash returns a guid without dash.
func GetUUIDWithoutDash() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
