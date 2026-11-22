// URLSM PARA PODER MANEJAR PREUEBAS LOCALMENTE, AMBIENTE DE PRUEBAS Y AMBIENTE DE PRODUCCION
const DOMINIO_ACTUAL = window.location.hostname;
export const API_BACKEND_RAILWAY = DOMINIO_ACTUAL == "localhost"
    // LOCAL
    ?  "http://localhost:8080"
    // PRUEBAS
    : DOMINIO_ACTUAL == "mabtrosoftweb-git-testing-programming10.vercel.app"
    ? "https://mabtrosoft-testing-production.up.railway.app"
    // PRODUCCIONN
    : "https://mabtrosoft-backend-production.up.railway.app"