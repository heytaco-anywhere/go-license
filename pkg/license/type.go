package license

import (
	"time"
)

type (
	// SigningData is the signing data.
	SigningData struct {
		Email     string    `json:"email"`
		ExpiredAt time.Time `json:"expiredAt"`
	}
)

// NewSigningData returns a new signing data.
// It would be enconded by base64.
func NewSigningData(email string, expiredAt time.Time) *SigningData {
	return &SigningData{
		Email:   email,
		ExpiredAt: expiredAt,
	}
}

func (s *SigningData) IsValid() bool {
	return s.ExpiredAt.IsZero()
}

func (s *SigningData) IsExpired(d time.Duration) bool {
	return s.ExpiredAt.
		Add(d).
		Before(time.Now())
}
