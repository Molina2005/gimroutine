import { login } from "./api.js"

export async function logicLogin(formData){
    const data = await login(formData);
    console.log(data)
    return data;
}

