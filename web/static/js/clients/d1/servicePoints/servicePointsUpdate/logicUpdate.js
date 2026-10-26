import { ConnectionServicePointInformation, ConnectionServicePointUpdate } from "./api.js";

export async function LogicServicePointInformation(id){
    const res = await ConnectionServicePointInformation(id)
    return res
}

export async function LogicServicePointsUpdate(id, data){
    const res = await ConnectionServicePointUpdate(id, data)
    return res
}