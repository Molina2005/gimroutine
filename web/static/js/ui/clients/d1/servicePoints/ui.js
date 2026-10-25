import { RenderServicePointsRow } from "../../../../cammon/renderDataServicePoints.js";
import { LogicAllServicePoints } from "../../../../clients/d1/servicePoints/allServicePoints/logicAllServicePoints.js";

document.addEventListener("DOMContentLoaded", async function(){
    const servicePoints = await LogicAllServicePoints();
    // Se obtiene tabla que es en donde se va a guardar la informacion
    const TableBody = document.getElementById("service-points-table")
    // Deja en vacio si llega a haber informacion y no agregue info basura
    TableBody.innerHTML = ""
    // Recorre los puntos de servicio uno por uno
    servicePoints.forEach(servicePoint =>{
        TableBody.appendChild(RenderServicePointsRow(servicePoint))
    });
});