// DOM: Permite interactuar con el html, da acceso a trabajar con el formulario y botones
document.addEventListener("DOMContentLoaded", function(){
    // Varible que obtiene informacion del formulario 
    const form = document.getElementById("userForm") 
    // Evento que al momento de dar oprimir el boton hace todo lo que se contiene dentro de la funcion
    form.addEventListener("submit", function(e){
        // evita el comportamiento por defecto del formulario (recargar la página) y perder la informacion.
        e.preventDefault();
        // se obtiene el elemento name y se usa value para obtener el valor actual del input en ese momento
        const name = document.getElementById("name").value;
        const email = document.getElementById("email").value;
        const age = document.getElementById("age").value;
        const weight = document.getElementById("weight").value;
        const height = document.getElementById("height").value;
        const password = document.getElementById("password").value;
        // Ruta para poder comunicarse con el backend(API) de creacion de usuarios
        fetch("/users",{
            // Metodo con el que se realizara la peticion
            method: "POST",
            // headers: indica el tipo de datos que se envian en el body
            // en este caso son datos en formato json
            headers: {"Content-Type": "application/json"},
            // Cambiar variables con los datos a cadena de texto 
            body: JSON.stringify({
                name, 
                email, 
                age : Number(age), 
                weight: Number(weight), 
                height: Number(height), 
                password}) 
        }) 
        // then: sirve para decir que hacer cuando una promesa se cumple
        // Response: respuesta del servidor 
            // Funcion para poder usar Response
        // Devuelve una promesa con los datos del json 
        .then(response => {
            // Si response es difente a que todo salio bien(ok)
                // muestre el error, si no convierte JSON a objeto JS
            if (!response.ok) { 
                throw new Error("Error al crear usuario");
            }
            return response.text();
        })
        // Data de los datos que se se pasaron a objeto js
        .then(data => {
            console.log("usuario creado", data);
        })
        // Resetear formulario al momento de agregar
        document.getElementById("userForm").reset();
    })
})

