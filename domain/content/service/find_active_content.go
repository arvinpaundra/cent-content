package service

import (
	"context"

	contentres "github.com/arvinpaundra/cent/content/application/response/content"
	"github.com/arvinpaundra/cent/content/domain/content/repository"
)

type FindActiveContentHandler struct {
	contentReader repository.ContentReader
}

func NewFindActiveContentHandler(contentReader repository.ContentReader) FindActiveContentHandler {
	return FindActiveContentHandler{
		contentReader: contentReader,
	}
}

func (s FindActiveContentHandler) Handle(ctx context.Context, userId int64) (contentres.FindActiveContent, error) {
	content, err := s.contentReader.FindActiveContentByUserId(ctx, userId)
	if err != nil {
		return contentres.FindActiveContent{}, err
	}

	var campaignId int64

	if !content.Campaign.IsEmpty() {
		campaignId = content.Campaign.ID
	}

	res := contentres.FindActiveContent{
		ID:         content.ID,
		CampaignId: campaignId,
	}

	return res, nil
}
