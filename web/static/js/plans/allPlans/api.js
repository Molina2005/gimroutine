// Conexion con backend para traer todos los planes creados
export async function ConnectionAllPlans(){
    const res = await fetch("AllPlans",{
        method: "GET",
        headers:{"Content-Type":"application/json"},
    })
    if (!res.ok){
        return {status:res.status};
    }
    return res.json();
}