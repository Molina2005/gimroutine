package exercises

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerExercises) HandlerUpdateInformationExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Verificacion de metodo
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	// Probar si es mejor con busqueda por id o solo consultar y despues actualizar
	IdParam := chi.URLParam(r, "id")
	IdExercise, err := strconv.Atoi(IdParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id_invalido",
		})
		return
	}
	// Lee la informacion que envia el usuario y lo organiza para poder usarla con formvalue
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "error_al_procesar_formulario",
		})
		return
	}
	// Se captura la informacion que envio el usuario por el formulario
	Id := r.FormValue("IdTypeOfExercise")
	IdTypeOfExercise, err := strconv.Atoi(Id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id_invalido",
		})
		return
	}
	name := r.FormValue("name")
	nameSpace := strings.TrimSpace(name)
	nameLower := strings.ToLower(nameSpace)
	description := r.FormValue("description")
	// variable que guarda el nombre del archivo
	var filename string
	// Se obtienen los datos de la imagen
	file, fileHeader, err := r.FormFile("images")
	if err != nil {
		// ErrMissingFile : maneja para que imagen se pueda utilizar de manera opcional
		// Si el error es == a que no se envio el archivo lo envia directamente a la funcion del servicio
		// Si es != a otro error pero que no es el de que no se envio archivo, entra directmente al error real
		if err != http.ErrMissingFile {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "error_obtener_imagen",
			})
			return
		}
	} else {
		// Se cierra el archivo cuando finaliza la ejecución del handler
		defer file.Close()
		// ruta de creacion de carpeta en donde se guardaran las imagenes
		routeFile := `./uploadsImg`
		// creacion carpeta
		if err := os.MkdirAll(routeFile, os.ModePerm); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "No_se_puede_crear_la_carpeta",
			})
			return
		}
		// Se obtiene solo nombre de la imagen
		filename = filepath.Base(fileHeader.Filename)
		// une informacion de ruta de carpeta y nombre de la imagen
		route := filepath.Join(routeFile, filename)
		// creacion archivo en donde se guardara la data de la imagen
		data, err := os.Create(route)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "no_se_pudo_guardar_la_imagen",
			})
			return
		}
		defer data.Close()
		// copia archivo que el usuario envio y lo deja en el archivo que se creo en el servidor
		// data : archivo creado en el servidor
		// file : archivo que el usuario envio en la peticion
		if _, err := io.Copy(data, file); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "error_al_guardar_imagen",
			})
			return
		}
	}
	// Funcion con logica del servicio
	if err := h.service.ServiceUpdateExercises(IdExercise, IdTypeOfExercise, nameLower, description, filename); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ejercicio_actualizado_correctamente",
	})
}
