package users

import "time"

type user struct {
	ID    string `json:"id"`
	Email struct {
		Address    string `json:"address"`
		IsVerified bool   `json:"is_verified"`
	} `json:"email"`
	Nickname    string    `json:"nickname"`
	GivenName   string    `json:"given_name"`
	FamilyName  string    `json:"family_name"`
	Groups      []string  `json:"groups"`
	Permissions []string  `json:"permissions"`
	Roles       []string  `json:"roles"`
	CreatedAt   time.Time `json:"created_at"`
	AvatarURL   string    `json:"avatar_url"`
	Enabled     bool      `json:"enabled"`
}
