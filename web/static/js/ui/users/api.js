// Funcion para recibir informacion del backend eliminacion de usuarios
import { DataReturn } from "./search.js";
export async function ConnectionUsersDelete(id){
    const data = await fetch(`/users/${id}`, {
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
        const res = await fetch(`/UsersSearch/${encodeURIComponent(search)}`,{
            method: "GET",
            headers: {"Content-Type":"application/json"},
        })
        const data = await res.json();
        DataReturn(data)
    }catch (err){
        throw {status:res.status}
    }
}