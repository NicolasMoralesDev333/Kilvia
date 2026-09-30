import 'package:flutter/material.dart';

import '../../core/mock_data.dart';
import '../widgets/screen_components.dart';

class TripsListScreen extends StatefulWidget {
  const TripsListScreen({super.key});

  @override
  State<TripsListScreen> createState() => _TripsListScreenState();
}

class _TripsListScreenState extends State<TripsListScreen> {
  String _filter = 'Todos';

  @override
  Widget build(BuildContext context) => KilviaFrame(
    tab: 1,
    title: 'Viajes',
    child: Column(
      children: [
        SingleChildScrollView(
          scrollDirection: Axis.horizontal,
          padding: const EdgeInsets.all(12),
          child: Row(
            children: [
              for (final label in [
                'Todos',
                'Activo',
                'Programado',
                'Finalizado',
              ])
                Padding(
                  padding: const EdgeInsets.only(right: 8),
                  child: ChoiceChip(
                    label: Text(label),
                    selected: _filter == label,
                    onSelected: (_) => setState(() => _filter = label),
                  ),
                ),
            ],
          ),
        ),
        Expanded(
          child: ListView(
            padding: const EdgeInsets.symmetric(horizontal: 12),
            children: [
              for (final trip in MockData.trips.where(
                (trip) => _filter == 'Todos' || trip['status'] == _filter,
              ))
                RecordCard(trip, '/trip-detail'),
            ],
          ),
        ),
      ],
    ),
  );
}

class TripDetailScreen extends StatelessWidget {
  const TripDetailScreen({super.key});

  @override
  Widget build(BuildContext context) => const DetailsScreen(
    title: 'Detalle del viaje',
    values: [
      ('Estado', 'Activo'),
      ('Origen', 'Buenos Aires, Buenos Aires'),
      ('Destino', 'Córdoba, Córdoba'),
      ('Chofer', 'Martín Ríos'),
      ('Vehículo', 'AB 123 CD'),
      ('Capacidad disponible', '6.400 kg · 32 m³'),
    ],
  );
}

class TripCreateScreen extends StatelessWidget {
  const TripCreateScreen({super.key});

  @override
  Widget build(BuildContext context) => const KilviaFormScreen(
    title: 'Nuevo viaje',
    fields: [
      'Origen',
      'Destino',
      'Fecha y hora',
      'Chofer',
      'Vehículo',
      'Capacidad total (kg)',
    ],
  );
}

class LoadsListScreen extends StatelessWidget {
  const LoadsListScreen({super.key});

  @override
  Widget build(BuildContext context) => KilviaFrame(
    tab: 2,
    title: 'Cargas',
    child: ListView(
      padding: const EdgeInsets.all(12),
      children: [
        for (final load in MockData.loads) RecordCard(load, '/load-detail'),
      ],
    ),
  );
}

class LoadDetailScreen extends StatelessWidget {
  const LoadDetailScreen({super.key});

  @override
  Widget build(BuildContext context) => const DetailsScreen(
    title: 'Detalle de carga',
    values: [
      ('Carga', 'Alimentos en pallets'),
      ('Peso', '8.400 kg'),
      ('Volumen', '32 m³'),
      ('Tipo', 'Refrigerada'),
      ('Restricciones', 'Mantener entre 2 y 8 °C'),
      ('Ruta', 'Buenos Aires → Córdoba'),
    ],
  );
}

class LoadCreateScreen extends StatelessWidget {
  const LoadCreateScreen({super.key});

  @override
  Widget build(BuildContext context) => const KilviaFormScreen(
    title: 'Publicar carga',
    fields: [
      'Descripción',
      'Origen',
      'Destino',
      'Peso (kg)',
      'Volumen (m³)',
      'Tipo de carga',
      'Restricciones',
    ],
  );
}
