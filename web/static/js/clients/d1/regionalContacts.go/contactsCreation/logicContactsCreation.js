import { ConnectionRegionalContactsCreate } from "./api.js"

export async function LogicCreateRegionalContacts(data){
    const res = await ConnectionRegionalContactsCreate(data)
    return res
}