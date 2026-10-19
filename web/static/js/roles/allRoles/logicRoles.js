import { connectionAllRoles } from "./api.js";
// Funcion que obtiene datos de los roles existentes y asi poder pasarlos para la creacion de usuarios
export async function LogicAllRoles(){
    const dataRoles  = await connectionAllRoles()
    return dataRoles
}