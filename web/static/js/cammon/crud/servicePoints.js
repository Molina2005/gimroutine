import { ConnectionServicePointDelete } from "../../ui/clients/d1/servicePoints/api.js";

// FUNCIONES REFERENTES A TODO LO RELACIONADO CON BOTONES DE CRUDS DE CLIENTES
export function UpdateButtonServicePoints(id){
    window.location.href = `/formUpdateServicePoints?id=${id}`;
}
export async function DeleteButtonServicePoints(id){
// Interfaz de sweetAlert2 para mostrar en pantalla alerta dinamica
    const result = await Swal.fire({
        title: "Estas seguro de eliminar?",
        text: "Punto servicio se eliminara definitivamente!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: "Eliminar",
        cancelButtonText: "Cancelar"
    });
    // Si la respuesta es no o saca de la interfaz dinamica
    if (!result.isConfirmed) return;
    // Si da confirmar elimina el usuario 
    try{
        await ConnectionServicePointDelete(id);
        await Swal.fire({
            icon:"success",
            title:"Eliminado",
            text:"Punto servicio fue eliminado correctamente"
        })
    }catch (err){
        Swal.fire({
            icon:"error",
            title:"error",
            text:"No se puede eliminar el punto servicio"
        })
    }
}