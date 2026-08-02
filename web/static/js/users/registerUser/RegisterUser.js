// Importa la función createUser desde api.js,
// que se encarga de comunicarse con el backend
import { connectionUser } from "./api.js";
// Función que maneja la lógica de registro de usuarios.
// Recibe los datos del formulario y delega la creación al backend.
export async function logicRegisterUser(formData){
    // Recibe la informacion que envia ConnectionUser si es correcta
    const data = await connectionUser(formData);
    console.log(data)
    return data;
} 