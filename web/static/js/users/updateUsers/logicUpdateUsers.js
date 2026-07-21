import { InfoUser, updateUser } from "./api.js";
// Recibe la informacion del usuario para poder obtener sus datos
export async function logicInfoUser(id){
   // Pasa el id que recibio del mainUpdate al fetch y al mismo tiempo en el momento que 
   // se obtiene la informacion en el fetch llega aqui y pasa al main.js
   const dataUser = await InfoUser(id);
   return dataUser
} 
export async function logicUpdateUser(id,data){
   // Pasa el id que recibio del mainUpdate al fetch y al mismo tiempo en el momento que 
   // se obtiene la informacion en el fetch llega aqui y pasa al main.js
   const dataUser = await updateUser(id, data);
   return dataUser
} 

