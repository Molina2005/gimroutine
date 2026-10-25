import { StateUserOptions } from "../../../../cammon/selectOption.js";
import { logicAllUserStates } from "../../../../user_states/logicAllUserStates.js";
import { LogicServicePointsUpdate } from "./logicUpdate.js";

document.addEventListener("DOMContentLoaded", async function(){
    const url = new URL(window.location.href);
    const id = url.searchParams.get("id");
    const data = await LogicServicePointsUpdate(id);
    const dataState = await logicAllUserStates();
    const stateSelect = document.getElementById("state");

    dataState.forEach(state =>{
        StateUserOptions(state, stateSelect);
    });

    document.getElementById("name").value = data.servicePoint.Name;
    document.getElementById("addres").value = data.servicePoint.Addres;
    document.getElementById("city").value = data.servicePoint.City;
    document.getElementById("state").value = data.servicePoint.IdState;

    const FormUpdate = document.getElementById("form-update-service-points")
    FormUpdate.addEventListener("submit", function(e){
        e.preventDefault();

        const dataServicePoint = {
            name :document.getElementById("name").value,
            addres: document.getElementById("addres").value, 
            city: document.getElementById("city").value ,
            id_state:document.getElementById("state").value
        }

    })
        
})