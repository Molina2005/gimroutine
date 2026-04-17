import { ConnectionClient } from "./api.js";
export async function LogicClient(formData) {
    const data = await ConnectionClient(formData);
    return data
}