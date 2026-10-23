// Funcion para hacer cconexion entre backend y frontend
export async function ConnectionCreateServicePoints(data){
    const res = await fetch("/servicePoints",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify(data)
    })
    if (!res.ok){
        throw {status: res.status};
    }
    return res.json();
} 