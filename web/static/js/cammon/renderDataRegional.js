import { ConsultButtonRegional, DeleteButtonRegional, UpdateButtonRegional } from "./crud/regionalCrud.js";
import { RenderCrudButton } from "./locationButtons.js";

// Funcion para creacion de filas - columnas, botones crud 
export function RenderRegionalRow(regional){
    const row = document.createElement("tr");
    const cellName = document.createElement("td");
    cellName.textContent = regional.Name;
    row.appendChild(cellName);
    const cellClient = document.createElement("td");
    cellClient.textContent = regional.Client;
    row.appendChild(cellClient); 
    const cellState = document.createElement("td");
    cellState.textContent = regional.State;
    row.appendChild(cellState);
    row.appendChild(RenderCrudButton(regional.Id,ConsultButtonRegional, UpdateButtonRegional, DeleteButtonRegional))
    return row
}