import { logicLogin } from "./Login.js";
document.addEventListener("DOMContentLoaded", function(){
    console.log("entro a login de inicio de sesion")
    const form =  document.getElementById("loginUserForm");
    form.addEventListener("submit", function(e){
        e.preventDefault();
        const formData = {
            name: document.getElementById("nameOrEmail").value,
            password: document.getElementById("password").value
        }
        logicLogin(formData);
        
    })
    .then(data =>{
        form.reset();
        window.location.replace("/Home")
    })
    .catch(err =>{
        alert("Credenciales incorrectas, Intente de nuevo")
    })
})