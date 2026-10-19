import { ConnectionCreateRegional } from "./api.js";

export async function LogicCreateRegional(data){
    const res = await ConnectionCreateRegional(data)
    return res
}