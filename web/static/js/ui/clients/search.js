import { ConnectionSearchClient } from "./api.js";
import { RenderClientsRow } from "../../cammon/renderDataClients.js";
import { SearchCreate } from "../../cammon/search.js";
// Envio de informacion de lo que el usuario digito en el buscador al fetch para que pueda procesarla con el backend
document.addEventListener("DOMContentLoaded", async function(){
    SearchCreate("searchInputClient", ConnectionSearchClient);
})
// Funcion que recibe la informacion que envia fetch y crea las columnas
export async function ReturnSearchClient(data){
    const tableSearch = document.getElementById("clients-table")
    tableSearch.innerHTML = ""
    data.forEach((clients) =>{
        tableSearch.appendChild(RenderClientsRow(clients));
    })
}