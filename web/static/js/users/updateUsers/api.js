// fetch para traer info usuario por id
export async function InfoUser(id){
    const res = await fetch(`/users/${id}`,{
        method: "GET",
        headers: {"Content-Type":"application/json"},
    })
    if (!res.ok){
        throw {status: res.status}
    }
    return await res.json();
} 

// fetch para actualizar info usuario
export async function updateUser(id,data){
    const res = await fetch(`/users/${id}`,{
        method: "PUT",
        headers: {"Content-Type":"application/json"},
        body:JSON.stringify(data)
    })
    if (!res.ok){
        throw {status: res.status}
    }
    return await res.json();
}