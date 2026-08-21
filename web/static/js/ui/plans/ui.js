import { logicAllPlans } from "../../plans/allPlans/logicAllPlans.js";
import { RenderPlansRow } from "../../cammon/renderDataPlans.js";
// Mostrar informacion de todos los planes con sus respectivos botones de cruds
    document.addEventListener("DOMContentLoaded",async function(){
        const plans = await logicAllPlans();
        const tableBody = document.getElementById("plan-table");
        tableBody.innerHTML = ""
        plans.forEach(plan =>{
            tableBody.appendChild(RenderPlansRow(plan))
        })
    })