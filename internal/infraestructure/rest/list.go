package rest

import (
	"encoding/json"
	"gcpfaqs-svc-faqs/internal/domain/persistence"
	"gcpfaqs-svc-faqs/internal/domain/services"
	"net/http"
)

type RESTListUseCase struct {
	repository persistence.FaqRepository
}

func NewRESTListUseCase(repository persistence.FaqRepository) *RESTListUseCase {
	return &RESTListUseCase{
		repository: repository,
	}
}

func (uc *RESTListUseCase) List(w http.ResponseWriter, r *http.Request) {
	listUC := services.NewListFaqUseCase(uc.repository)
	faqs, err := listUC.Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := make([]faqDTO, 0, len(faqs))
	for _, a := range faqs {
		response = append(response, faqDTO{
			ID:              a.ID.Value(),
			Title:           a.Title.Value(),
			Content:         a.Content.Value(),
			Poster:          a.Poster.Value(),
			VideoURL:        a.VideoURL.Value(),
			VideoPlaylistID: a.VideoPlaylistID.Value(),
			Tags:            a.Tags.Value(),
			YoutubeID:       a.YoutubeID.Value(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
