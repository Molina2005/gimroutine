    import { LogicAllClient } from "./logicAllClient.js"
    export async function GetClients() {
        try{
            // Info que viene del logicClient se guarda en res
            const res=await LogicAllClient()
            // Se selecciona el id clients-table
            const tableBody=document.getElementById("clients-table")
            // Se vacia todo el contenedor para añadir lo que viene del backend
            tableBody.innerHTML=""
            // Recorrer cada uno de los clientes
            res.forEach(client =>{
                // ================== Nombre ==================
                // Creacion contenedor para guardar la informacion que esta dentro, contiene las celdas
                const row = document.createElement("tr");
                // Elemento que crea una celda la cual guarda el nombre del cliente
                const clientName = document.createElement("td");
                // Agrega el nombre a la celda ClientName
                clientName.textContent = client.Name;
                // Agrega a row lo que se tiene en clientName
                row.appendChild(clientName);
                // ================== Documento ==================
                const clientDoc = document.createElement("td");
                clientDoc.textContent=client.Document;
                row.appendChild(clientDoc);
                // ================== Correo ==================
                const clientGmail = document.createElement("td");
                clientGmail.textContent=client.Gmail;
                row.appendChild(clientGmail);
                // ================== Telefono ==================
                const clientPhone = document.createElement("td");
                clientPhone.textContent=client.Phone;
                row.appendChild(clientPhone);
                // ================== Fecha ingreso ==================
                const clientEnterDate = document.createElement("td");
                clientEnterDate.textContent=client.EnterDate;
                row.appendChild(clientEnterDate);
                // ================== Estado ==================
                const clientState = document.createElement("td");
                clientState.textContent=client.State;
                row.appendChild(clientState);
                // Agrega a tableBody lo que se tiene en row
                tableBody.appendChild(row);
            })
        }catch (err){
            alert("Error al obtener clientes", err)
        }
    }
    document.addEventListener("DOMContentLoaded",GetClients);