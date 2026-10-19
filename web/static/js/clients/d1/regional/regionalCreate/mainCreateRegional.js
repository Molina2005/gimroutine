import { StateClientOptions, StateUserOptions } from "../../../../cammon/selectOption.js";
import { logicAllUserStates } from "../../../../user_states/logicAllUserStates.js";
import { LogicAllClient } from "../../../allClients/logicAllClient.js";
import { LogicCreateRegional } from "./logicRegional.js";

// Funcion para obtener informacion de formulario y crear regional
document.addEventListener("DOMContentLoaded", async function(){
    const dataClient = await LogicAllClient()
    const dataRegionalState = await logicAllUserStates()
    const regionalClientSelect = document.getElementById("client")
    const regionalStateSelect = document.getElementById("state")
    // Se recorre cada uno de los clientes y estados existentes y los muestra en la web
    dataClient.forEach(client =>{
        StateClientOptions(client,regionalClientSelect )
    })
    dataRegionalState.forEach(state =>{
        StateUserOptions(state, regionalStateSelect)
    })
    // Varible que obtiene informacion del formulario 
    const form = document.getElementById("CreationRegionalform")
    // Evento que al momento de dar oprimir el boton hace todo lo que se contiene dentro de la funcion
    form.addEventListener("submit", async function(e){
        // evita el comportamiento por defecto del formulario (recargar la página) y perder la informacion.
        e.preventDefault();
        // FormData va a guardar toda la informacion que se obtega dentro de esa const
        const formData = {
            // se obtiene los elementos y se usa value para obtener el valor actual del input en ese momento
            name: document.getElementById("name").value,
            id_client: Number(document.getElementById("client").value),
            id_user_state: Number(document.getElementById("state").value),
        }     
        try{
            // Se le pasa a la función LogicCreateRegional la información del formData
            const data = await LogicCreateRegional(formData);
            alert("Regional creada correctamente",data);
            form.reset();
        }catch (err){
            if (err.status == "409"){
                alert("Regional ya existe en el sistema")
            }else{
                alert("Error en el servidor")
            }
        }
    });
});

