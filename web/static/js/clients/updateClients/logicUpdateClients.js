import { InfoClient } from "./api.js";
import { UpdateClient } from "./api.js";
export async function logicInfoClient(id){
   // Pasa el id que recibio del mainUpdate al fetch y al mismo tiempo en el momento que 
   // se obtiene la informacion en el fetch llega aqui y pasa al main.js
   const dataClient = await InfoClient(id);
   return dataClient
} 

export async function logicUpdateClients(id,data){
   // Pasa el id que recibio del mainUpdate al fetch y al mismo tiempo en el momento que 
   // se obtiene la informacion en el fetch llega aqui y pasa al main.js
   const dataClient = await UpdateClient(id, data);
   return dataClient
} 