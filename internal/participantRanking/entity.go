package participantRanking

import (
	"time"

	"gorm.io/gorm"
)

type ParticipantRanking struct {
	gorm.Model
	ID              string    `json:"id"`
	ParticipationId string    `json:"participation_id"`
	RankingId       int       `json:"ranking_id"`
	Grade           float64   `json:"grade"`
	Note            string    `json:"note"`
	Position        int       `json:"position"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}