import { ConnectionCreateServicePoints } from "./api.js"

export async function LogicCreateServicePoints(data){
    const res = await ConnectionCreateServicePoints(data)
    return res
}