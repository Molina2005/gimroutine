import { ConnectionRegionalContactDelete } from "../../ui/clients/d1/regionalContacts/api.js";

// FUNCIONES REFERENTES A TODO LO RELACIONADO CON BOTONES DE CRUDS DE CONTACTOS
export function UpdateButtonRegionalContacts(id){
    window.location.href = `/formUpdateRegionalContacts?id=${id}`;
}
export async function DeleteButtonRegionalContacts(id){
// // Interfaz de sweetAlert2 para mostrar en pantalla alerta dinamica
    const result = await Swal.fire({
        title: "Estas seguro de eliminar?",
        text: "Contacto regional se eliminara definitivamente!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: "Eliminar",
        cancelButtonText: "Cancelar"
    });
    // Si la respuesta es no o saca de la interfaz dinamica
    if (!result.isConfirmed) return;
    // Si da confirmar elimina el usuario 
    try{
        await ConnectionRegionalContactDelete(id);
        await Swal.fire({
            icon:"success",
            title:"Eliminar",
            text:"Contacto fue eliminada correctamente"
        })
    }catch (err){
        Swal.fire({
            icon:"error",
            title:"error",
            text:"No se puede eliminar contacto"
        })
    }
}