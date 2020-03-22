package items

import "time"

// Item is a dummy element with user ownership. It just exists as example of use
// for authorization roles.
type Item struct {
	ID          int64     `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
