export async function ConnectionAllClient(data){
    const res = await fetch("/AllClients",{
        method:"GET",
        headers: {"Content-Type":"application/json",},
    });
    if (!res.ok){
        throw {status:res.status};
    }
    return res.json();
}