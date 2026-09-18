import { ConsultButtonUsers, DeleteButtonUsers, UpdateButtonUsers } from "./crud/usersCrud.js";
import { RenderCrudButton } from "./locationButtons.js";
// Funcion para poder renderizar y crear informacion de los usuarios y mostrarlas en pantalla, ademas de los botones de cruds
export function RenderUsersRow(user){
    const row = document.createElement("tr");
    const cellName = document.createElement("td");
    cellName.textContent = user.Name;
    row.appendChild(cellName);
    const cellLastName = document.createElement("td");
    cellLastName.textContent = user.LastName;
    row.appendChild(cellLastName);
    const cellEmail = document.createElement("td");
    cellEmail.textContent = user.Gmail;
    row.appendChild(cellEmail);
    const cellPassword = document.createElement("td");
    cellPassword.textContent = user.Password;
    row.appendChild(cellPassword);
    const cellRole = document.createElement("td");
    cellRole.textContent = user.Role;
    row.appendChild(cellRole);
    const cellPhone = document.createElement("td");
    cellPhone.textContent = user.Phone;
    row.appendChild(cellPhone);
    const cellState = document.createElement("td");
    cellState.textContent = user.Users_state;
    row.appendChild(cellState);
    const cellEntryDate = document.createElement("td");
    cellEntryDate.textContent = user.EntryDate;
    row.appendChild(cellEntryDate);
    row.appendChild(RenderCrudButton(user.Id, ConsultButtonUsers, UpdateButtonUsers, DeleteButtonUsers))
    return row
}