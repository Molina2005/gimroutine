import { SearchCreate } from "../../cammon/search.js";
import { ConnectionSearchUser } from "./api.js";
import { RenderUsersRow } from "../../cammon/renderDataUsers.js";
// Envio de informacion de lo que el usuario digito en el buscador al fetch para que pueda procesarla con el backend
document.addEventListener("DOMContentLoaded", async function(){
    SearchCreate("searchInputUser", ConnectionSearchUser);
})
// Funcion que recibe la informacion que envia fetch y crea las columnas
export async function DataReturn(data){
    const tableSearch = document.getElementById("users-table")
    tableSearch.innerHTML = ""
    data.forEach((users) =>{
        tableSearch.appendChild(RenderUsersRow(users));
    })
}


