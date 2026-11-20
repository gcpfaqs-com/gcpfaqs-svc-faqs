package models

import "gcpfaqs-svc-faqs/internal/domain/types"

type Faq struct {
	ID              *types.ID              `json:"id"`
	Title           *types.Title           `json:"title"`
	Content         *types.Content         `json:"content"`
	Poster          *types.Poster          `json:"poster"`
	VideoURL        *types.VideoURL        `json:"videoURL"`
	VideoPlaylistID *types.VideoPlaylistID `json:"videoPlaylistID"`
	Tags            *types.Tags            `json:"tags"`
	YoutubeID       *types.YoutubeID       `json:"youtubeID"`
}

func NewFaq(title string, content string, poster string, tags []string, videoURL string, videoPlaylistID string, youtubeID string) (*Faq, error) {
	faqTitle, err := types.NewTitle(title)
	if err != nil {
		return nil, err
	}

	faqContent, err := types.NewContent(content)
	if err != nil {
		return nil, err
	}

	faqPoster, err := types.NewPoster(poster)
	if err != nil {
		return nil, err
	}

	faqTags, err := types.NewTags(tags)
	if err != nil {
		return nil, err
	}

	faqVideoURL, err := types.NewVideoURL(videoURL)
	if err != nil {
		return nil, err
	}

	faqVideoPlaylistID, err := types.NewVideoPlaylistID(videoPlaylistID)
	if err != nil {
		return nil, err
	}

	faqID, err := types.NewID()
	if err != nil {
		return nil, err
	}

	faqYoutubeID, err := types.NewYoutubeID(youtubeID)
	if err != nil {
		return nil, err
	}

	return &Faq{
		ID:              faqID,
		Title:           faqTitle,
		Content:         faqContent,
		Poster:          faqPoster,
		VideoURL:        faqVideoURL,
		VideoPlaylistID: faqVideoPlaylistID,
		Tags:            faqTags,
		YoutubeID:       faqYoutubeID,
	}, nil
}
