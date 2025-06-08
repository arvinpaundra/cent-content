package model

import (
	"time"

	"github.com/guregu/null/v6"
)

type Content struct {
	ID          int64
	UserId      int64
	RingtoneUrl string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   null.Time
}
