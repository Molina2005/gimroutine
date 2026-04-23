// Funcion para la redireccion de cada boton de acciones de clientes 
// Recibe como parametro, clase del boton y identificador del cliente en este caso document
export async function LocationButtons(btnClass, clientData){
    // Se obtiene cada uno de los botones 
    const buttons = document.querySelectorAll(`.${btnClass}`);
    // se recorren cada uno de los botones y se guardan en btn
    buttons.forEach(btn =>{
        // Se crea evento que al dar click obtiene el id(documento) de ese cliente y luego lo pasa al parametro clientData
        btn.addEventListener("click", ()=>{
            const clientId = btn.dataset.id
            clientData(clientId)
        })
    })
}