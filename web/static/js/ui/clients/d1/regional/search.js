import { RenderRegionalRow } from "../../../../cammon/renderDataRegional.js";
import { SearchCreate } from "../../../../cammon/search.js";
import { ConnectionRegionalSearch } from "./api.js";

// Envio de informacion de lo que el usuario digito en el buscador al fetch para que pueda procesarla con el backend
document.addEventListener("DOMContentLoaded", async function(){
    SearchCreate("searchInputRegional", ConnectionRegionalSearch);
})

// Funcion que recibe la informacion que envia fetch y crea las columnas
export async function ReturnSearchRegional(data){
    const tableSearch = document.getElementById("regional-table")
    tableSearch.innerHTML = ""
    data.forEach((regional) =>{
        tableSearch.appendChild(RenderRegionalRow(regional));
    })
}