// Funcion para manejar la conexion con el backend por medio de fetch
// Recibe la data
export async function connectionUser(data){
    // Ruta para poder comunicarse con el backend(API) de creacion de usuarios
    const res = await fetch("/users",{
        // Metodo con el que se realizara la peticion
        method:"POST",
        // headers: indica el tipo de datos que se envian en el body
        // en este caso son datos en formato json
        headers:{"Content-Type":"application/json"},
        // Cambiar variables con los datos a cadena de texto 
        body:JSON.stringify(data)
    })
    // Si res que viene de fetch es diferente a ok genera el error si no pasa a retornar el res y lo pasa a json
    if (!res.ok){
        throw {status: res.status};
    }
    return res.json();
}
