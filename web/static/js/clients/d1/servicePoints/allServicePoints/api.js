import { API_BACKEND_RAILWAY } from "../../../../config/config.js";
// Conexion con el backend para poder obtener informacion de los puntos de servicio
export async function ConnectionAllServicePoints(idRegional){
    const res = await fetch(`${API_BACKEND_RAILWAY}/servicePoints/${idRegional}`, {
        method:"GET",
    });
    if (!res.ok){
        return {status:res.status};
    }
    return await res.json();
}