import { ConnectionAllClient } from "./api.js";
export async function LogicAllClient(formData){
    const data = await ConnectionAllClient(formData);
    return data
}