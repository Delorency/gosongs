package schema

import "github.com/go-playground/validator"

var validate = validator.New()

type Song struct {
	Id           string `json:"id"`
	Group        string `json:"group" validate:"required"`
	Song         string `json:"song" validate:"required"`
	Release_date string `json:"release_date" validate:"required"`
	Text         string `json:"text" validate:"required"`
	Link         string `json:"link" validate:"required"`
}
