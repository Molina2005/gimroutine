// Funcion para recibir informacion de planes que vienen del backend
export async function ConnectionPlansIndex(){
    const data = await fetch("PlansIndex", {
        method:"GET",
        headers:{"Content-Type":"application/json"},
    })
    if (!data.ok){
        return {status:data.status}
    }
    return await data.json()
}
