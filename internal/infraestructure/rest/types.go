package rest

type faqDTO struct {
	ID              string   `json:"id,omitempty"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	Poster          string   `json:"poster"`
	VideoURL        string   `json:"video_url"`
	VideoPlaylistID string   `json:"video_playlist_id"`
	Tags            []string `json:"tags"`
	YoutubeID       string   `json:"youtube_id"`
}
