package rest

import (
	"encoding/json"
	"gcpfaqs-svc-faqs/internal/domain/persistence"
	"gcpfaqs-svc-faqs/internal/domain/services"
	"net/http"
)

type RESTUpdateUseCase struct {
	repository persistence.FaqRepository
}

func NewRESTUpdateUseCase(repository persistence.FaqRepository) *RESTUpdateUseCase {
	return &RESTUpdateUseCase{
		repository: repository,
	}
}

func (uc *RESTUpdateUseCase) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	var dto faqDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updateUC := services.NewUpdateFaqUseCase(uc.repository)
	err := updateUC.Do(idStr, dto.Title, dto.Content, dto.Poster, dto.Tags, dto.VideoURL, dto.VideoPlaylistID, dto.YoutubeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto.ID = idStr
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto)
}
