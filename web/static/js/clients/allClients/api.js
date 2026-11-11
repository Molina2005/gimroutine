import { API_BACKEND_RAILWAY } from "../../config/config.js";
export async function ConnectionAllClient(){
    const res = await fetch(`${API_BACKEND_RAILWAY}/AllClients`,{
        method:"GET",
        headers: {"Content-Type":"application/json",},
    });
    if (!res.ok){
        throw {status:res.status};
    }
    return res.json();
}