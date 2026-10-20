import { ReturnSearchRegionalContact } from "./search.js";

// Funcion obtener informacion para eliminar contacto de sistema 
export async function ConnectionRegionalContactDelete(id){
    const res = await fetch(`/regionalContacts/${id}`, {
        method:"DELETE",
    });
    if (!res.ok){
        return {status: res.status}
    }
    return res.json()
}

// Conexion entre backend y frontend 
export async function ConnectionRegionalContactSearch(search, idRegional){
    try{
        const res = await fetch(`/regionalContacts/${encodeURIComponent(search)}/${idRegional}`,{
            method: "GET",
            headers: {"Content-Type":"application/json"},
        })
        const data = await res.json();
        ReturnSearchRegionalContact(data)
    }catch (err){
        throw {status:res.status}
    }
} 