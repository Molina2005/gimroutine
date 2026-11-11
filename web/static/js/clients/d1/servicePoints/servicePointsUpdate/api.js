import { API_BACKEND_RAILWAY } from "../../../../config/config.js";
export async function ConnectionServicePointInformation(id){
    const res = await fetch(`${API_BACKEND_RAILWAY}/servicePoint/${id}`, {
        method:"GET",
    });
    if (!res.ok){
        return {status:res.status};
    }
    return await res.json();
}

export async function ConnectionServicePointUpdate(id, data){
    const res = await fetch(`${API_BACKEND_RAILWAY}/servicePoint/${id}`, {
        method:"PUT",
        headers:{"Content-Type":"application/json"},
        body: JSON.stringify(data)
    });
    if (!res.ok) {
        return {status:res.status};
    }
    return res.json();
}