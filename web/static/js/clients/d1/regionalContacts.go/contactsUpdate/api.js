// Funcion para poder obtener informacion del contacto regional segun id 
export async function ConnectionInfoRegionalContacts(id){
    const res = await fetch(`/regionalContact/${id}`, {
        method:"GET",
        headers:{"Content-Type":"application/json"} 
    });
    if (!res.ok){
        return {status:res.status}
    }
    return await res.json();
}   

export async function ConnectionUpdateRegionalContacts(id, data){
    const res = await fetch(`/regionalContacts/${id}`,{
        method:"PUT",
        headers: {"Content-Type":"application/json"},
        body:JSON.stringify(data)
    })
    if (!res.ok){
        return {status:res.status}
    }
    return res.json();
}


