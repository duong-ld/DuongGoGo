package userdto

import (
	"github.com/google/uuid"
	"time"
)

type UserResponse struct {
	ID       uuid.UUID
	Email    string
	BirthDay time.Time
}
