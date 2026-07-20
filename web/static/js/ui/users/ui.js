    import { logicAllUsers } from "../../users/allUsers/logicAllUsers.js"
    import { RenderUsersRow } from "../../cammon/renderDataUsers.js";
    // Mostrar informacion de todos los usuarios con sus respectivos botones de cruds
    document.addEventListener("DOMContentLoaded",async function(){
        const users = await logicAllUsers();
        const tableBody = document.getElementById("users-table");
        tableBody.innerHTML = ""
        users.forEach(user =>{
            tableBody.appendChild(RenderUsersRow(user))
        })
    })