abstract final class MockData {
  static const trips = [
    {
      'title': 'Buenos Aires → Córdoba',
      'meta': 'Hoy, 14:30 · Martín Ríos · AB 123 CD',
      'status': 'Activo',
    },
    {
      'title': 'Rosario → Mendoza',
      'meta': 'Mañana, 08:00 · Lucía Gómez · AC 456 EF',
      'status': 'Programado',
    },
    {
      'title': 'Córdoba → Buenos Aires',
      'meta': '18 jun · Diego Paz · AD 789 GH',
      'status': 'Finalizado',
    },
  ];
  static const loads = [
    {
      'title': 'Alimentos en pallets',
      'meta': 'Buenos Aires → Córdoba · 8.400 kg',
      'status': 'Refrigerada',
    },
    {
      'title': 'Materiales de construcción',
      'meta': 'Rosario → Mendoza · 12.000 kg',
      'status': 'General',
    },
    {
      'title': 'Insumos médicos',
      'meta': 'La Plata → Santa Fe · 2.100 kg',
      'status': 'Sensible',
    },
  ];
  static const vehicles = [
    {
      'title': 'AB 123 CD',
      'meta': 'Semirremolque · 28.000 kg',
      'status': 'Disponible',
    },
    {
      'title': 'AC 456 EF',
      'meta': 'Camión térmico · 18.000 kg',
      'status': 'En viaje',
    },
    {
      'title': 'AD 789 GH',
      'meta': 'Camión rígido · 8.000 kg',
      'status': 'Disponible',
    },
  ];
  static const drivers = [
    {
      'title': 'Martín Ríos',
      'meta': 'Licencia E1 · +54 9 11 5555-0123',
      'status': 'Disponible',
    },
    {
      'title': 'Lucía Gómez',
      'meta': 'Licencia E1 · +54 9 341 555-0456',
      'status': 'En viaje',
    },
    {
      'title': 'Diego Paz',
      'meta': 'Licencia E2 · +54 9 351 555-0789',
      'status': 'Disponible',
    },
  ];
}
