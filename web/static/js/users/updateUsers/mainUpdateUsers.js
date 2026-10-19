import { logicInfoUser, logicUpdateUser } from "./logicUpdateUsers.js"
import { LogicAllRoles } from "../../roles/allRoles/logicRoles.js";
import { logicAllUserStates } from "../../user_states/logicAllUserStates.js";
import { RolOptions, StateUserOptions } from "../../cammon/selectOption.js";
document.addEventListener("DOMContentLoaded", async function(){
    // Obtiene la URL en que se encuentra actualmennte
    const url = new URL(window.location.href);
    // Busca el valor id dentro de la url
    const id = url.searchParams.get("id");
    // Se pasa el id a la funcion de logic, de igual manera aqui recibe la informacion cuando llega de logic
    const data = await logicInfoUser(id);
    // contiene toda la informacion de los roles creados
    const dataRoles = await LogicAllRoles()
    // Contiene toda la informacion de los estados de usuarios creados
    const dataUserStates = await logicAllUserStates()
    // Se obtiene id del selector para guardar la informacion 
    const rolSelect = document.getElementById("role")
    const UserStatesSelect = document.getElementById("state")
    // Se recorre cada uno de los roles y estados creados y los muestra en la web
    dataRoles.forEach(rol =>{
        RolOptions(rol, rolSelect)
    });
    dataUserStates.forEach(state =>{
        StateUserOptions(state, UserStatesSelect )
    })
    const formUpdate = document.getElementById("form-update-users");
    document.getElementById("name").value = data.user.Name
    document.getElementById("lastname").value = data.user.LastName
    document.getElementById("email").value = data.user.Email
    document.getElementById("role").value = data.user.IdRole
    document.getElementById("phone").value = data.user.Phone
    document.getElementById("state").value = data.user.Id_users_state
    // Evento para cuando el usuario de click a boton actualizar se actualizen los datos con lo que ingreso en los inputs 
    formUpdate.addEventListener("submit", async function(e){
        e.preventDefault();
        // Obtiene la informacion de que lo que envian los inputs
        const dataUser = {
            name : document.getElementById("name").value,
            email : document.getElementById("email").value,
            password : document.getElementById("password").value,
            idrol : Number(document.getElementById("role").value),
            lastname: document.getElementById("lastname").value,
            phone: document.getElementById("phone").value,
            id_state: Number(document.getElementById("state").value),
        };
        // Retorna la informacion a la siguiente funcion para que cumnpla su proceso
        const dataUpdate = await logicUpdateUser(id,dataUser);
        alert("usuario actualizado correctamente");
        window.location.replace("/usersInfo");
        return dataUpdate
    })
})

