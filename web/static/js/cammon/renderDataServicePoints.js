import { DeleteButtonServicePoints, UpdateButtonServicePoints } from "./crud/servicePoints.js";
import { RenderCrudButtonEspecials } from "./locationButtons.js";

// Funcion para creacion de filas - columnas, botones crud 
export function RenderServicePointsRow(servicePoint){
    const row = document.createElement("tr");
    const cellName = document.createElement("td");
    cellName.textContent = servicePoint.Name;
    row.appendChild(cellName);
    const cellAddres = document.createElement("td");
    cellAddres.textContent = servicePoint.Addres;
    row.appendChild(cellAddres); 
    const cellCity = document.createElement("td");
    cellCity.textContent = servicePoint.City;
    row.appendChild(cellCity); 
    const cellState = document.createElement("td");
    cellState.textContent = servicePoint.State;
    row.appendChild(cellState);
    row.appendChild(RenderCrudButtonEspecials(servicePoint.Id, UpdateButtonServicePoints, 
        DeleteButtonServicePoints))
    return row
}