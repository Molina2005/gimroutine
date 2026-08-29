// Funcion para poder crear apartados en donde se van a registrar los datos del plan
export function PlanIndex(plan){
    // Creacion de div y se le añade una clase a cada contenedor segun la cantidad de planes
    const card = document.createElement("div");
    card.className = "plan-card";
    // Creacion de cart con sus datos
    const PlanName = document.createElement("h3");
    PlanName.className = "plan-name";
    PlanName.textContent = plan.Name
    const PlanDescription = document.createElement("h3");
    PlanDescription.className = "plan-description";
    PlanDescription.textContent = plan.Description
    const PlanUserMax = document.createElement("h3");
    PlanUserMax.className = "plan-usermax";
    PlanUserMax.textContent = `${plan.UserMax} Usuarios`
    const PlanMonths = document.createElement("h3");
    PlanMonths.className = "plan-months";
    PlanMonths.textContent = `${plan.Months} / Meses`
    const PlanPrice = document.createElement("h3");
    PlanPrice.className = "plan-price";
    PlanPrice.textContent = plan.Price
    // Se pasan cards con informacion al card donde se van a alojar todos para luego
    // pasarlos al containerPlan para que este lo muestre en pantalla
    card.appendChild(PlanName)
    card.appendChild(PlanDescription)
    card.appendChild(PlanUserMax)
    card.appendChild(PlanMonths)
    card.appendChild(PlanPrice)
    // Retorna la card con toda la informacion
    return card
}