// Se obtiene el formulario html por medio del atributo id
// En cuaanto se envie el formulario se ejecuta todo lo que esta dentro de la funcion
document.getElementById("userForm").addEventListener("submit", function(e) {
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
        headers: {"Content-Type": "aplication/json"},
        // body: JSON.stringify({name, email, age, weight, height, password})
    })
})