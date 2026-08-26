export function CreatePricesPlan(){
    const price = document.createElement("div")

    price.className = "price-item"

    price.innerHTML  = `
        <input type="number" class="months" placeholder="Meses"> 
        <input type="number" class="price" placeholder="Valor"> 
    `
    return price
}