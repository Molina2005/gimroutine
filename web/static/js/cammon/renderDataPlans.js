import { ConsultButtonPlans, DeleteButtonPlans, UpdateButtonPlans } from "./crud/plansCrud.js";
import { RenderCrudButton } from "./locationButtons.js";
// Funcion para creacion de filas - columnas, botones crud 
export function RenderPlansRow(plans){
    const row = document.createElement("tr");
    const cellName = document.createElement("td");
    cellName.textContent = plans.Name;
    row.appendChild(cellName);
    const cellDescription = document.createElement("td");
    cellDescription.textContent = plans.Description;
    row.appendChild(cellDescription);
    const cellUserMax = document.createElement("td");
    cellUserMax.textContent = plans.UserMax;
    row.appendChild(cellUserMax);
    const cellDateCreation = document.createElement("td");
    cellDateCreation.textContent = plans.CretionDate;
    row.appendChild(cellDateCreation);
    row.appendChild(RenderCrudButton(plans.Id, ConsultButtonPlans, UpdateButtonPlans, DeleteButtonPlans))
    return row
}