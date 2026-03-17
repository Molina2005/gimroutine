import 'package:flutter/material.dart';
import 'pages/login.dart';
// punto de entrada de app
void main(){
  runApp(const MasterGPro());
}
/// Widget raíz de la aplicación, donde va a llegar toda la informacion.
class MasterGPro extends StatelessWidget{
  const MasterGPro({super.key});
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(),
      // llamado widget login
      home: LoginPage(),
    );
  }
}
