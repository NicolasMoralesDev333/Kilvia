import 'package:flutter/material.dart';

import '../widgets/screen_components.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) => KilviaFrame(
    tab: 0,
    title: '192.168.1.33:4200/home',
    child: ListView(
      padding: const EdgeInsets.fromLTRB(16, 14, 16, 20),
      children: [
        const SizedBox(height: 8),
        Container(
          padding: const EdgeInsets.all(18),
          decoration: const BoxDecoration(
            color: Color(0xFF102E4A),
            borderRadius: BorderRadius.all(Radius.circular(18)),
          ),
          child: Row(
            children: [
              SizedBox(
                width: 120,
                height: 120,
                child: Stack(
                  alignment: Alignment.center,
                  children: [
                    SizedBox(
                      width: 110,
                      height: 110,
                      child: CircularProgressIndicator(
                        value: 0.94,
                        strokeWidth: 12,
                        backgroundColor: Colors.white.withValues(alpha: 0.15),
                        valueColor: const AlwaysStoppedAnimation<Color>(
                          Color(0xFF22C06E),
                        ),
                      ),
                    ),
                    const Text(
                      '94%',
                      style: TextStyle(
                        color: Colors.white,
                        fontSize: 34,
                        fontWeight: FontWeight.w800,
                      ),
                    ),
                  ],
                ),
              ),
              const SizedBox(width: 20),
              const Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'MATCH',
                      style: TextStyle(
                        color: Color(0xFF6ED7A3),
                        letterSpacing: 1.2,
                        fontSize: 18,
                        fontWeight: FontWeight.w800,
                      ),
                    ),
                    SizedBox(height: 6),
                    Text(
                      'Alta compatibilidad',
                      style: TextStyle(
                        color: Colors.white,
                        fontSize: 22,
                        fontWeight: FontWeight.w700,
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
        const SizedBox(height: 18),
        Container(
          decoration: const BoxDecoration(
            color: Color(0xFF102E4A),
            borderRadius: BorderRadius.all(Radius.circular(16)),
          ),
          child: const Row(
            children: [
              Expanded(
                child: Padding(
                  padding: EdgeInsets.all(18),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text('Desvío', style: TextStyle(color: Colors.white70)),
                      SizedBox(height: 8),
                      Text(
                        '+38 km',
                        style: TextStyle(
                          color: Colors.white,
                          fontSize: 28,
                          fontWeight: FontWeight.w800,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              Expanded(
                child: Padding(
                  padding: EdgeInsets.all(18),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        'Tiempo extra',
                        style: TextStyle(color: Colors.white70),
                      ),
                      SizedBox(height: 8),
                      Text(
                        '+35 min',
                        style: TextStyle(
                          color: Colors.white,
                          fontSize: 28,
                          fontWeight: FontWeight.w800,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              Expanded(
                child: Padding(
                  padding: EdgeInsets.all(18),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        'Capacidad',
                        style: TextStyle(color: Colors.white70),
                      ),
                      SizedBox(height: 8),
                      Text(
                        '58% utilizada',
                        style: TextStyle(
                          color: Colors.white,
                          fontSize: 22,
                          fontWeight: FontWeight.w800,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
        const SizedBox(height: 18),
        GestureDetector(
          onTap: () => Navigator.pushNamed(context, '/match-detail'),
          child: Container(
            height: 70,
            decoration: const BoxDecoration(
              color: Color(0xFF1E6BFF),
              borderRadius: BorderRadius.all(Radius.circular(16)),
            ),
            child: const Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Padding(
                  padding: EdgeInsets.only(left: 20),
                  child: Text(
                    'Evaluar match',
                    style: TextStyle(
                      color: Colors.white,
                      fontSize: 26,
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                ),
                Padding(
                  padding: EdgeInsets.only(right: 22),
                  child: Icon(
                    Icons.chevron_right,
                    color: Colors.white,
                    size: 34,
                  ),
                ),
              ],
            ),
          ),
        ),
        const SizedBox(height: 24),
        const Text(
          'CORREDORES EN AGENDA',
          style: TextStyle(
            color: Color(0xFF8FB8FF),
            letterSpacing: 1.4,
            fontSize: 18,
            fontWeight: FontWeight.w800,
          ),
        ),
        const SizedBox(height: 14),
        const Text(
          'Viajes programados',
          style: TextStyle(
            color: Colors.white,
            fontSize: 36,
            fontWeight: FontWeight.w800,
          ),
        ),
        const SizedBox(height: 8),
        const Text(
          'Capacidad y matching',
          style: TextStyle(color: Colors.white70, fontSize: 18),
        ),
        const SizedBox(height: 18),
        Container(
          padding: const EdgeInsets.all(18),
          decoration: BoxDecoration(
            color: const Color(0xFF102E4A),
            borderRadius: BorderRadius.circular(18),
            border: Border.all(color: Colors.white.withValues(alpha: 0.12)),
          ),
          child: Column(
            children: [
              Row(
                children: [
                  Container(
                    width: 12,
                    height: 12,
                    decoration: const BoxDecoration(
                      color: Color(0xFF5DA9FF),
                      shape: BoxShape.circle,
                    ),
                  ),
                  const SizedBox(width: 12),
                  const Expanded(
                    child: Text(
                      'VJ-024 · 12 sep',
                      style: TextStyle(color: Colors.white70, fontSize: 18),
                    ),
                  ),
                  Container(
                    padding: const EdgeInsets.symmetric(
                      horizontal: 12,
                      vertical: 8,
                    ),
                    decoration: BoxDecoration(
                      color: const Color(0xFFB9F1D2),
                      borderRadius: BorderRadius.circular(12),
                    ),
                    child: const Text(
                      '3 oportunidades',
                      style: TextStyle(
                        color: Color(0xFF0B1D3A),
                        fontWeight: FontWeight.w800,
                      ),
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 18),
              const Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'Buenos Aires → Mendoza',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 26,
                    fontWeight: FontWeight.w700,
                  ),
                ),
              ),
            ],
          ),
        ),
      ],
    ),
  );
}

class ProfileScreen extends StatelessWidget {
  const ProfileScreen({super.key});

  @override
  Widget build(BuildContext context) => KilviaFrame(
    tab: 4,
    title: 'Mi perfil',
    child: ListView(
      padding: const EdgeInsets.all(16),
      children: [
        const Card(
          child: ListTile(
            leading: CircleAvatar(child: Text('MR')),
            title: Text('Martín Ríos'),
            subtitle: Text('Administrador · Transportes del Sur'),
          ),
        ),
        const Card(
          child: Padding(
            padding: EdgeInsets.all(16),
            child: Column(
              children: [
                InfoLine('Correo', 'martin@transportes.com'),
                InfoLine('Teléfono', '+54 9 11 5555-0123'),
                InfoLine('Rol', 'Administrador'),
              ],
            ),
          ),
        ),
        const RecordCard({
          'title': 'Configuración de empresa',
          'meta': 'CUIT · datos operativos',
        }, '/company'),
        OutlinedButton.icon(
          onPressed: () => Navigator.pushNamedAndRemoveUntil(
            context,
            '/login',
            (_) => false,
          ),
          icon: const Icon(Icons.logout),
          label: const Text('Cambiar sesión'),
        ),
      ],
    ),
  );
}

class CompanySettingsScreen extends StatefulWidget {
  const CompanySettingsScreen({super.key});

  @override
  State<CompanySettingsScreen> createState() => _CompanySettingsScreenState();
}

class _CompanySettingsScreenState extends State<CompanySettingsScreen> {
  bool _alerts = true;

  @override
  Widget build(BuildContext context) => KilviaFrame(
    title: 'Empresa',
    child: ListView(
      padding: const EdgeInsets.all(16),
      children: [
        for (final field in const [
          ('Razón social', 'Transportes del Sur S.A.'),
          ('CUIT', '30-71234567-8'),
          ('Base operativa', 'Buenos Aires, Argentina'),
        ]) ...[
          TextFormField(
            initialValue: field.$2,
            decoration: InputDecoration(labelText: field.$1),
          ),
          const SizedBox(height: 12),
        ],
        SwitchListTile(
          contentPadding: EdgeInsets.zero,
          value: _alerts,
          title: const Text('Alertas de oportunidades'),
          onChanged: (value) => setState(() => _alerts = value),
        ),
        ElevatedButton(
          onPressed: () => ScaffoldMessenger.of(context)
              .showSnackBar(const SnackBar(content: Text('Cambios guardados'))),
          child: const Text('Guardar cambios'),
        ),
      ],
    ),
  );
}
