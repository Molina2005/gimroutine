// Importa la función createUser desde api.js,
// que se encarga de comunicarse con el backend
import { createUser } from "./api.js";
// Función que maneja la lógica de registro de usuarios.
// Recibe los datos del formulario y delega la creación al backend.
export function registerUser(formData){
    return createUser(formData)
    .then(data =>{
        console.log("usuario creado")
        return data
    });
} 