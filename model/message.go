package model

import (
	"time"

	"github.com/guregu/null/v6"
)

type Message struct {
	ID          int64
	ContentId   int64
	BgColor     string
	TextColor   string
	IsTtsActive bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   null.Time
}
