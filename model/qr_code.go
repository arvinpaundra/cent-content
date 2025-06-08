package model

import (
	"time"

	"github.com/guregu/null/v6"
)

type QrCode struct {
	ID        int64
	ContentId int64
	Code      string
	BgColor   string
	QrColor   string
	Text      null.String
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
