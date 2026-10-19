export async function ConnectionAllUserStates(){
    const res = await fetch("/AllUserStates",{
        method: "GET",
        headers:{"Content-Type":"application/json"},
    })
    if (!res.ok){
        return {status:res.status};
    }
    return res.json();
}