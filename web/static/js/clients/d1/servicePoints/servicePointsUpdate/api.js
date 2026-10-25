export async function ConnectionServicePointInformation(id){
    const res = await fetch(`/servicePoint/${id}`, {
        method:"GET",
        headers:{"Content-Type":"application/json"},
    });
    if (!res.ok){
        return {status:res.status};
    }
    return await res.json();
}