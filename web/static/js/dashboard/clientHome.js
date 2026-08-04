import { Logout } from "../login/tokenLogout.js";
import { VerifyToken } from "../login/tokenVerification.js";

// DOM para poder manejar que al momento que el cliente de click en cerrar sesion el token se elimine y no lo deje ingresar
document.addEventListener("DOMContentLoaded",()=>{
    VerifyToken();
    const Button = document.getElementById("logoutButtonClient");
    Button.addEventListener("click", Logout);
})