import { ConnectionAllRegional } from "./api.js";
// Funcion que recibe los datos que obtiene fetch y los almacena aqui
export async function LogicAllRegional() {
    const data = await ConnectionAllRegional()
    return data
}