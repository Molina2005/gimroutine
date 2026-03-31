export async function ConnectionClient(data) {
    const res = await fetch("/Clients",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify(data)
    });
    if (!res.ok){
        throw {status: res.status};
    }
    return res.json();
}