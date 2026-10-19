import { StateClientOptions, StateUserOptions } from "../../../../cammon/selectOption.js";
import { logicAllUserStates } from "../../../../user_states/logicAllUserStates.js";
import { LogicAllClient } from "../../../allClients/logicAllClient.js";
import { LogicInfoRegional, LogicUpdateRegional } from "./logicUpdate.js";

document.addEventListener("DOMContentLoaded", async function(){
    // Obtiene la url actual
    const ulr = new URL(window.location.href);
    // Se obtiene dato id dentro de esa url
    const id = ulr.searchParams.get("id")
    // Se pasa el id a la funcion de logic, de igual manera aqui recibe la informacion cuando llega de logic
        const data = await LogicInfoRegional(id);
        // contiene toda la informacion de los clientes creados
        const dataClients = await LogicAllClient()
        // Contiene toda la informacion de los estados de usuarios creados
        const dataUserStates = await logicAllUserStates()
        // Se obtiene id del selector para guardar la informacion 
        const clientSelect = document.getElementById("client")
        const UserStatesSelect = document.getElementById("state")
        // Se recorre cada uno de los clientes y estados y los muestra en la web
        dataClients.forEach(client =>{
            StateClientOptions(client, clientSelect)
        });
        dataUserStates.forEach(state =>{
            StateUserOptions(state, UserStatesSelect )
        })
        const formUpdate = document.getElementById("form-update-regional");
        document.getElementById("name").value = data.regional.Name
        document.getElementById("client").value = data.regional.IdClient
        document.getElementById("state").value = data.regional.IdUserState
        // Evento para cuando el usuario de click a boton actualizar se actualizen los datos con lo que ingreso en los inputs 
        formUpdate.addEventListener("submit", async function(e){
            e.preventDefault();
            // Obtiene la informacion de que lo que envian los inputs
            const dataRegion = {
                name : document.getElementById("name").value,
                id_client : Number(document.getElementById("client").value),
                id_state : Number(document.getElementById("state").value),
            };
            // Retorna la informacion a la siguiente funcion para que cumpla su proceso
            const dataUpdate = await LogicUpdateRegional(id,dataRegion);
            alert("regional actualizada correctamente");
            window.location.replace("/listRegional");
            return dataUpdate
        })

})