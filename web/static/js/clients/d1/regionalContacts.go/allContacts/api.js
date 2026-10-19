// Funcion para conectar backend y frontend 
export async function ConnectionAllRegionalContacts(id_regional){
    const res = await fetch(`AllregionalContacts/${id_regional}`, {
        method:"GET",
        headers:{"Content-Type":"application/json"}
    })
    if (!res.ok){
        return {status: res.status}
    }
    return res.json()
}