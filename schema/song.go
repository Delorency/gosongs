package schema

type Song struct {
	Id          string `json:"id"`
	Group       string `json:"group" validate:"required"`
	Song        string `json:"song" validate:"required"`
	Releasedate string `json:"releasedate" validate:"required"`
	Text        string `json:"text" validate:"required"`
	Link        string `json:"link" validate:"required"`
}
type SaveSongInput struct {
	Group string `json:"group" swagger:"required"`
	Song  string `json:"song" swagger:"required"`
}
type UpdateSong struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	Releasedate string `json:"releasedate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
