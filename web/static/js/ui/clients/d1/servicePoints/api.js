import { API_BACKEND_RAILWAY } from "../../../../config/config.js";
import { ReturnSearchServicePoints } from "./search.js";

// Conexiones entre backend y frontend para manejar buscador 
export async function ConnectionServicePointsSearch(search, idRegional){
    try{
        const res = await fetch(`${API_BACKEND_RAILWAY}/servicePoints/${encodeURIComponent(search)}/${idRegional}`,{
            method: "GET",
        });
        const data = await res.json();
        ReturnSearchServicePoints(data);
    }catch (err){
        throw {status:res.status};
    }
}   
// Conexiones entre backend y frontend para manejar eliminador 
export async function ConnectionServicePointDelete(id){
    const res = await fetch(`${API_BACKEND_RAILWAY}/servicePoint/${id}`, {
        method:"DELETE"
    })
    if (!res.ok){
        return {status:res.status}
    }
    return await res.json()
}

