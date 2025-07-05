package router

import (
	"github.com/arvinpaundra/cent/content/application/sse/handler"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Router struct {
	g  *gin.Engine
	db *gorm.DB
	nc *nats.Conn
}

func Register(g *gin.Engine, db *gorm.DB, nc *nats.Conn, logger *zap.Logger) {
	handler := handler.NewHandler(db, nc, logger)

	g.GET("/message", handler.ShowDonationMessageHandler)
}
