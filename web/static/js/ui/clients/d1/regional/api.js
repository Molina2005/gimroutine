import { ReturnSearchRegional } from "./search.js";

// Conexion entre backend y frontend para eliminacion de regionales
export async function ConnectionRegionalDelete(id){
    const res = await fetch(`/regional/${id}`, {
        method:"DELETE",
    })
    if (!res.ok){
        return {status:res.status}
    }
    return res.json();
}

// Conexion entre backend y frontend 
export async function ConnectionRegionalSearch(search){
    try{
        const res = await fetch(`/regionalSearch/${encodeURIComponent(search)}`,{
            method: "GET",
            headers: {"Content-Type":"application/json"},
        })
        const data = await res.json();
        ReturnSearchRegional(data)
    }catch (err){
        throw {status:res.status}
        
    }
} 

