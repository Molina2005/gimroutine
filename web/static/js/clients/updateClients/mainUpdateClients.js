import { StateUserOptions } from "../../cammon/selectOption.js";
import { logicAllUserStates } from "../../user_states/logicAllUserStates.js";
import { logicInfoClient } from "./logicUpdateClients.js";
import { logicUpdateClients } from "./logicUpdateClients.js";
document.addEventListener("DOMContentLoaded", async function(){
    // Obtiene la URL en que se encuentra actualmennte
    const url = new URL(window.location.href);
    // Busca el valor id dentro de la url
    const id = url.searchParams.get("id");
    // Se para el id a la funcion de logic, de igual manera aqui recibe la informacion cuando llega de logic
    const data = await logicInfoClient(id);
    // Se obtienen los estados existentes [activo : inactivo]
    const dataClientState = await logicAllUserStates()
    // Se obtiene id del selector para guardar la informacion de roles
    const ClientStateSelect = document.getElementById("state")
    dataClientState.forEach(state =>{
        StateUserOptions(state, ClientStateSelect)
    })
    const formUpdate = document.getElementById("form-update-clients")
    document.getElementById("name").value = data.client.Name
    document.getElementById("phone").value = data.client.Phone
    document.getElementById("state").value = data.client.IdState
    document.getElementById("nit").value = data.client.Nit
    document.getElementById("gmail").value = data.client.Gmail    
    // Evento para cuando el usuario de click a boton actualizar se actualizen los datos con lo que ingreso en los inputs 
    formUpdate.addEventListener("submit", async function(e){
        e.preventDefault();
        // Obtiene la informacion de que lo que envian los inputs
        const dataClient = {
            name : document.getElementById("name").value,
            phone : document.getElementById("phone").value,
            id_users_state : Number(document.getElementById("state").value),
            nit : document.getElementById("nit").value,
            gmail : document.getElementById("gmail").value
        };
        // Retorna la informacion a la siguiente funcion para que cumpla su proceso
        const dataUpdate = await logicUpdateClients(id,dataClient);
        alert("cliente actualizado correctamente")
        window.location.replace("/clientsInfo")
        return dataUpdate
    })
})

