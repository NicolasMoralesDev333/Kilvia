import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';

import '../widgets/screen_components.dart';

class OpportunitiesScreen extends StatefulWidget {
  const OpportunitiesScreen({super.key});

  @override
  State<OpportunitiesScreen> createState() => _OpportunitiesScreenState();
}

class _OpportunitiesScreenState extends State<OpportunitiesScreen> {
  bool _offerSent = false;

  @override
  Widget build(BuildContext context) => KilviaFrame(
    tab: 3,
    title: 'Oportunidades',
    child: ListView(
      padding: const EdgeInsets.fromLTRB(16, 20, 16, 20),
      children: [
        const Text(
          'CARGA RECOMENDADA · DATOS DEMO',
          style: TextStyle(
            color: Color(0xFF8FB8FF),
            fontSize: 18,
            letterSpacing: 1.5,
            fontWeight: FontWeight.w800,
          ),
        ),
        const SizedBox(height: 14),
        const Text(
          'Buenos Aires → San Luis',
          style: TextStyle(
            color: Colors.white,
            fontSize: 52,
            fontWeight: FontWeight.w800,
            height: 0.95,
          ),
        ),
        const SizedBox(height: 16),
        const Text(
          'Compatible con el viaje Buenos Aires → Mendoza',
          style: TextStyle(color: Colors.white70, fontSize: 18),
        ),
        const SizedBox(height: 18),
        GestureDetector(
          onTap: () {
            setState(() => _offerSent = true);
            ScaffoldMessenger.of(context).showSnackBar(
              const SnackBar(
                content: Text('Oferta enviada correctamente'),
                behavior: SnackBarBehavior.floating,
                duration: Duration(seconds: 2),
              ),
            );
          },
          child: Container(
            height: 64,
            decoration: BoxDecoration(
              color: _offerSent
                  ? const Color(0xFF22C06E)
                  : const Color(0xFF1E6BFF),
              borderRadius: const BorderRadius.all(Radius.circular(14)),
            ),
            child: Center(
              child: Text(
                _offerSent ? 'Oferta enviada' : 'Enviar oferta',
                style: const TextStyle(
                  color: Colors.white,
                  fontSize: 26,
                  fontWeight: FontWeight.w700,
                ),
              ),
            ),
          ),
        ),
        const SizedBox(height: 18),
        Container(
          padding: const EdgeInsets.all(12),
          decoration: BoxDecoration(
            color: const Color(0xFF102E4A),
            borderRadius: BorderRadius.circular(18),
            border: Border.all(color: Colors.white.withValues(alpha: 0.12)),
          ),
          child: const Column(
            children: [
              _RouteMapCard(),
              SizedBox(height: 16),
              Row(
                children: [
                  Expanded(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          'Destino carga',
                          style: TextStyle(color: Colors.white70, fontSize: 16),
                        ),
                        SizedBox(height: 8),
                        Text(
                          'San Luis',
                          style: TextStyle(
                            color: Colors.white,
                            fontSize: 22,
                            fontWeight: FontWeight.w800,
                          ),
                        ),
                      ],
                    ),
                  ),
                  Expanded(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          'Origen compartido',
                          style: TextStyle(color: Colors.white70, fontSize: 16),
                        ),
                        SizedBox(height: 8),
                        Text(
                          'Buenos Aires',
                          style: TextStyle(
                            color: Colors.white,
                            fontSize: 22,
                            fontWeight: FontWeight.w800,
                          ),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ],
    ),
  );
}

class MatchDetailScreen extends StatefulWidget {
  const MatchDetailScreen({super.key});

  @override
  State<MatchDetailScreen> createState() => _MatchDetailScreenState();
}

class _MatchDetailScreenState extends State<MatchDetailScreen> {
  bool _offerSent = false;

  @override
  Widget build(BuildContext context) => KilviaFrame(
    tab: 3,
    title: 'Oportunidades / Detalle del match',
    child: ListView(
      padding: const EdgeInsets.fromLTRB(16, 20, 16, 20),
      children: [
        const Text(
          'capacidad libre',
          style: TextStyle(
            color: Colors.white,
            fontSize: 40,
            fontWeight: FontWeight.w800,
          ),
        ),
        const SizedBox(height: 10),
        const Text(
          'La visualización compara el peso de la carga con el espacio disponible antes del match.',
          style: TextStyle(color: Colors.white70, fontSize: 20),
        ),
        const SizedBox(height: 24),
        const Row(
          children: [
            Expanded(
              child: Text(
                'Carga',
                style: TextStyle(color: Colors.white70, fontSize: 18),
              ),
            ),
            Expanded(
              child: Align(
                alignment: Alignment.centerRight,
                child: Text(
                  'Capacidad libre',
                  style: TextStyle(color: Colors.white70, fontSize: 18),
                ),
              ),
            ),
          ],
        ),
        const SizedBox(height: 6),
        const Row(
          children: [
            Expanded(
              child: Text(
                '8,2 t',
                style: TextStyle(
                  color: Colors.white,
                  fontSize: 34,
                  fontWeight: FontWeight.w800,
                ),
              ),
            ),
            Expanded(
              child: Align(
                alignment: Alignment.centerRight,
                child: Text(
                  '14,1 t',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 34,
                    fontWeight: FontWeight.w800,
                  ),
                ),
              ),
            ),
          ],
        ),
        const SizedBox(height: 18),
        Container(
          height: 80,
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(14),
            border: Border.all(color: Colors.white.withValues(alpha: 0.2)),
            color: const Color(0xFF1B3950),
          ),
          child: Row(
            children: [
              Expanded(
                flex: 58,
                child: Container(
                  decoration: const BoxDecoration(
                    color: Color(0xFF22C06E),
                    borderRadius: BorderRadius.only(
                      topLeft: Radius.circular(14),
                      bottomLeft: Radius.circular(14),
                    ),
                  ),
                  alignment: Alignment.center,
                  child: const Text(
                    '58% EN USO',
                    style: TextStyle(
                      color: Colors.white,
                      fontSize: 20,
                      fontWeight: FontWeight.w800,
                    ),
                  ),
                ),
              ),
              Expanded(
                flex: 42,
                child: Container(
                  decoration: BoxDecoration(
                    color: Colors.white.withValues(alpha: 0.12),
                    borderRadius: const BorderRadius.only(
                      topRight: Radius.circular(14),
                      bottomRight: Radius.circular(14),
                    ),
                  ),
                  alignment: Alignment.center,
                  child: const Text(
                    '5,9 t libres',
                    style: TextStyle(
                      color: Colors.white,
                      fontSize: 20,
                      fontWeight: FontWeight.w800,
                    ),
                  ),
                ),
              ),
            ],
          ),
        ),
        const SizedBox(height: 26),
        const Text(
          'VALIDACIÓN OPERATIVA',
          style: TextStyle(
            color: Color(0xFF8FB8FF),
            letterSpacing: 1.5,
            fontSize: 18,
            fontWeight: FontWeight.w800,
          ),
        ),
        const SizedBox(height: 12),
        const Text(
          '¿Por qué es compatible?',
          style: TextStyle(
            color: Colors.white,
            fontSize: 42,
            fontWeight: FontWeight.w800,
            height: 1,
          ),
        ),
        const SizedBox(height: 12),
        const Text(
          'Señales principales para evaluar este match.',
          style: TextStyle(color: Colors.white70, fontSize: 20),
        ),
        const SizedBox(height: 18),
        ...[
          ('Ruta', 'Muy compatible', true),
          ('Fecha', 'Compatible', true),
          ('Vehículo', 'Compatible', true),
          ('Capacidad', 'Compatible', true),
        ].map(
          (entry) => Container(
            padding: const EdgeInsets.symmetric(vertical: 18),
            decoration: const BoxDecoration(
              border: Border(bottom: BorderSide(color: Color(0xFF22466A))),
            ),
            child: Row(
              children: [
                Checkbox(
                  value: true,
                  activeColor: const Color(0xFF22C06E),
                  onChanged: (_) {},
                ),
                Expanded(
                  child: Text(
                    entry.$1,
                    style: const TextStyle(
                      color: Colors.white,
                      fontSize: 24,
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                ),
                Text(
                  entry.$2,
                  style: TextStyle(
                    color: entry.$3 ? const Color(0xFF22C06E) : Colors.white70,
                    fontSize: 22,
                    fontWeight: FontWeight.w700,
                  ),
                ),
              ],
            ),
          ),
        ),
        const SizedBox(height: 20),
        GestureDetector(
          onTap: () {
            setState(() => _offerSent = true);
            ScaffoldMessenger.of(context).showSnackBar(
              const SnackBar(
                content: Text('Oferta enviada correctamente'),
                behavior: SnackBarBehavior.floating,
                duration: Duration(seconds: 2),
              ),
            );
          },
          child: Container(
            height: 64,
            decoration: BoxDecoration(
              color: _offerSent
                  ? const Color(0xFF22C06E)
                  : const Color(0xFF1E6BFF),
              borderRadius: BorderRadius.circular(14),
            ),
            child: Center(
              child: Text(
                _offerSent ? 'Oferta enviada' : 'Enviar oferta',
                style: const TextStyle(
                  color: Colors.white,
                  fontSize: 26,
                  fontWeight: FontWeight.w700,
                ),
              ),
            ),
          ),
        ),
      ],
    ),
  );
}

class _RouteMapCard extends StatelessWidget {
  const _RouteMapCard();

  @override
  Widget build(BuildContext context) {
    const start = LatLng(-34.6037, -58.3816);
    const end = LatLng(-32.8908, -68.8446);
    const mid = LatLng(-33.8, -63.3);

    return ClipRRect(
      borderRadius: BorderRadius.circular(18),
      child: SizedBox(
        height: 260,
        child: FlutterMap(
          options: const MapOptions(
            initialCenter: mid,
            initialZoom: 5.1,
            interactionOptions: InteractionOptions(flags: InteractiveFlag.all),
          ),
          children: [
            TileLayer(
              urlTemplate: 'https://tile.openstreetmap.org/{z}/{x}/{y}.png',
              userAgentPackageName: 'com.example.mobile',
            ),
            PolylineLayer(
              polylines: [
                Polyline(
                  points: [start, end],
                  strokeWidth: 4,
                  color: Color(0xFF1E6BFF),
                ),
              ],
            ),
            MarkerLayer(
              markers: [
                Marker(
                  point: start,
                  width: 26,
                  height: 26,
                  child: const DecoratedBox(
                    decoration: BoxDecoration(
                      color: Color(0xFF1E6BFF),
                      shape: BoxShape.circle,
                      border: Border.fromBorderSide(
                        BorderSide(color: Colors.white, width: 3),
                      ),
                    ),
                  ),
                ),
                Marker(
                  point: end,
                  width: 26,
                  height: 26,
                  child: const DecoratedBox(
                    decoration: BoxDecoration(
                      color: Color(0xFF22C06E),
                      shape: BoxShape.circle,
                      border: Border.fromBorderSide(
                        BorderSide(color: Colors.white, width: 3),
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
