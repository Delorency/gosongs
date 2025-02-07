package validator

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

// Функция обработки ошибок валидации
func HandleValidationErrors(w http.ResponseWriter, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		fields := []string{}

		for _, fieldErr := range validationErrors {
			fields = append(fields, fieldErr.Field())
		}

		response := map[string]interface{}{
			"error":  "Не все поля заполнены",
			"fields": fields,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	}
}
