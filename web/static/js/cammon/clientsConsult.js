export function RenderClients(list){
    const tableBody = await document.createElement("clients-table");
    tableBody.innerHTML = "";
    list.forEach(client =>{
        const row = document.createElement("tr")
        row.innerHTML = `
            <td>${client.Name}</td>
            <td>${client.Document}</td>
            <td>${client.Gmail}</td>
            <td>${client.Phone}</td>
            <td>${client.EnterDate}</td>
            <td>${client.State}</td>`
        tableBody.appendChild(row);
    });
}