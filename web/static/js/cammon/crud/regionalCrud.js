import { ConnectionRegionalDelete } from "../../ui/clients/d1/regional/api.js";

// FUNCIONES REFERENTES A TODO LO RELACIONADO CON BOTONES DE CRUDS DE CLIENTES
export async function ConsultButtonRegional(id){
    // Se guarda id para reutilizarlo en paginas posteriores
    sessionStorage.setItem("idRegional", id)
    window.location.href = "/moduleMenu";
}
export function UpdateButtonRegional(id){
    window.location.href = `/formUpdateRegional?id=${id}`;
}
export async function DeleteButtonRegional(id){
// Interfaz de sweetAlert2 para mostrar en pantalla alerta dinamica
    const result = await Swal.fire({
        title: "Estas seguro de eliminar?",
        text: "regional se eliminara definitivamente!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: "Eliminar",
        cancelButtonText: "Cancelar"
    });
    // Si la respuesta es no o saca de la interfaz dinamica
    if (!result.isConfirmed) return;
    // Si da confirmar elimina el usuario 
    try{
        await ConnectionRegionalDelete(id);
        await Swal.fire({
            icon:"success",
            title:"Eliminar",
            text:"la regional fue eliminada correctamente"
        })
    }catch (err){
        Swal.fire({
            icon:"error",
            title:"error",
            text:"No se puede eliminar la regional"
        })
    }
}