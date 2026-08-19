import { ConnectionPlans } from "./api.js";
export async function LogicPlans(formData) {
    const data = await ConnectionPlans(formData);
    return data
}