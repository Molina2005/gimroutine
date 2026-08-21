// FUNCIONES REFERENTES A TODO LO RELACIONADO CON BOTONES DE CRUDS DE PLANES
export function ConsultButtonPlans(id){
    window.location.href = ``
}
export function UpdateButtonPlans(id){
    window.location.href = ``
}
export async function DeleteButtonPlans(id){
// Interfaz de sweetAlert2 para mostrar en pantalla alerta dinamica
    const result = await Swal.fire({
        title: "Estas seguro de eliminar?",
        text: "Plan se eliminara definitivamente!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonText: "Eliminar",
        cancelButtonText: "Cancelar"
    });
    // Si la respuesta es no o saca de la interfaz dinamica
    if (!result.isConfirmed) return;
    // Si da confirmar elimina el plan
    try{
        await ConnectionClientDelete(id);
        await Swal.fire({
            icon:"success",
            title:"Eliminado",
            text:"el cliente fue eliminado correctamente"
        })
    }catch (err){
        Swal.fire({
            icon:"error",
            title:"error",
            text:"No se puede eliminar el cliente"
        })
    }
}