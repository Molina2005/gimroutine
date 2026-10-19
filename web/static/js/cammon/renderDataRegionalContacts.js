import { DeleteButtonRegionalContacts, UpdateButtonRegionalContacts } from "./crud/regionalContacts.js";
import { RenderCrudButtonEspecials } from "./locationButtons.js";

// Funcion para creacion de filas - columnas, botones crud 
export function RenderRegionalContactsRow(regionalContacts){
    const row = document.createElement("tr");
    const cellName = document.createElement("td");
    cellName.textContent = regionalContacts.Name;
    row.appendChild(cellName);
    const cellLastName = document.createElement("td");
    cellLastName.textContent = regionalContacts.LastName;
    row.appendChild(cellLastName); 
    const cellGmail = document.createElement("td");
    cellGmail.textContent = regionalContacts.Gmail;
    row.appendChild(cellGmail);
    const cellPhone = document.createElement("td");
    cellPhone.textContent = regionalContacts.Phone;
    row.appendChild(cellPhone);
    const cellState = document.createElement("td");
    cellState.textContent = regionalContacts.Job_position;
    row.appendChild(cellState);
    const cellCargo = document.createElement("td");
    cellCargo.textContent = regionalContacts.UserState;
    row.appendChild(cellCargo);
    row.appendChild(RenderCrudButtonEspecials(regionalContacts.Id, 
        UpdateButtonRegionalContacts, DeleteButtonRegionalContacts))
    return row
}