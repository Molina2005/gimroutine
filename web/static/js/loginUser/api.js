export async function login(data){
    try {
        const res = await fetch("/login",{
            method:"POST",
            headers:{"Content-Type":"application/json"},
            body:JSON.stringify(data)        
        });
        const json = await res.json();
        if (!res.ok){
            return {error: true, message: json.error || "Error el login"};
        }
        return json;
    } catch (err){
        return {error: true, message: json.error || "Error inesperado en login"};
    }
}  