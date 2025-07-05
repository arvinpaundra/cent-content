package handler

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	db     *gorm.DB
	nc     *nats.Conn
	logger *zap.Logger
}

func NewHandler(db *gorm.DB, nc *nats.Conn, logger *zap.Logger) Handler {
	return Handler{
		db:     db,
		nc:     nc,
		logger: logger,
	}
}
