import { RenderRegionalRow } from "../../../../cammon/renderDataRegional.js";
import { LogicAllRegional } from "../../../../clients/d1/regional/allRegional/logicAllregional.js";

document.addEventListener("DOMContentLoaded", async function(){
    const regional = await LogicAllRegional();
    // Se obtiene tabla que es en donde se va a guardar la informacion
    const TableBody = document.getElementById("regional-table")
    // Deja en vacio si llega a haber informacion y no agregue info basura
    TableBody.innerHTML = ""
    // Recorre los clientes uno por uno
    regional.forEach(region =>{
        TableBody.appendChild(RenderRegionalRow(region))
    });
});