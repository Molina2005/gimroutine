import { logicInfoUser } from "./logicUpdateUsers.js"
import { logicUpdateUser } from "./logicUpdateUsers.js"
document.addEventListener("DOMContentLoaded", async function(){
    // Obtiene la URL en que se encuentra actualmennte
    const url = new URL(window.location.href);
    // Busca el valor id dentro de la url
    const id = url.searchParams.get("id");
    // Se para el id a la funcion de logic, de igual manera aqui recibe la informacion cuando llega de logic
    const data = await logicInfoUser(id);
    const formUpdate = document.getElementById("form-update-users")
    document.getElementById("name").value = data.user.Name
    document.getElementById("email").value = data.user.Email
    document.getElementById("role").value = data.user.Role      
    // Evento para cuando el usuario de click a boton actualizar se actualizen los datos con lo que ingreso en los inputs 
    formUpdate.addEventListener("submit", async function(e){
        e.preventDefault();
        // Obtiene la informacion de que lo que envian los inputs
        const dataUser = {
            Name : document.getElementById("name").value,
            Email : document.getElementById("email").value,
            Password : document.getElementById("password").value,
            Role : document.getElementById("role").value
        };
        // Retorna la informacion a la siguiente funcion para que cumnpla su proceso
        const dataUpdate = await logicUpdateUser(id,dataUser);
        alert("usuario actualizado correctamente");
        window.location.replace("/usersInfo");
        return dataUpdate
    })
})

