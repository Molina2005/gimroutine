import { ConnectionAllUserStates } from "./api.js";

export async function logicAllUserStates(){
    const data = await ConnectionAllUserStates();
    return data 
}