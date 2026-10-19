import { StateJobPosition, StateUserOptions } from "../../../../cammon/selectOption.js";
import { logicAllUserStates } from "../../../../user_states/logicAllUserStates.js";
import { LogicAllJobPosition } from "../../jobPosition/logicJobPosition.js";
import { LogicInfoRegionalContacts, LogicUpdateRegionalContacts } from "./logicUpdate.js";

document.addEventListener("DOMContentLoaded", async function(){
    // Se obtiene url para luego poder obtener id del contacto regional
    const url = new URL(window.location.href);
    const id = url.searchParams.get("id");
    // Se obtiene id de la reginal guardado en sessionStorage
    const id_regional = Number(sessionStorage.getItem("idRegional"));
    // Se obtiene cada una de la informacion correspondiente para poder procesarla
    const data = await LogicInfoRegionalContacts(id);
    const dataJobPosition = await LogicAllJobPosition();
    const dataState = await logicAllUserStates();
    // Se obtienen selectores del html y luego se procesan para poder pasarlos al selector y verlos en pantalla
    const jobpositionSelect = document.getElementById("jobposition");
    const stateSelect = document.getElementById("state");
    dataJobPosition.forEach(jobposition => {
        StateJobPosition(jobposition, jobpositionSelect)
    })
    dataState.forEach(state => {
        StateUserOptions(state, stateSelect)
    })
    // Mostrar en cada input el valor actual que tiene cada contacto
    document.getElementById("name").value = data.regionaContact.Name
    document.getElementById("lastname").value = data.regionaContact.LastName
    document.getElementById("gmail").value = data.regionaContact.Gmail
    document.getElementById("phone").value = data.regionaContact.Phone
    document.getElementById("jobposition").value = data.regionaContact.IdJogPosition
    document.getElementById("state").value = data.regionaContact.IdUserState
    // Se obtiene formulario y se obtiene info de input para actualizar y pasar a actualizarla 
    const formUpdate = document.getElementById("form-update-contacts")
    formUpdate.addEventListener("submit", async function(e){
        e.preventDefault();
        const dataRegionalContact = {
            name:document.getElementById("name").value,
            lastname:document.getElementById("lastname").value,
            gmail:document.getElementById("gmail").value,
            phone: document.getElementById("phone").value,
            id_regional: id_regional,
            id_job_position:  Number(document.getElementById("jobposition").value),
            id_state: Number(document.getElementById("state").value),
        }
        const dataUpdate = await LogicUpdateRegionalContacts(id, dataRegionalContact)
        alert("Contacto actualizado correctamente");
        window.location.replace("/regionalContacts");
        return dataUpdate
    })  
})
