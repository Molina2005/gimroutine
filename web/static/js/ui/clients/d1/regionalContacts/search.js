import { RenderRegionalContactsRow } from "../../../../cammon/renderDataRegionalContacts.js";
import { KeySearchCreate } from "../../../../cammon/search.js";
import { ConnectionRegionalContactSearch } from "./api.js";

// Envio de informacion de lo que el usuario digito en el buscador al fetch para que pueda procesarla con el backend
// ademas se pasa idregional para que solo procese la informacion correspondiente y no la de otras regionales
document.addEventListener("DOMContentLoaded", async function(){
    const id_regional = Number(sessionStorage.getItem("idRegional"));
    KeySearchCreate("searchInputContact", ConnectionRegionalContactSearch, id_regional);
})

// Funcion que recibe la informacion que envia fetch y crea las columnas
export async function ReturnSearchRegionalContact(data){
    const tableSearch = document.getElementById("regional-contacts-table")
    tableSearch.innerHTML = ""
    data.forEach(regionalContact =>{
        tableSearch.appendChild(RenderRegionalContactsRow(regionalContact));
    })
}