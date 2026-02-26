// funcion para cargar menus reutilizables en otro archivo html
function cargarMenu(nombreArchivo, idContenedor = 'container-nav'){
    console.log("JS CARGADO")
    // archivo el cual contiene la data  
    fetch(nombreArchivo) 
    // Promesa que convierte data que llega en texto para poder utilizarlo
    // Tiene que cumplirse el primer then para poder avanzar al siguiente
    .then(res => res.text())
    // resultado de la promesa anterior
    .then(data =>{
        // Se intercambia lo que esta dentro de idContenedor (lo que esta dentro de container-nav) por lo que esta dentro de data(lo que esta dentro del archivo cargado)
        document.getElementById(idContenedor).innerHTML = data
    });
}


