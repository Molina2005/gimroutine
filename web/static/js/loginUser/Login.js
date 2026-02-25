import { login } from "./api.js"
export async function logicLogin(formData){
    const data = await login(formData);
    if (data.error){
        throw new Error(data.message || "Credenciales invalidas");
    }
    return data;
}