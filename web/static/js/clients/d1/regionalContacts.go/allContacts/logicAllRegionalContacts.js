import { ConnectionAllRegionalContacts } from "./api.js"

// Funcion que recibe los datos que obtiene fetch y los almacena aqui
export async function LogicAllRegionalContacts() {
    const id_regional = Number(sessionStorage.getItem("idRegional"));
    const data = await ConnectionAllRegionalContacts(id_regional)
    return data
}
