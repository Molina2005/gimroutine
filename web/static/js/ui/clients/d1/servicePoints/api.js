import { ReturnSearchServicePoints } from "./search.js";

// Conexiones entre backend y frontend para manejar buscador 
export async function ConnectionServicePointsSearch(search, idRegional){
    try{
        const res = await fetch(`/servicePoints/${encodeURIComponent(search)}/${idRegional}`,{
            method: "GET",
            headers: {"Content-Type":"application/json"},
        });
        const data = await res.json();
        ReturnSearchServicePoints(data);
    }catch (err){
        throw {status:res.status};
    }
}   
// Conexiones entre backend y frontend para manejar eliminador 
export async function ConnectionServicePointDelete(id){
    const res = await fetch(`/servicePoint/${id}`, {
        method:"DELETE"
    })
    if (!res.ok){
        return {status:res.status}
    }
    return await res.json()
}

