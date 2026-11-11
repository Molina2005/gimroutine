import { API_BACKEND_RAILWAY } from "../../../../config/config.js";
// Funcion para hacer cconexion entre backend y frontend
export async function ConnectionCreateRegional(data){
    const res = await fetch(`${API_BACKEND_RAILWAY}/regional`,{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify(data)
    })
    if (!res.ok){
        throw {status: res.status};
    }
    return res.json();
}