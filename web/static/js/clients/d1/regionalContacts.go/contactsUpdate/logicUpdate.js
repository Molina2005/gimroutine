import { ConnectionInfoRegionalContacts, ConnectionUpdateRegionalContacts } from "./api.js"

export async function LogicInfoRegionalContacts(id){
    const data = await ConnectionInfoRegionalContacts(id);
    return data
}

export async function LogicUpdateRegionalContacts(id, data){
    const dataRegionalContact = await ConnectionUpdateRegionalContacts(id, data);
    return dataRegionalContact
}