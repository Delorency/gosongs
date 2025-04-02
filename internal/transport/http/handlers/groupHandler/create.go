package grouphandler

import (
	"encoding/json"
	"main/internal/models"
	dto "main/internal/transport/http/dto"
	e "main/internal/transport/http/httperror"
	v "main/internal/validator"
	"net/http"

	"github.com/go-playground/validator"
)

type requestCreate struct {
	Name string `json:"name" validate:"required"`
}

func (gh *groupHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req requestCreate

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.NewResponse(
			e.NewError(""),
			http.StatusInternalServerError,
			w,
		)
		return
	}

	if err := validator.New().Struct(req); err != nil {
		data, f := v.HandleValidationErrors(w, err)
		if f {
			dto.NewResponse(
				data,
				http.StatusBadRequest,
				w,
			)
			return
		}
	}

	obj := models.Group{Name: req.Name}

	if err := gh.service.Create(&obj); err != nil {
		dto.NewResponse(
			e.NewError("Такая группа уже была создана"),
			http.StatusBadRequest,
			w,
		)
		return
	}

	dto.NewResponse(
		obj,
		http.StatusOK,
		w,
	)
}
