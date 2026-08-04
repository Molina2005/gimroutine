// Funcion para poder verificar el token JWT y que sea el mismo que se obtubo al iniciar sesion
export function VerifyToken(){
    // Se obtiene token que se creo en el momento en que se inicio sesion
    const token = localStorage.getItem("token")
    // Si es diferente lo redirije a apartado de login
    if (!token){
        window.location.replace("/index")
        return
    }
}