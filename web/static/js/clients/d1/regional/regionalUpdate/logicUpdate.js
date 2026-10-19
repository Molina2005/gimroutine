import { ConnectionInfoRegional, ConnectionUpdateRegional } from "./api.js";

// Funcion para obtener id de regional
export async function LogicInfoRegional(id){
    const data = await ConnectionInfoRegional(id)
    return data
}

// Funcion para mantener informacion de regional a actualizar
export async function LogicUpdateRegional(id,data){
    const dataRegional = await ConnectionUpdateRegional(id, data)
    return dataRegional
}