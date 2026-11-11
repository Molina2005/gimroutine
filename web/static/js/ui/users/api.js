import { API_BACKEND_RAILWAY } from "../../config/config.js";
// Funcion para recibir informacion del backend eliminacion de usuarios
import { DataReturn } from "./search.js";
export async function ConnectionUsersDelete(id){
    const data = await fetch(`${API_BACKEND_RAILWAY}/users/${id}`, {
        method: "DELETE"
    })
    if (!data.ok){
        throw {status:data.status}
    }
    return data.json();
}

// Funcion para recibir informacion del backend del buscador de usuarios
export async function ConnectionSearchUser(search){
    try{
        const res = await fetch(`${API_BACKEND_RAILWAY}/UsersSearch/${encodeURIComponent(search)}`,{
            method: "GET",
        })
        const data = await res.json();
        DataReturn(data)
    }catch (err){
        throw {status:res.status}
    }
}