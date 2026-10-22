package servicepoints

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (h *HandlerServicePoints) HandlerServicePointsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	var input struct {
		Name        string `json:"name"`
		Addres      string `json:"addres"`
		City        string `json:"city"`
		Id_Regional int    `json:"id_regional"`
		Id_State    int    `json:"id_state"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json_invalido",
		})
		return
	}
	if err := h.service.ServiceServicePointsCreation(input.Name, input.Addres, input.City, input.Id_Regional, input.Id_State); err != nil {
		if errors.Is(err, ErrServicePointsAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "punto_servicio_creado_correctamente",
	})

}
