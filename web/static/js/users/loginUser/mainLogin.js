import { logicLogin } from "./Login.js";
document.addEventListener("DOMContentLoaded", async function(){
    const form =  document.getElementById("loginUserForm");
    form.addEventListener("submit", async function(e){
        e.preventDefault();
        const formData = { 
            email: document.getElementById("email").value,
            password: document.getElementById("password").value
        }; 
        // Try maneja el error sin romper el programa 
        // catch cumple con enviar el error en caso de que falle
        try{
            const result = await logicLogin(formData);
            // Se guarda token de usuario en el navegador y queda listo para usarlo en cualquier parte 
            localStorage.setItem("token",result.token);            
            // Objeto el cual guarda las rutas de cada uno de los roles
            const routes = {
                administrator : "/adminHome",
                supervisor : "/supervisorHome",
                technical : "/TechnicalHome",
            }
            // Redireccion a cada pagina de inicio segun el rol
            window.location.replace(routes[result.role])
            form.reset
        }catch (err){
            // Si error es Estado no autorizado(401) entra a alert si no es un error del servidor
            if (err.status == "401") {
                alert("Credenciales incorrectas intente de nuevo")
            }else{
                alert("Error en el servidor")
            }
        } 
    })
})