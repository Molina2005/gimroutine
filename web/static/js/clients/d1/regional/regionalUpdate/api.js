
export async function ConnectionInfoRegional(id){
        const res = await fetch(`/regional/${id}`, {
        method:"GET",
        headers:{"Content-Type":"application/json"}
    })
    if (!res.ok){
        return {status: res.status} 
    }
    return await res.json()
}

// fetch para actualizar info usuario
export async function ConnectionUpdateRegional(id,data){
    const res = await fetch(`/regional/${id}`,{
        method: "PUT",
        headers: {"Content-Type":"application/json"},
        body:JSON.stringify(data)
    })
    if (!res.ok){
        throw {status: res.status}
    }
    return await res.json();
}