import { VerifyToken } from "../login/tokenVerification.js";
import { Logout } from "../login/tokenLogout.js";

// DOM para poder manejar que al momento que el usaurio de click en cerrar sesion el token se elimine y no lo deje ingresar
document.addEventListener("DOMContentLoaded",()=>{
    VerifyToken();
    const Button = document.getElementById("logoutButtontechnical");
    Button.addEventListener("click", Logout);
})