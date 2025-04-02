package grouphandler

import (
	"main/internal/schemes"
	dto "main/internal/transport/http/dto"
	e "main/internal/transport/http/httperror"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (gh *groupHandler) List(w http.ResponseWriter, r *http.Request) {
	var limit, page int
	var err error

	if r.URL.Query().Get("limit") == "" && r.URL.Query().Get("page") == "" {
		limit = -1
		page = 1
	} else {
		limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil || limit <= 0 {
			dto.NewResponse(
				e.NewError("limit должно быть числом > 0"),
				http.StatusBadRequest,
				w,
			)
			return
		}

		page, err = strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page <= 0 {
			dto.NewResponse(
				e.NewError("page должно быть числом > 0"),
				http.StatusBadRequest,
				w,
			)
			return
		}
	}

	pag := schemes.Pagination{Limit: limit, Page: page}

	groups, err := gh.service.List(&pag)

	if err != nil {
		dto.NewResponse(
			e.NewError("Ошибка получения данных"),
			http.StatusInternalServerError,
			w,
		)
	}

	dto.NewResponse(
		groups,
		http.StatusOK,
		w,
	)

}

func (gh *groupHandler) Retireve(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		dto.NewResponse(
			e.NewError("Идентификатор должен быть числом"),
			http.StatusNotFound,
			w,
		)
		return
	}

	group, err := gh.service.Retrieve(uint(id))

	if err != nil {
		dto.NewResponse(
			e.NewError("Ошибка получения данных"),
			http.StatusInternalServerError,
			w,
		)
	}

	dto.NewResponse(
		group,
		http.StatusOK,
		w,
	)
}
