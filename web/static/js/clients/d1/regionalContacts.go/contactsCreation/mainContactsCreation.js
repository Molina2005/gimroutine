import { StateJobPosition, StateUserOptions } from "../../../../cammon/selectOption.js";
import { logicAllUserStates } from "../../../../user_states/logicAllUserStates.js";
import { LogicAllJobPosition } from "../../jobPosition/logicJobPosition.js";
import { LogicCreateRegionalContacts } from "./logicContactsCreation.js";

document.addEventListener("DOMContentLoaded", async function(){
    const dataJobPosition = await LogicAllJobPosition()
    const dataUserStates = await logicAllUserStates()
    const JobPositionSelect = document.getElementById("jobposition") 
    const StateSelect = document.getElementById("state") 
    // Se obtiene cargo del contacto regional para realizar selector 
    dataJobPosition.forEach(position => {
        StateJobPosition(position, JobPositionSelect)
    })
    // Se obtiene id almacenado segun la reginal seleccionada 
    const IdRegional = Number(sessionStorage.getItem("idRegional"));
    // Se obtiene estado para realizar selector 
    dataUserStates.forEach(state =>{
        StateUserOptions(state,StateSelect )
    })
    const formRegioanalCreate = document.getElementById("form-regional-create")
    formRegioanalCreate.addEventListener("submit", async function(e){
        e.preventDefault();        
        const FormData = {
            name : document.getElementById("name").value,
            lastname : document.getElementById("lastname").value,
            gmail : document.getElementById("gmail").value,
            phone : document.getElementById("phone").value,
            id_job_position : Number(document.getElementById("jobposition").value),
            id_regional : IdRegional,
            id_user_state: Number(document.getElementById("state").value),  
        }
        try{
           await LogicCreateRegionalContacts(FormData);
           alert("Contacto creado correctamente");
           formRegioanalCreate.reset();
        }catch(err) {
            if (err.status == "409"){
                alert("contacto ya existe en el sistema")
            }else{
                alert("Error en el servidor")
            }
        }
    })
})