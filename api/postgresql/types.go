package postgresql

import (
	"time"
)

// Migration defines applied database migration.
type Migration struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	AppliedBy string    `json:"applied_by"`
	Created   time.Time `json:"created"`
}
