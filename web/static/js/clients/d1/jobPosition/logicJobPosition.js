import { ConnectionAllJobPosition } from "./api.js"

// Funcion que recibe los datos que obtiene fetch y los almacena aqui
export async function LogicAllJobPosition() {
    const data = await ConnectionAllJobPosition()
    return data
}