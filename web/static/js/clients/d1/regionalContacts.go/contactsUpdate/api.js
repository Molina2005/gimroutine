import { API_BACKEND_RAILWAY } from "../../../../config/config.js";
// Funcion para poder obtener informacion del contacto regional segun id 
export async function ConnectionInfoRegionalContacts(id){
    const res = await fetch(`${API_BACKEND_RAILWAY}/regionalContact/${id}`, {
        method:"GET",
    });
    if (!res.ok){
        return {status:res.status}
    }
    return await res.json();
}   

export async function ConnectionUpdateRegionalContacts(id, data){
    const res = await fetch(`${API_BACKEND_RAILWAY}/regionalContacts/${id}`,{
        method:"PUT",
        headers: {"Content-Type":"application/json"},
        body:JSON.stringify(data)
    })
    if (!res.ok){
        return {status:res.status}
    }
    return res.json();
}


