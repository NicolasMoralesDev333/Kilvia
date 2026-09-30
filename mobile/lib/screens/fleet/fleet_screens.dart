import 'package:flutter/material.dart';

import '../../core/mock_data.dart';
import '../widgets/screen_components.dart';

class VehiclesListScreen extends StatelessWidget {
  const VehiclesListScreen({super.key});

  @override
  Widget build(BuildContext context) => KilviaFrame(
    title: 'Flota',
    child: ListView(
      padding: const EdgeInsets.all(12),
      children: [
        for (final vehicle in MockData.vehicles)
          RecordCard(vehicle, '/trip-detail'),
      ],
    ),
  );
}

class DriversListScreen extends StatelessWidget {
  const DriversListScreen({super.key});

  @override
  Widget build(BuildContext context) => KilviaFrame(
    title: 'Choferes',
    child: ListView(
      padding: const EdgeInsets.all(12),
      children: [
        for (final driver in MockData.drivers)
          RecordCard(driver, '/trip-detail'),
      ],
    ),
  );
}
