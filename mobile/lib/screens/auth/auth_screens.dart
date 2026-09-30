import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

import '../../core/app_theme.dart';
import '../widgets/screen_components.dart';

class SplashScreen extends StatelessWidget {
  const SplashScreen({super.key});

  @override
  Widget build(BuildContext context) => Scaffold(
    backgroundColor: AppTheme.navy,
    body: SafeArea(
      child: LayoutBuilder(
        builder: (context, constraints) => SingleChildScrollView(
          child: ConstrainedBox(
            constraints: BoxConstraints(minHeight: constraints.maxHeight),
            child: Center(
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  SvgPicture.asset(
                    'assets/images/kilvia_route.svg',
                    width: 72,
                    height: 72,
                    semanticsLabel: 'Marca Kilvia',
                  ),
                  const SizedBox(height: 12),
                  const Text(
                    'kilvia',
                    style: TextStyle(
                      color: Colors.white,
                      fontSize: 36,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                  const Text(
                    'Logística conectada',
                    style: TextStyle(color: Color(0xFFCFD8E6)),
                  ),
                  const SizedBox(height: 36),
                  SizedBox(
                    width: 220,
                    child: ElevatedButton(
                      onPressed: () =>
                          Navigator.pushReplacementNamed(context, '/login'),
                      child: const Text('Continuar'),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    ),
  );
}

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final _key = GlobalKey<FormState>();
  String _role = 'Administrador';

  @override
  Widget build(BuildContext context) => Scaffold(
    appBar: AppBar(title: const Text('Iniciar sesión')),
    body: Center(
      child: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Form(
          key: _key,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              const Text(
                'Buen día',
                style: TextStyle(
                  fontSize: 28,
                  fontWeight: FontWeight.bold,
                  color: AppTheme.navy,
                ),
              ),
              const Text(
                'Accedé a tu operación logística',
                style: TextStyle(color: AppTheme.slate),
              ),
              const SizedBox(height: 24),
              TextFormField(
                decoration: const InputDecoration(
                  labelText: 'Correo electrónico',
                  prefixIcon: Icon(Icons.mail_outline),
                ),
                keyboardType: TextInputType.emailAddress,
                validator: (value) => value != null && value.contains('@')
                    ? null
                    : 'Ingresá un correo válido',
              ),
              const SizedBox(height: 12),
              TextFormField(
                decoration: const InputDecoration(
                  labelText: 'Contraseña',
                  prefixIcon: Icon(Icons.lock_outline),
                ),
                obscureText: true,
                validator: (value) => value != null && value.length >= 4
                    ? null
                    : 'Ingresá tu contraseña',
              ),
              const SizedBox(height: 12),
              DropdownButtonFormField<String>(
                initialValue: _role,
                decoration: const InputDecoration(labelText: 'Rol de prueba'),
                items: const [
                  DropdownMenuItem(
                    value: 'Administrador',
                    child: Text('Administrador'),
                  ),
                  DropdownMenuItem(value: 'Operador', child: Text('Operador')),
                  DropdownMenuItem(
                    value: 'Transportista',
                    child: Text('Transportista'),
                  ),
                ],
                onChanged: (value) => setState(() => _role = value ?? _role),
              ),
              const SizedBox(height: 20),
              ElevatedButton(
                onPressed: () {
                  if (_key.currentState!.validate()) {
                    Navigator.pushReplacementNamed(context, '/home');
                  }
                },
                child: const Text('Ingresar'),
              ),
              TextButton(
                onPressed: () => Navigator.pushNamed(context, '/register'),
                child: const Text('Crear cuenta de empresa'),
              ),
            ],
          ),
        ),
      ),
    ),
  );
}

class RegisterScreen extends StatelessWidget {
  const RegisterScreen({super.key});

  @override
  Widget build(BuildContext context) => const KilviaFormScreen(
    title: 'Alta de empresa',
    fields: [
      'Empresa',
      'CUIT',
      'Nombre del administrador',
      'Correo electrónico',
      'Contraseña',
    ],
  );
}
