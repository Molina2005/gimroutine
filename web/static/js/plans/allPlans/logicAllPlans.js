import { ConnectionAllPlans } from "./api.js";
export async function logicAllPlans(){
    const data = await ConnectionAllPlans();
    return data 
}