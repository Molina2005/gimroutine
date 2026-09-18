import { StateUserOptions } from "../../cammon/selectOption.js";
import { logicAllUserStates } from "../../user_states/logicAllUserStates.js";
import { LogicClient } from "./LogicClientCreate.js";
document.addEventListener("DOMContentLoaded", async function(){
        const form = document.getElementById("addClientForm")
        // Se obtienen los estados existentes [activo : inactivo]
        const dataClientState = await logicAllUserStates()
        // Se obtiene id del selector para guardar la informacion de roles
        const ClientStateSelect = document.getElementById("state")
        dataClientState.forEach(state =>{
            StateUserOptions(state, ClientStateSelect)
        })
    form.addEventListener("submit", async function(e){
        e.preventDefault();
        const formData = {
            name: document.getElementById("name").value,
            phone: document.getElementById("phone").value,
            id_users_state: Number(document.getElementById("state").value),
            gmail: document.getElementById("gmail").value,
            nit: document.getElementById("nit").value,
        }
        try{
            const res = await LogicClient(formData);
            alert("Cliente creado correctamente", res)
            form.reset();
        }catch (err){
            if (err.status == "409"){
                alert("Cliente ya existe en el sistema")
            }else{
                alert("Error en el servidor")   
            }
        }
    })
})