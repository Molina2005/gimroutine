import { LogicAllClient } from "../../clients/allClients/logicAllClient.js";
import { RenderClientsRow } from "../../cammon/renderDataClients.js";

document.addEventListener("DOMContentLoaded", async function(){
    const clients = await LogicAllClient();
    // Se obtiene tabla que es en donde se va a guardar la informacion
    const TableBody = document.getElementById("clients-table")
    // Deja en vacio si llega a haber informacion y no agregue info basura
    TableBody.innerHTML = ""
    // Recorre los clientes uno por uno
    clients.forEach(client =>{
        TableBody.appendChild(RenderClientsRow(client))
    });
});