// Conexion con el backend para poder obtener informacion de los puntos de servicio
export async function ConnectionAllServicePoints(idRegional){
    const res = await fetch(`/servicePoints/${idRegional}`, {
        method:"GET",
        headers:{"Content-Type":"application/json"},
    });
    if (!res.ok){
        return {status:res.status};
    }
    return await res.json();
}