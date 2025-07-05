package router

import (
	"github.com/arvinpaundra/cent/content/application/rest/handler"
	"github.com/arvinpaundra/cent/content/application/rest/middleware"
	"github.com/arvinpaundra/cent/content/core/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type router struct {
	g   *gin.Engine
	db  *gorm.DB
	vld *validator.Validator
	hdl handler.Handler
}

func Register(g *gin.Engine, db *gorm.DB, vld *validator.Validator) {
	h := handler.NewHandler(db, vld)

	g.Use(middleware.Cors())
	g.Use(gin.Recovery())
	g.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}))

	r := router{g, db, vld, h}

	r.public()
	r.private()
}

func (r *router) public() {
	// v1 := r.g.Group("/api/v1")
}

func (r *router) private() {
	// v1 := r.g.Group("/api/v1")
}
