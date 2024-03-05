package dtos

type MusixmatchResponse struct {
	Message Message `json:"message"`
}

type Track struct {
	TrackID   int    `json:"track_id"`
	TrackName string `json:"track_name"`
	Artist    string `json:"artist_name"`
}

type TrackList struct {
	Track Track `json:"track"`
}

type Body struct {
	TrackList []TrackList `json:"track_list"`
}

type Message struct {
	Body Body `json:"body"`
}

//

type Artist struct {
	ArtistID   int    `json:"artist_id"`
	ArtistName string `json:"artist_name"`
	ArtistURL  string `json:"artist_share_url"`
	Bio        string `json:"artist_bio"`
}

type ArtistBody struct {
	Artist Artist `json:"artist"`
}

type ArtistMessage struct {
	Body ArtistBody `json:"body"`
}

type ArtistResponse struct {
	Message ArtistMessage `json:"message"`
}
