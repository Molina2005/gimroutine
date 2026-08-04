import { ConnectionUsersDelete } from "../../ui/users/api.js";
// FUNCIONES REFERENTES A TODO LO RELACIONADO CON BOTONES DE CRUDS DE USUARIOS
export function ConsultButtonUsers(id){
    window.location.href = `/PrivacyPolicy`
}
export function UpdateButtonUsers(id){
    window.location.href = `/FormUpdateUsers?id=${id}`
}
export async function DeleteButtonUsers(id){
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
            text:"el usuario fue eliminado correctamente"
        })
    }catch (err){
        Swal.fire({
            icon:"error",
            title:"error",
            text:"No se puede eliminar el usuario"
        })
    }
}