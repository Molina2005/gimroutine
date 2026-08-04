// Funcion para eliminar token JWT de cuando inicio sesion el usuario 
export function Logout(){
    localStorage.removeItem("token")
    window.location.replace("/index")
}