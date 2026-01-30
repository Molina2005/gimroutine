export function login(data){
    return fetch("/login",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify(data)        
    })
    .then(response =>{
        if (!response.ok){
            throw new Error("Error ingresar al sistema");
        }
        return response.text()
    })
}