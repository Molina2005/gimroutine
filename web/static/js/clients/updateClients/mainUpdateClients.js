import { logicInfoClient } from "./logicUpdateClients.js";
import { logicUpdateClients } from "./logicUpdateClients.js";
document.addEventListener("DOMContentLoaded", async function(){
    // Obtiene la URL en que se encuentra actualmennte
    const url = new URL(window.location.href);
    // Busca el valor id dentro de la url
    const id = url.searchParams.get("id");
    // Se para el id a la funcion de logic, de igual manera aqui recibe la informacion cuando llega de logic
    const data = await logicInfoClient(id);
    const formUpdate = document.getElementById("form-update-clients")
    document.getElementById("name").value = data.client.Name
    document.getElementById("document").value = data.client.Document
    document.getElementById("gmail").value = data.client.Gmail
    document.getElementById("phone").value = data.client.Phone
    document.getElementById("state").value = data.client.State      
    // Evento para cuando el usuario de click a boton actualizar se actualizen los datos con lo que ingreso en los inputs 
    formUpdate.addEventListener("submit", async function(e){
        e.preventDefault();
        // Obtiene la informacion de que lo que envian los inputs
        const dataClient = {
            Name : document.getElementById("name").value,
            Document : document.getElementById("document").value,
            Gmail : document.getElementById("gmail").value,
            Phone : document.getElementById("phone").value,
            Password : document.getElementById("password").value,
            State : document.getElementById("state").value
        };
        // Retorna la informacion a la siguiente funcion para que cumpla su proceso
        const dataUpdate = await logicUpdateClients(id,dataClient);
        alert("cliente actualizado correctamente")
        window.location.replace("/clientsInfo")
        return dataUpdate
    })
})

