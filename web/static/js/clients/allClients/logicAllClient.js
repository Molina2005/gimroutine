import { ConnectionAllClient } from "./api.js";
export async function LogicAllClient(){
    const data = await ConnectionAllClient();
    return data
}