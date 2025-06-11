package grpc

import (
	contentgrpc "github.com/arvinpaundra/cent/content/application/grpc/content"
	"github.com/arvinpaundra/cent/content/core/validator"
	"github.com/arvinpaundra/centpb/gen/go/content/v1"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func Register(srv *grpc.Server, db *gorm.DB, vld *validator.Validator) {
	content.RegisterContentServiceServer(srv, contentgrpc.NewContentService(db))
}
