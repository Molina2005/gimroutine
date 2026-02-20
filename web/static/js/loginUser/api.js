export function login(data){
    return fetch("/login",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify(data)        
    })
    .then(res => res.json())   
    .then(json =>{
        console.log(json)
        return json;
    })
    .catch(err =>{
        console.log("Error en login", err)
    });
}  