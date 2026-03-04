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
            alert("Bienvenido a App", result)
            form.reset();
            window.location.replace("/home")
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