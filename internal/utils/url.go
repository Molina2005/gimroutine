package utils

import (
	"net/http"
)

var Domins = map[string]bool{
	// ENTORNO LOCAL
	"http://localhost:3000": true,
	// ENTORNO PRUEBAS
	"https://mabtrosoftweb-git-testing-programming10.vercel.app": true,
	// ENTORNO PRODUCCION
	"https://mabtrosoftweb.vercel.app": true,
}

// Funcion para poder manejar dominio en el que se encuentra trabajando
func GetFrontEndtUrl(url *http.Request) string {
	// SE OBTIENE ENCABEZADOS DE LA PETICION DEL CLIENTE
	// EN ESTE CASO PARA EL MANEJO DE QUE DOMINIO SE UTILIZO
	origin := url.Header.Get("Origin")
	if Domins[origin] {
		return ""
	}
	return origin
}
