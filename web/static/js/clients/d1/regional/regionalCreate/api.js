// Funcion para hacer cconexion entre backend y frontend
export async function ConnectionCreateRegional(data){
    const res = await fetch("/regional",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify(data)
    })
    if (!res.ok){
        throw {status: res.status};
    }
    console.log("fetch", res)
    return res.json();
}