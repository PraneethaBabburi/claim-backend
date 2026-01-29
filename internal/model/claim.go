package model

import "time"

const (
	StatusSubmitted   = "SUBMITTED"
	StatusUnderReview = "UNDER_REVIEW"
	StatusApproved    = "APPROVED"
	StatusDenied      = "DENIED"
)

type Claim struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
