import { RenderCrudButton } from "./locationButtons.js";
// Funcion para poder renderizar y crear informacion de los usuarios y mostrarlas en pantalla, ademas de los botones de cruds
export function RenderUsersRow(user){
    const row = document.createElement("tr");
    const cellName = document.createElement("td");
    cellName.textContent = user.Name;
    row.appendChild(cellName);
    const cellEmail = document.createElement("td");
    cellEmail.textContent = user.Gmail;
    row.appendChild(cellEmail);
    const cellPassword = document.createElement("td");
    cellPassword.textContent = user.Password;
    row.appendChild(cellPassword);
    const cellEntryDate = document.createElement("td");
    cellEntryDate.textContent = user.EntryDate;
    row.appendChild(cellEntryDate);
    const cellRole = document.createElement("td");
    cellRole.textContent = user.Role;
    row.appendChild(cellRole);
    row.appendChild(RenderCrudButton(user.Id))
    return row
}