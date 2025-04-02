package grouphandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"main/internal/models"
	v "main/internal/validator"

	dto "main/internal/transport/http/dto"
	e "main/internal/transport/http/httperror"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
)

type requestUpdate struct {
	Name string `json:"name" validate:"required"`
}

func (gh *groupHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req requestUpdate

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		dto.NewResponse(
			e.NewError("Идентификатор должен быть числом"),
			http.StatusNotFound,
			w,
		)
		return
	}

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

	data := models.Group{Name: req.Name}

	obj, err := gh.service.Update(uint(id), &data)
	if err != nil {
		dto.NewResponse(
			e.NewError("Такая группа уже была создана"),
			http.StatusBadRequest,
			w,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}
