package content

import (
	"context"

	"github.com/arvinpaundra/cent/content/domain/content/service"
	contentinfra "github.com/arvinpaundra/cent/content/infrastructure/content"
	"github.com/arvinpaundra/centpb/gen/go/content/v1"
	"gorm.io/gorm"
)

type ContentService struct {
	db *gorm.DB

	content.UnimplementedContentServiceServer
}

func NewContentService(db *gorm.DB) ContentService {
	return ContentService{db: db}
}

func (c ContentService) FindActiveContent(ctx context.Context, req *content.FindActiveContentRequest) (*content.FindActiveContentResponse, error) {
	userId := req.GetUserId()

	handler := service.NewFindActiveContentHandler(
		contentinfra.NewContentReaderRepository(c.db),
	)

	result, err := handler.Handle(ctx, userId)
	if err != nil {
		return nil, err
	}

	res := content.FindActiveContentResponse{
		Id:         result.ID,
		CampaignId: result.CampaignId,
	}

	return &res, nil
}
