import { StateUserOptions } from "../../../../cammon/selectOption.js";
import { logicAllUserStates } from "../../../../user_states/logicAllUserStates.js";
import { LogicCreateServicePoints } from "./logicCreate.js";

// Funcion para obtener informacion de formulario y crear regional
document.addEventListener("DOMContentLoaded", async function(){
    // Se obtiene id del sessionStorage para poder enviar esa creacion con id de la reginal seleccionada
    const id_regional = Number(sessionStorage.getItem("idRegional"));
    // 
    const dataRegionalState = await logicAllUserStates()
    const regionalStateSelect = document.getElementById("state")

    // Se recorre cada uno de los clientes y estados existentes y los muestra en la web
    dataRegionalState.forEach(state =>{
        StateUserOptions(state, regionalStateSelect)
    })
    // Varible que obtiene informacion del formulario 
    const form = document.getElementById("form-servicepoints-create")
    // Evento que al momento de dar oprimir el boton hace todo lo que se contiene dentro de la funcion
    form.addEventListener("submit", async function(e){
        // evita el comportamiento por defecto del formulario (recargar la página) y perder la informacion.
        e.preventDefault();
        // FormData va a guardar toda la informacion que se obtega dentro de esa const
        const formData = {
            // se obtiene los elementos y se usa value para obtener el valor actual del input en ese momento
            name: document.getElementById("name").value,
            addres: document.getElementById("addres").value,
            city: document.getElementById("city").value,
            id_regional: id_regional,
            id_state: Number(document.getElementById("state").value),
        }
        try{
            const data = await LogicCreateServicePoints(formData);
            alert("Punto servicio creada correctamente",data);
            form.reset();
        }catch (err){
            if (err.status == "409"){
                alert("Punto servicio ya existe en el sistema")
            }else{
                alert("Error en el servidor")
            }
        }
    });
});

