export async function ConnectionPlans(data){
    const res = await fetch("/addPlans",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body: JSON.stringify(data)
    });

    if (!res.ok){
        throw {status:res.status}
    }
    const x = await res.json();
    return x
}