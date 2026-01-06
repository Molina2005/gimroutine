package exercises

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerExercises) HandlerUpdateInformationExercise(w http.ResponseWriter, r *http.Request) {
	// Verificacion de metodo
	if r.Method != http.MethodPut {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	// Probar si es mejor con busqueda por id o solo consultar y despues actualizar
	IdParam := chi.URLParam(r, "id")
	IdExercise, err := strconv.Atoi(IdParam)
	if err != nil {
		http.Error(w, "ID invalido", 400)
		return
	}
	// Lee la informacion que envia el usuario y lo organiza para poder usarla con formvalue
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "error al procesar el formulario", http.StatusBadRequest)
		return
	}
	// Se captura la informacion que envio el usuario por el formulario
	Id := r.FormValue("IdTypeOfExercise")
	IdTypeOfExercise, err := strconv.Atoi(Id)
	if err != nil {
		http.Error(w, "ID invalido", 400)
		return
	}
	name := r.FormValue("name")
	nameSpace := strings.TrimSpace(name)
	nameLower := strings.ToLower(nameSpace)
	description := r.FormValue("description")
	// Se obtienen los datos de la imagen
	file, fileHeader, err := r.FormFile("images")
	if err != nil {
		http.Error(w, "error al obtener imagen", 400)
		return
	}
	// Se cierra el archivo cuando finaliza la ejecución del handler
	defer file.Close()
	// ruta de creacion de carpeta en donde se guardaran las imagenes
	routeFile := `./uploadsImg`
	// creacion carpeta
	if err := os.MkdirAll(routeFile, os.ModePerm); err != nil {
		http.Error(w, "No se puede crear la carpeta", 500)
		return
	}
	// Se obtiene solo nombre de la imagen
	filename := filepath.Base(fileHeader.Filename)
	// une informacion de ruta de carpeta y nombre de la imagen
	route := filepath.Join(routeFile, filename)
	// creacion archivo en donde se guardara la data de la imagen
	data, err := os.Create(route)
	if err != nil {
		http.Error(w, "no se pudo guardar la imagen", 500)
		return
	}
	defer data.Close()
	// copia archivo que el usuario envio y lo deja en el archivo que se creo en el servidor
	// data : archivo creado en el servidor
	// file : archivo que el usuario envio en la peticion
	if _, err := io.Copy(data, file); err != nil {
		http.Error(w, "error al guardar la imagen", 500)
		return
	}
	if err := h.service.ServiceUpdateExercises(IdExercise, IdTypeOfExercise, nameLower, description, filename); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	w.Write([]byte("ejercicio actualizado correctamente"))
}
