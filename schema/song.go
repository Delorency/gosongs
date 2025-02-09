package schema

type Song struct {
	Id          string `json:"id"`
	Group       string `json:"group" validate:"required"`
	Song        string `json:"song" validate:"required"`
	Releasedate string `json:"releasedate" validate:"required"`
	Text        string `json:"text" validate:"required"`
	Link        string `json:"link" validate:"required"`
}
