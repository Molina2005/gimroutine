import { API_BACKEND_RAILWAY } from "../../config/config.js";
// Funcion para obtener roles creados 
export async function connectionAllRoles(){
    const res = await fetch(`${API_BACKEND_RAILWAY}/Allroles`,{
        method:"GET",
    })
    if (!res.ok){
        throw {status: res.status};
    }
    return res.json();
}