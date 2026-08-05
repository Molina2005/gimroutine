document.addEventListener("DOMContentLoaded", async function(){
    const form = document.getElementById("addClientForm")
    form.addEventListener("submit", async function(e){
        e.preventDefault();
        const formData = {
            name: document.getElementById("name").value,
            document: document.getElementById("document").value,
            gmail:document.getElementById("gmail").value,
            phone: document.getElementById("phone").value,
            password: document.getElementById("password").value,
            state: document.getElementById("state").value
        }
        try{
            const res = await LogicClient(formData);
            alert("Cliente creado correctamente", res)
            form.reset();
        }catch (err){
            if (err.status == "409"){
                alert("Cliente ya existe en el sistema")
            }else{
                alert("Error en el servidor")   
            }
        }
    })
})