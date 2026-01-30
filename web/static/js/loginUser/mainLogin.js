import { logicLogin } from "./Login.js";
document.addEventListener("DOMContentLoaded", function(){
    const form =  document.getElementById("loginUserForm");
    form.addEventListener("submit", async function(e){
        e.preventDefault();
        const formData = {
            name: document.getElementById("nameOrEmail").value,
            password: document.getElementById("password").value
        };
        // Try maneja el error sin rompero el programa 
        // catch cumple con enviar el error en caso de que falle
        try{
            // Await se usa para esperar una promesa, hasta que esto se cumpla pasa a lo siguiente
            await logicLogin(formData);
            form.reset();
            // url para que redirija a la pagina de Home
            window.location.replace("/Home");
        }catch (err){
            alert("Credenciales incorrectas intente de nuevo")
        } 
    })
})