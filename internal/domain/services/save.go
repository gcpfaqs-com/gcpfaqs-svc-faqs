package services

import (
	"gcpfaqs-svc-faqs/internal/domain/models"
	"gcpfaqs-svc-faqs/internal/domain/persistence"
)

type SaveFaqUseCase struct {
	repository persistence.FaqRepository
}

func NewSaveFaqUseCase(repo persistence.FaqRepository) *SaveFaqUseCase {

	return &SaveFaqUseCase{repository: repo}

}

func (uc *SaveFaqUseCase) Do(title, content, poster string, tags []string, videoURL, videoPlaylistID string, youtubeID string) (string, error) {
	faq, err := models.NewFaq(title, content, poster, tags, videoURL, videoPlaylistID, youtubeID)
	if err != nil {
		return "", err
	}

	err = uc.repository.Save(faq)
	if err != nil {
		return "", err
	}

	return faq.ID.Value(), nil
}
