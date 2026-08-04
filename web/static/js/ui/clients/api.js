import { ReturnSearchClient } from "./search.js";
// Funcion para recibir informacion del backend eliminacion de clientes
export async function ConnectionClientDelete(id){
    const data = await fetch(`/client/${id}`, {
        method: "DELETE"
    })
    if (!data.ok){
        throw {status:data.status}
    }
    return data.json();
}

// Funcion para recibir informacion del backend del buscador de clientes
export async function ConnectionSearchClient(search){
    try{
        const res = await fetch(`/ClientSearch/${encodeURIComponent(search)}`,{
            method: "GET",
            headers: {"Content-Type":"application/json"},
        })
        const data = await res.json();
        ReturnSearchClient(data)
    }catch (err){
        throw {status:res.status}
    }
}