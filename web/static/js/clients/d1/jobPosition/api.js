import { API_BACKEND_RAILWAY } from "../../../config/config.js"
// Funcion para conectar backend y frontend 
export async function ConnectionAllJobPosition(){
    const res = await fetch(`${API_BACKEND_RAILWAY}AllJobPosition`, {
        method:"GET",
    })
    if (!res.ok){
        return {status: res.status}
    }
    return res.json()
}