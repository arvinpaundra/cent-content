package service

import (
	"context"

	contentres "github.com/arvinpaundra/cent/content/application/response/content"
	"github.com/arvinpaundra/cent/content/domain/content/repository"
)

type FindActiveContent struct {
	contentReader repository.ContentReader
}

func NewFindActiveContent(contentReader repository.ContentReader) FindActiveContent {
	return FindActiveContent{
		contentReader: contentReader,
	}
}

func (s FindActiveContent) Exec(ctx context.Context, userId int64) (contentres.FindActiveContent, error) {
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
