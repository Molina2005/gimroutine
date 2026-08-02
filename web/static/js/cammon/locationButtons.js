// Funcion para poder guardar los botones de cruds y asi puedan ser reutilizables
export function RenderCrudButton(id, consult, update, remove){
    // Container el cua va a guardar toda la informacion
    const buttonContainer = document.createElement("td")
    // Botones para cada una de las cruds 
    const consultButton = document.createElement("button")
    consultButton.textContent = "Consultar"
    consultButton.onclick = () => consult(id)
    const updateButton = document.createElement("button")
    updateButton.textContent = "Actualizar"
    updateButton.onclick = () => update(id)
    const deleteButton = document.createElement("button")
    deleteButton.textContent = "Eliminar"
    deleteButton.onclick = () => remove(id)
    // Se agrega botones a container y retorna cada uno de ellos 
    buttonContainer.appendChild(consultButton)
    buttonContainer.appendChild(updateButton)
    buttonContainer.appendChild(deleteButton)
    return buttonContainer
}