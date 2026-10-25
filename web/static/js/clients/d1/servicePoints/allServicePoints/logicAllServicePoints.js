import { ConnectionAllServicePoints } from "./api.js"

export async function LogicAllServicePoints(){
    const idRegional = Number(sessionStorage.getItem("idRegional"));
    const dataServicePoints = await ConnectionAllServicePoints(idRegional)
    return dataServicePoints
}