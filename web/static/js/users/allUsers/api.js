export async function ConnectionAllUsers(){
    const res = await fetch("allUsers",{
        method: "GET",
        headers:{"Content-Type":"application/json"},
    })
    if (!res.ok){
        return {status:res.status};
    }
    return res.json();
}