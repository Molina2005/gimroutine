import { LogicPlans } from "./logicPlans.js";
document.addEventListener("DOMContentLoaded", async function(){
    const form = document.getElementById("addPlansForm")
    form.addEventListener("submit", async function(e){
        e.preventDefault();
        const formData = {
            name: document.getElementById("name").value,
            description: document.getElementById("description").value,
            usermax: Number(document.getElementById("usermax").value),
            months: Number(document.getElementById("months").value),
            price: Number(document.getElementById("price").value)
        }
        buttonNewPrice = document.getElementById("new-price")
        buttonNewPrice.addEventListener("click", function(){
            // Estrucutta para poder imgresar nuevo precio
            // y que pueda ingesar esos datos y procesarlos como un nuevo precio
        })
        try{
            const res = await LogicPlans(formData);
            alert("Plan creado correctamente", res)
            form.reset();
        }catch (err){
            if (err.status == "409"){
                alert("Plan ya existe en el sistema")
            }else{
                alert("Error en el servidor")   
            }
        }
    })
})