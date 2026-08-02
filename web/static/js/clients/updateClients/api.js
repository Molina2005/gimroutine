// fetch para traer info client por id
export async function InfoClient(id){
    const res = await fetch(`/client/${id}`,{
        method: "GET",
        headers: {"Content-Type":"application/json"},
    })
    if (!res.ok){
        throw {status: res.status}
    }
    return await res.json();
    
} 

// Fetch para actulizar informacion cliente
export async function UpdateClient(id, data){
    const res = await fetch(`/client/${id}`,{
        method:"PUT",
        headers:{"Content-Type":"application/json"},
        body: JSON.stringify(data)
    })
    if(!res){
        throw {status:res.status}
    }
    return await res.json();
}