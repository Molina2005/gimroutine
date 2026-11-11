import { API_BACKEND_RAILWAY } from "../../../../config/config.js"
// Funcion para conectar backend y frontend 
export async function ConnectionAllRegionalContacts(id_regional){
    const res = await fetch(`${API_BACKEND_RAILWAY}AllregionalContacts/${id_regional}`, {
        method:"GET",
    })
    if (!res.ok){
        return {status: res.status}
    }
    return res.json()
}