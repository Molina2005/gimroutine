import { ConnectionAllUsers } from "./api.js";
export async function logicAllUsers(){
    const data = await ConnectionAllUsers();
    console.log(data)
    return data 
}