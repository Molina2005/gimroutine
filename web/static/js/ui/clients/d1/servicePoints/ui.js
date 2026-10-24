
document.addEventListener("DOMContentLoaded", async function(){
    const regional = await LogicAllRegionalContacts();
    // Se obtiene tabla que es en donde se va a guardar la informacion
    const TableBody = document.getElementById("regional-contacts-table")
    // Deja en vacio si llega a haber informacion y no agregue info basura
    TableBody.innerHTML = ""
    // Recorre los clientes uno por uno
    regional.forEach(regionContacts =>{
        TableBody.appendChild(RenderRegionalContactsRow(regionContacts))
    });
});