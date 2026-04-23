// Importacion de funcion que trae recorre a los clientes
import { LocationButtons } from "../cammon/locationButtons.js";
import { LogicAllClient } from "../clients/allClients/logicAllClient.js";

document.addEventListener("DOMContentLoaded", async function(){
    const clients = await LogicAllClient();
    // Se obtiene tabla que es en donde se va a guardar la informacion
    const TableBody = document.getElementById("clients-table")
    // Deja en vacio si llega a haber informacion y no agregue info basura
    TableBody.innerHTML = ""
    // Recorre los clientes uno por uno
    clients.forEach(client =>{
        // Contenedor para guardar la info de las celdas td
        const row = document.createElement("tr")
        row.innerHTML = `
            <td>${client.Name}</td>
            <td>${client.Document}<td/>
            <td>${client.Gmail}<td/>
            <td>${client.Phone}<td/>
            <td>${client.EnterDate}<td/>
            <td>${client.State}<td/>`
        // Creacion de celda la cual guarda los botones con sus acciones 
        const actions = document.createElement("td")
        actions.innerHTML = `
            <button class="ConsultButton" data-id="${client.Document}">Consultar</button>
            <button class="UpdateButton" data-id="${client.Document}">Actualizar</button>
            <button class="DeleteButton" data-id="${client.Document}">Eliminar</button>`
        // Se agregan a su respectivo campo
        row.appendChild(actions);
        TableBody.appendChild(row);     
        // Direccionamientos botones acciones
        // Se pasa la clase del boton y el document que es el identificador de cada cliente
        LocationButtons("ConsultButton", (document) =>{
            window.location.href = ""
        });
        LocationButtons("UpdateButton", (document) =>{
            window.location.href = ""
        });
        LocationButtons("DeleteButton", (document) =>{
            window.location.href = ""
        });
    });
});