import { login } from "./api.js"
export function logicLogin(formData){
    console.log("entro a logica login")
    return login(formData)
    .then(data =>{
        console.log("Ingreso al sistema")
        return data
    });
}