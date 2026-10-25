import { ConnectionServicePointInformation } from "./api.js";

export async function LogicServicePointsUpdate(id){
    const res = await ConnectionServicePointInformation(id)
    return res
}