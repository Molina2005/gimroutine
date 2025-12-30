package exercises

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (h *HandlerExercises) HandlerCreationExercises(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	// Lee el contenido que el usuario envio (parsea peticion usuario)
	// Organizado los datos para poder acceder a ellos por medio de r.FormValue y r.FormFile
	// Usa hasta 10MB de memoria RAM
	// Si esta entre los 10MB se mantiene en memoria temporalmente
	// Si los archivos superan ese tamaño, se guardan en archivos temporales en disco
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "error al procesar el formulario", http.StatusBadRequest) // Solicitud estado incorrecta
		return
	}
	// Captura la informacion que el usuario envia en el formulario
	IdTypeOfExercise := r.FormValue("idTypeOfExercise")
	id, err := strconv.Atoi(IdTypeOfExercise)
	if err != nil {
		http.Error(w, "ID invalido", 400)
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
		http.Error(w, "error al obtener imagen", 400)
		return
	}
	// Cierra la conexion cuando termina su funcion
	defer file.Close()
	// Creacion de carpeta automatica para guardar imagenes
	routeFile := `c:\Users\CRISTIAN MOLINA\go\src/gimroutine/uploadsImg`
	if err := os.MkdirAll(routeFile, os.ModePerm); err != nil {
		http.Error(w, "No se puede crear la carpeta", 500)
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
		http.Error(w, "no se pudo guardar la imagen", 500)
		return
	}
	defer data.Close()
	// copia archivo que el usuario envio y lo deja en el archivo que se creo en el servidor
	// data : archivo creado en el servidor
	// file : archivo que el usuario envio en la peticion
	if _, err := io.Copy(data, file); err != nil {
		http.Error(w, "Error al guardar la imagen", 500)
	}
	// pase de servicio con toda la informacion
	if err := h.service.ServiceCreationExercises(id, nameLower, description, filename); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte("ejercicio creado correctamente"))
}
