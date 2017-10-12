package models

type Episode struct {
	ID    int `json:"id"`
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Title             string  `json:"title,omitempty"`
	SerialID          *int    `json:"serialID,omitempty"`
	EpisodeOrder      string  `json:"episodeOrder,omitempty"`
	OriginalAirDate   string  `json:"originalAirDate,omitempty"`
	Runtime           string  `json:"runtime,omitempty"`
	UkViewersMM       float64 `json:"ukViewersMM,omitempty"`
	AppreciationIndex int     `json:"appreciationIndex,omitempty"`
	Missing           bool    `json:"missing,omitempty"`
	Recreated         bool    `json:"recreated,omitempty"`
}
