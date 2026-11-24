package rest

import (
	"encoding/json"
	"gcpfaqs-svc-faqs/internal/domain/persistence"
	"gcpfaqs-svc-faqs/internal/domain/services"
	"net/http"
)

type RESTGetUseCase struct {
	repository persistence.FaqRepository
}

func NewRESTGetUseCase(repository persistence.FaqRepository) *RESTGetUseCase {
	return &RESTGetUseCase{
		repository: repository,
	}
}

func (uc *RESTGetUseCase) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	getUC := services.NewGetFaqUseCase(uc.repository)
	faq, err := getUC.Do(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := faqDTO{
		ID:              faq.ID.Value(),
		Title:           faq.Title.Value(),
		Content:         faq.Content.Value(),
		Poster:          faq.Poster.Value(),
		VideoURL:        faq.VideoURL.Value(),
		VideoPlaylistID: faq.VideoPlaylistID.Value(),
		Tags:            faq.Tags.Value(),
		YoutubeID:       faq.YoutubeID.Value(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
