import { ConsultButtonClients, DeleteButtonClients, UpdateButtonClients } from "./crud/clientsCrud.js";
import { RenderCrudButton } from "./locationButtons.js";
// Funcion para creacion de filas - columnas, botones crud 
export function RenderClientsRow(client){
    const row = document.createElement("tr");
    const cellName = document.createElement("td");
    cellName.textContent = client.Name;
    row.appendChild(cellName);
    const cellDocument = document.createElement("td");
    cellDocument.textContent = client.Document;
    row.appendChild(cellDocument);
    const cellEmail = document.createElement("td");
    cellEmail.textContent = client.Gmail;
    row.appendChild(cellEmail);
    const cellPhone = document.createElement("td");
    cellPhone.textContent = client.Phone;
    row.appendChild(cellPhone);
    const cellEntryDate = document.createElement("td");
    cellEntryDate.textContent = client.EntryDate;
    row.appendChild(cellEntryDate);
    const cellState = document.createElement("td");
    cellState.textContent = client.State;
    row.appendChild(cellState);
    row.appendChild(RenderCrudButton(client.Id, ConsultButtonClients, UpdateButtonClients, DeleteButtonClients))
    return row
}