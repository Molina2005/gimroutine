import { env } from "dotenv";
env()
const domainLocal = process.env.DOMAIN_LOCAL_WEB
console.log(domainLocal)
const domainTesting = process.env.DOMAIN_MOMENT_CURRECT
console.log(domainTesting)
const domainRailway = process.env.DOMAIN_TESTING_RAILWAY
const domainProduction = process.env.MAIN_PRODUCTION_WEB
// URLSM PARA PODER MANEJAR PREUEBAS LOCALMENTE, AMBIENTE DE PRUEBAS Y AMBIENTE DE PRODUCCION
const DOMINIO_ACTUAL = window.location.hostname;
export const API_BACKEND_RAILWAY = DOMINIO_ACTUAL == "localhost"
    // LOCAL
    ?  domainLocal
    // PRUEBAS
    : DOMINIO_ACTUAL == domainTesting
    ? domainRailway
    // PRODUCCIONN
    : domainProduction