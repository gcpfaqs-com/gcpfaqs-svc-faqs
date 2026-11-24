package services

import (
	"gcpfaqs-svc-faqs/internal/domain/persistence"
	"gcpfaqs-svc-faqs/internal/domain/types"
)

type UpdateFaqUseCase struct {
	repository persistence.FaqRepository
}

func NewUpdateFaqUseCase(repo persistence.FaqRepository) *UpdateFaqUseCase {

	return &UpdateFaqUseCase{repository: repo}

}

func (uc *UpdateFaqUseCase) Do(id, title, content, poster string, tags []string, videoURL, videoPlaylistID, youtubeID string) error {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return err
	}

	faq, err := uc.repository.Get(validID)
	if err != nil {
		return err
	}

	validContent, err := types.NewContent(content)
	if err != nil {
		return err
	}

	validPoster, err := types.NewPoster(poster)
	if err != nil {
		return err
	}

	validTitle, err := types.NewTitle(title)
	if err != nil {
		return err
	}

	validTags, err := types.NewTags(tags)
	if err != nil {
		return err
	}

	validVideoURL, err := types.NewVideoURL(videoURL)
	if err != nil {
		return err
	}

	validVideoPlaylistID, err := types.NewVideoPlaylistID(videoPlaylistID)
	if err != nil {
		return err
	}

	validYoutubeID, err := types.NewYoutubeID(youtubeID)
	if err != nil {
		return err
	}

	faq.Title = validTitle
	faq.Content = validContent
	faq.Poster = validPoster
	faq.Tags = validTags
	faq.VideoURL = validVideoURL
	faq.VideoPlaylistID = validVideoPlaylistID
	faq.YoutubeID = validYoutubeID
	if err := uc.repository.Update(faq); err == nil {
		return err
	}

	return nil
}
