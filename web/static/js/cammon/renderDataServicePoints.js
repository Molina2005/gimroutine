import { ConsultButtonServicePoints, DeleteButtonServicePoints, UpdateButtonServicePoints } from "./crud/servicePoints.js";
import { RenderCrudButton } from "./locationButtons.js";

// Funcion para creacion de filas - columnas, botones crud 
export function RenderServicePointsRow(servicePoints){
    const row = document.createElement("tr");
    const cellName = document.createElement("td");
    cellName.textContent = servicePoints.Name;
    row.appendChild(cellName);
    const cellAddres = document.createElement("td");
    cellAddres.textContent = servicePoints.Addres;
    row.appendChild(cellAddres); 
    const cellCity = document.createElement("td");
    cellCity.textContent = servicePoints.City;
    row.appendChild(cellCity); 
    const cellState = document.createElement("td");
    cellState.textContent = servicePoints.State;
    row.appendChild(cellState); 
    row.appendChild(RenderCrudButton(servicePoints.Id, ConsultButtonServicePoints, UpdateButtonServicePoints, 
        DeleteButtonServicePoints))
    return row
}