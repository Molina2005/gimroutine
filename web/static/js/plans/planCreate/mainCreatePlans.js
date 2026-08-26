import { CreatePricesPlan } from "../../cammon/prices.js";
import { LogicPlans } from "./logicPlans.js";
document.addEventListener("DOMContentLoaded", async function(){
    // Formulario en el cual se van a digitar los datos
    const form = document.getElementById("addPlansForm");
    // Se guarda la informacion que trae la funcion CreatePricesPlan (maquetacion para nuevo precio con nueva clase)
    const ResultNewPrice =  document.getElementById("content-new-price");
    // Se obtiene boton de html para manejar el evento cuando ususario de click
    // Cuando el usaurio de click se ejecutara la funcion y pasara la informacion a ResultNewPrice
    const buttonNewPrice = document.getElementById("new-price-button");
    buttonNewPrice.addEventListener("click", function(){
        ResultNewPrice.appendChild(CreatePricesPlan());
    });
    form.addEventListener("submit", async function(e){
        e.preventDefault();
        // de la informacion recolectada de la funcion y guardada en ResultNewPrice se obtiene el div creado en esa funcion con la esa clase 
        // y lo que contiene lo que esta dentro de el
        const prices = ResultNewPrice.querySelectorAll(".price-item")
        // Array para guardar los datos que se recorren 
        const dataNewPrice = []
        // Se recorren cada uno de los valores con la clase indicada y se guardan con push
        prices.forEach(priceData => {
            dataNewPrice.push({
                months : Number(priceData.querySelector(".months").value),
                price : Number(priceData.querySelector(".price").value)
            })
        })
        // Se recoleta toda la informacion y se pasa a logicPlans para poder enviarla al siguiente proceso y pueda llegar a
        // procersarla el backend
        const formData = {
            name: document.getElementById("name").value,
            description: document.getElementById("description").value,
            usermax: Number(document.getElementById("usermax").value),
            monthsandprice:dataNewPrice
        }        

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