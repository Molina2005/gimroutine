// Funcion para manejar todo lo relacionado con los buscadores
export async function SearchCreate(IdInput, functionPas){
    const search = document.getElementById(IdInput)
    // evento input para poder ver resultados de busqueda mientras de digita informacion
    search.addEventListener("input", function(){
        // Obtiene el valot y  trim para controlar  espacios al inicio y final de cada palabra 
        const data = search.value.trim();
        if (data === ""){return};
        functionPas(data);
    });
};


