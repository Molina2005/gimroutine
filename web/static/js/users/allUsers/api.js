import { API_BACKEND_RAILWAY } from "../../config/config.js";
export async function ConnectionAllUsers(){
    const res = await fetch(`${API_BACKEND_RAILWAY}/allUsers`,{
        method: "GET",
    })
    if (!res.ok){
        return {status:res.status};
    }
    return res.json();
}