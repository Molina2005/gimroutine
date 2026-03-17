import 'package:flutter/material.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({super.key});
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFFF5F6FA), // Fondo claro
      body: Center(
        child: SingleChildScrollView(
          padding: const EdgeInsets.symmetric(horizontal: 32, vertical: 24),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              //logo
              Image.asset("assets/logo.png"),
              const SizedBox(height: 32),
              // Título
              const Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  "Iniciar sesion",
                  style: TextStyle(
                    fontSize: 28,
                    fontWeight: FontWeight.bold,
                    color: Colors.black87,
                  ),
                ),
              ),
              const SizedBox(height: 8),
              const SizedBox(height: 24),
              // Email
              const Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  "Correo",
                  style: TextStyle(
                    color: Color.fromARGB(255, 77, 0, 210),
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
              TextField(
                decoration: InputDecoration(
                  hintText: "ejemplo@egmail.com",
                  hintStyle: const TextStyle(fontWeight: FontWeight.bold),
                  border: const UnderlineInputBorder(),
                  focusedBorder: const UnderlineInputBorder(
                    borderSide: BorderSide(color: Color.fromARGB(255, 77, 0, 210)),
                  ),
                ),
              ),
              const SizedBox(height: 16),
              // Password
              const Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  "Contraseña",
                  style: TextStyle(
                    color: Color.fromARGB(255, 77, 0, 210),
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
              TextField(
                obscureText: true,
                decoration: InputDecoration(
                  suffixIcon: Icon(Icons.visibility_off_outlined, color: Colors.grey[400]),
                  border: const UnderlineInputBorder(),
                  focusedBorder: const UnderlineInputBorder(
                    borderSide: BorderSide(color: Color.fromARGB(255, 77, 0, 210)),
                  ),
                ),
              ),
              const SizedBox(height: 28),
              // Botón Sign in
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Color.fromARGB(255, 77, 0, 210),
                    padding: const EdgeInsets.symmetric(vertical: 16),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(6),
                    ),
                    elevation: 2,
                  ),
                  onPressed: () {},
                  child: const Text(
                    "Ingresar",
                    style: TextStyle(fontSize: 18, color: Colors.white),
                  ),
                ),
              ),
              const SizedBox(height: 16),
              // Texto social
              const Text(
                "Ingresa con tu cuenta de google",
                style: TextStyle(color: Colors.grey, fontSize: 13),
              ),
              const SizedBox(height: 12),
              // Botone social
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Expanded(
                    child: ElevatedButton.icon(
                      style: ElevatedButton.styleFrom(
                        backgroundColor: const Color(0xFF1DA1F2),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(6),
                        ),
                        elevation: 0,
                      ),
                      onPressed: () {},
                      icon: const Icon(Icons.alternate_email, color: Colors.white),
                      label: const Text("Google", style: TextStyle(color: Colors.white)),
                    ),
                  ),
                  const SizedBox(width: 12),
                ],
              ),
              const SizedBox(height: 18),
              // Enlaces
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  TextButton(
                    onPressed: () {},
                    child: const Text(
                      "Olvidaste la contraseña?",
                      style: TextStyle(color: Colors.grey),
                    ),
                  ),
                  TextButton(
                    onPressed: () {},
                    child: const Text(
                      "Crear cuenta",
                      style: TextStyle(color: Color.fromARGB(255, 77, 0, 210), fontWeight: FontWeight.bold),
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}