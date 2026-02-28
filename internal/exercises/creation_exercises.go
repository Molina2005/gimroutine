package exercises

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (h *HandlerExercises) HandlerCreationExercises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "metodo_no_permitido",
		})
		return
	}
	// Lee el contenido que el usuario envio (parsea peticion usuario)
	// Organizado los datos para poder acceder a ellos por medio de r.FormValue y r.FormFile
	// Usa hasta 10MB de memoria RAM
	// Si esta entre los 10MB se mantiene en memoria temporalmente
	// Si los archivos superan ese tamaño, se guardan en archivos temporales en disco
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "error_al_procesar_formulario",
		})
		return
	}
	// Captura la informacion que el usuario envia en el formulario
	IdTypeOfExercise := r.FormValue("IdTypeOfExercise")
	id, err := strconv.Atoi(IdTypeOfExercise)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id_invalido",
		})
		return
	}
	name := r.FormValue("name")
	// Manejo de espacios y mayusculas para el nombre del ejercicio
	nameSpace := strings.TrimSpace(name)
	nameLower := strings.ToLower(nameSpace)
	description := r.FormValue("description")
	// devuelve: archivo en stream para leerlo pedazo a pedazo, datos del archivo(nombre, tamaño, etc), error
	file, fileHeader, err := r.FormFile("images")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "error_obtener_imagen",
		})
		return
	}
	// Cierra la conexion cuando termina su funcion
	defer file.Close()
	// Creacion de carpeta automatica para guardar imagenes
	routeFile := `./uploadsImg`
	if err := os.MkdirAll(routeFile, os.ModePerm); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "No_se_puede_crear_la_carpeta",
		})
		return
	}
	// Solo toma solo el nombre de la imagen sin importar sus direcciones
	filename := filepath.Base(fileHeader.Filename)
	// Une la informacion de la ruta del archivo y el nombre de la imagen
	route := filepath.Join(routeFile, filename)
	// Crea una ruta en la cual se guarda el nombre del archivo de forma permanente
	// Crea el archivo físico dentro de la carpeta uploadsImg usando la ruta definida
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
	// pase de servicio con toda la informacion
	if err := h.service.ServiceCreationExercises(id, nameLower, description, filename); err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ejercicio_creado_correctamente",
	})
}
