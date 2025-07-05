package handler

import (
	"github.com/arvinpaundra/cent/content/core/validator"
	"gorm.io/gorm"
)

type Handler struct {
	db  *gorm.DB
	vld *validator.Validator
}

func NewHandler(db *gorm.DB, vld *validator.Validator) Handler {
	return Handler{
		db:  db,
		vld: vld,
	}
}
