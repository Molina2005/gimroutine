export async function ConnectionRegionalContactsCreate(data){
    const res = await fetch("regionalContacts", {
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body: JSON.stringify(data)
    })
    if (!res.ok){
        throw {status:res.status}
    }
    return res.json()
}