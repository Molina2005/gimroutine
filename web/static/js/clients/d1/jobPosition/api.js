// Funcion para conectar backend y frontend 
export async function ConnectionAllJobPosition(){
    const res = await fetch("AllJobPosition", {
        method:"GET",
        headers:{"Content-Type":"application/json"}
    })
    if (!res.ok){
        return {status: res.status}
    }
    return res.json()
}