import { API_BACKEND_RAILWAY } from "../../config/config.js";

export async function ConnectionClient(data) {
    const res = await fetch(`${API_BACKEND_RAILWAY}/AddClients`,{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify(data)
    });
    if (!res.ok){
        throw {status: res.status};
    }
    return res.json();
}