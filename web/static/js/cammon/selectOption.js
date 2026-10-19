// Funcion para controlar estados en general, crea cada una de las opciones
export function RolOptions(data,Select){
    const option = document.createElement("option");
    option.value = data.Id_rol;
    option.textContent = data.Post;
    Select.appendChild(option);
    return option
}

export function StateUserOptions(data,Select){
    const option = document.createElement("option");
    option.value = data.Id;
    option.textContent = data.Name;
    Select.appendChild(option);
    return option
}

export function StateClientOptions(data, Select){
    const option = document.createElement("option");
    option.value = data.Id;
    option.textContent = data.Name;
    Select.appendChild(option);
    return option
}

export function StateJobPosition(data, Select){
    const option = document.createElement("option");
    option.value = data.Id;
    option.textContent = data.Name;
    Select.appendChild(option);
    return option
}

