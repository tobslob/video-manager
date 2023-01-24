package token

import (
	"time"

	"github.com/google/uuid"
)

// Maker interface for managing tokens
type Maker interface {
	CreateToken(userId uuid.UUID, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
