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
            if (result.error){
                alert(result.message);
            }else{
                form.reset();
                window.location.replace("home")
            }

        }catch (err){
            alert( err.message || "Credenciales incorrectas intente de nuevo")
        } 
    })
})