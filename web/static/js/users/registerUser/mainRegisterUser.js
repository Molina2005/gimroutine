import { logicRegisterUser } from "./RegisterUser.js";
import { LogicAllRoles } from "../../roles/allRoles/logicRoles.js";
import { logicAllUserStates } from "../../user_states/logicAllUserStates.js";
import { RolOptions } from "../../cammon/selectOption.js";
import { StateUserOptions } from "../../cammon/selectOption.js";

// DOM: Permite interactuar con el html, da acceso a trabajar con el formulario y botones
document.addEventListener("DOMContentLoaded", async function(){
    // contiene toda la informacion de los roles creados
    const dataRoles = await LogicAllRoles()
    const dataUserState = await logicAllUserStates()
    // Se obtiene id del selector para guardar la informacion de roles
    const rolSelect = document.getElementById("role")
    const UserStateSelect = document.getElementById("state")
    // Se recorre cada uno de los roles creados y los muestra en la web
    dataRoles.forEach(rol =>{
        RolOptions(rol, rolSelect)
    })
    dataUserState.forEach(state =>{
        StateUserOptions(state, UserStateSelect)
    })
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
            lastname: document.getElementById("lastname").value,
            id_rol: Number(document.getElementById("role").value),
            phone: Number(document.getElementById("phone").value),
            id_users_state: Number(document.getElementById("state").value),
        }       
        try{
            // Se le pasa a la función login la información del formData
            // para que ejecute la lógica de inicio de sesión (comunicación con el backend)
            const data = await logicRegisterUser(formData);
            alert("Usuario creado correctamente",data);
            form.reset();
        }catch (err){
            if (err.status == "409"){
                alert("Usuario ya existe en el sistema")
            }else{
                alert("Error en el servidor")
            }
        }
    });
});

