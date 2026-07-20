import { VerifyToken } from "../users/loginUser/tokenVerification.js";
import { Logout } from "../users/loginUser/tokenLogout.js";

// DOM para poder manejar que al momento que el usaurio de click en cerrar sesion el token se elimine y no lo deje ingresar
document.addEventListener("DOMContentLoaded",()=>{
    VerifyToken();
    const Button = document.getElementById("logoutButtonTechnical");
    Button.addEventListener("click", Logout);
})