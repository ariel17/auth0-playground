package items

import (
	"time"

	"github.com/google/uuid"
)

type newItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type item struct {
	ID          uuid.UUID  `json:"id"`
	UserID      string     `json:"user_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
