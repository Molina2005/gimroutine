// Funcion para manejar la conexion con el backend por medio de fetch
// Recibe la data
export function connectionUser(data){
    // Ruta para poder comunicarse con el backend(API) de creacion de usuarios
    return fetch("/users",{
        // Metodo con el que se realizara la peticion
        method:"POST",
        // headers: indica el tipo de datos que se envian en el body
        // en este caso son datos en formato json
        headers:{"Content-Type":"application/json"},
        // Cambiar variables con los datos a cadena de texto 
        body:JSON.stringify(data)
    })
    // then: sirve para decir que hacer cuando una promesa se cumple
    // Response: respuesta del servidor 
    // Funcion para poder usar Response
    // Devuelve una promesa con los datos del json 
    .then(response =>{
        // Si response es difente a que todo salio bien(ok)
        // muestre el error, si no convierte JSON a objeto JS
        if (!response.ok ) {
            throw new Error("error al crear usuario");
        }
        return response.text();
    })
}
