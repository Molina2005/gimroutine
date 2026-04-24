import { RenderClients } from "../cammon/clientsConsult";
import { LogicAllClient } from "../clients/allClients/logicAllClient";

document.addEventListener("DOMContentLoaded", async function(){
    const clients = await LogicAllClient()
    RenderClients(clients)
})