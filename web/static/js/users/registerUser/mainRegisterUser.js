import { logicRegisterUser } from "./RegisterUser.js";
// DOM: Permite interactuar con el html, da acceso a trabajar con el formulario y botones
document.addEventListener("DOMContentLoaded", function(){
    // Varible que obtiene informacion del formulario 
    const form = document.getElementById("userForm")
    // Evento que al momento de dar oprimir el boton hace todo lo que se contiene dentro de la funcion
    form.addEventListener("submit", async function(e){
        // evita el comportamiento por defecto del formulario (recargar la página) y perder la informacion.
        e.preventDefault();
        // FormData va a guardar toda la informacion que se obtega dentro de esa const
        const formData = {
            // se obtiene los elementos y se usa value para obtener el valor actual del input en ese momento
            name: document.getElementById("name").value,
            email: document.getElementById("email").value,
            password: document.getElementById("password").value,
            role: document.getElementById("role").value,
        }
        console.log(formData)
        
        try{
            // Se le pasa a la función login la información del formData
            // para que ejecute la lógica de inicio de sesión (comunicación con el backend)
            const data = await logicRegisterUser(formData);
            alert("Usuario creado correctamente",data);
            form.reset();
            alert("Usuario creado correctamente")
        }catch (err){
            if (err.status == "409"){
                alert("Usuario ya existe en el sistema")
            }else{
                alert("Error en el servidor")
            }
        }
    });
});

