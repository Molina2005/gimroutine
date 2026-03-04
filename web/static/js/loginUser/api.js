export async function login(data){
    const res = await fetch("/login",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify(data)        
    });
    if (!res.ok){
        // se usa status ya que es la forma de manejar el verdadero error que envia el backend por es estado
        throw {status : res.status}
    }
    return res.json();
}  