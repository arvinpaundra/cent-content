package model

import (
	"time"

	"github.com/guregu/null/v6"
)

type Campaign struct {
	ID            int64
	ContentId     int64
	TargetAmount  float64
	CurrentAmount float64
	IsActive      bool
	Text          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     null.Time
}
