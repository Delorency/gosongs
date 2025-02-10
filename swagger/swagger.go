package swagger

type Input struct {
	Group string `json:"group" swagger:"required"`
	Song  string `json:"song" swagger:"required"`
}

type SongOutput struct {
	Id          string `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	Releasedate string `json:"releasedate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
