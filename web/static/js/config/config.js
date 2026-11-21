// URLSM PARA PODER MANEJAR PREUEBAS LOCALMENTE Y AMBIENTE DE PRODUCCION
export const API_BACKEND_RAILWAY = window.location.hostname == "localhost"
    ?  "http://localhost:8080"
    : "https://mabtrosoft-backend-production.up.railway.app"