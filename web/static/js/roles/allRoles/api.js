// Funcion para obtener roles creados 
export async function connectionAllRoles(data){
    const res = await fetch("Allroles",{
        method:"GET",
        headers:{"Content-Type":"application/json"},
    })
    if (!res.ok){
        throw {status: res.status};
    }
    return res.json();
}