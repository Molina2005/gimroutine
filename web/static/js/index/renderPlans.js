import { PlanIndex } from "../cammon/plansIndex.js";
import { ConnectionPlansIndex } from "./index.js";
// Funcion para poder agregar planes al index
document.addEventListener("DOMContentLoaded", async function(){
    // Funcion que trae todos los planes que se recopilaron del backend
    const dataPlans = await ConnectionPlansIndex();
    // Se obtiene contenedor div del htm, y es en donde va a alojar toda la informacion 
    const containerPlan = document.getElementById("plans-container");
    // Se recorren cada uno de los planes 
    dataPlans.forEach(plan =>{
        // Condicion para poder controlar que solo se muestren los planes que tienen para un mes
        if (plan.Months == 1) {
            // Guarda informacion del plan con paramtro plan y trae la funcion PlanIndex que se encarga de guardarlas 
            // la informacion correcta en su casilla correcta
            const card = PlanIndex(plan)
            // Creacion de boton para seleccionar plan
            const button = document.createElement("button")
            // Se nombra boton y se guarda en card (donde esta toda la informacion del plan)
            button.textContent = "Elegir plan"
            card.appendChild(button)
            // Se para toda la informacion de la card a containerPlan y se agrega ruta para redireccionar al dar click
            containerPlan.appendChild(card)
            button.addEventListener("click", function () {
                window.location.href = "/CreateCompany";
            });
        }
    })
})
