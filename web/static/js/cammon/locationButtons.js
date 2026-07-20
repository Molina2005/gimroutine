import { ConnectionUsersDelete } from "../ui/users/api.js";
// FUNCIONES REFERENTES A TODO LO RELACIONADO CON BOTONES DE CRUDS
export function ConsultButton(id){
    window.location.href = `/PrivacyPolicy`
}
export function UpdateButton(id){
    window.location.href = `/FormUpdateUsers?id=${id}`
}
export async function DeleteButton(id){
// Interfaz de sweetAlert2 para mostrar en pantalla alerta dinamica
    const result = await Swal.fire({
        title: "Estas seguro de eliminar?",
        text: "Usuario se eliminara definitivamente!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: "Eliminar",
        cancelButtonText: "Cancelar"
    });
    // Si la respuesta es no o saca de la interfaz dinamica
    if (!result.isConfirmed) return;
    // Si da confirmar elimina el usuario 
    try{
        await ConnectionUsersDelete(id);
        await Swal.fire({
            icon:"success",
            title:"Eliminado",
            text:"el texto fue eliminado correctamente"
        })
    }catch (err){
        Swal.fire({
            icon:"error",
            title:"error",
            text:"No se puede eliminar el usuario"
        })
    }
}
// Funcion para poder guardar los botones de cruds y asi puedan ser reutilizables
export function RenderCrudButton(id){
    // Container el cua va a guardar toda la informacion
    const buttonContainer = document.createElement("td")
    // Botones para cada una de las cruds 
    const consultButton = document.createElement("button")
    consultButton.textContent = "Consultar"
    consultButton.onclick = () => ConsultButton(id)
    const updateButton = document.createElement("button")
    updateButton.textContent = "Actualizar"
    updateButton.onclick = () => UpdateButton(id)
    const deleteButton = document.createElement("button")
    deleteButton.textContent = "Eliminar"
    deleteButton.onclick = () => DeleteButton(id)
    // Se agrega botones a container y retorna cada uno de ellos 
    buttonContainer.appendChild(consultButton)
    buttonContainer.appendChild(updateButton)
    buttonContainer.appendChild(deleteButton)
    return buttonContainer
}

