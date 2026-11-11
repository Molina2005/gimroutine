import { API_BACKEND_RAILWAY } from "../config/config.js";
export async function ConnectionAllUserStates(){
    const res = await fetch(`${API_BACKEND_RAILWAY}/AllUserStates`,{
        method: "GET",
        headers:{"Content-Type":"application/json"},
    })
    if (!res.ok){
        return {status:res.status};
    }
    return res.json();
}