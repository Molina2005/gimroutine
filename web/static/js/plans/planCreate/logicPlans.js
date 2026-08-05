import { ConnectionPlans } from "./api";
async function LogicPlans(formData) {
    const data = await ConnectionPlans(formData);
    return data
}