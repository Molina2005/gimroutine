// Funcion para conectar backend y frontend 
export async function ConnectionAllRegional(){
    const res = await fetch("allRegional", {
        method:"GET",
        headers:{"Content-Type":"application/json"}
    })
    if (!res.ok){
        return {status: res.status}
    }
    return res.json()
}