package domain

import "time"

type Visit struct {
	ID         uint32    `json:"id"` //unique
	LocationID uint32    `json:"location"`
	UserID     uint32    `json:"user"`
	VisitedAt  time.Time `json:"visited_at"`
	Mark       int       `json:"mark"` // 0-5
}
