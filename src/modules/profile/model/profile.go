package model

import (
	"time"
)

// Profile ...initialize users profile
type Profile struct {
	ID        string    `bson: "id"`
	Name      string    `bson: "name"`
	CreatedAt time.Time `bson: "created_at"`
	UpdatedAt time.Time `bson: "updated_at"`
}

// Profiles ...
type Profiles []Profile
