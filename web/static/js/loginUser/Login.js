import { login } from "./api.js"
export function logicLogin(formData){
    return login(formData)
    .then(data =>{
        console.log("Ingreso al sistema")
        return data
    })
    .catch(err =>{
        throw new Error("Credenciales invalidas", err);
    });
}