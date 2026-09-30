import 'package:flutter/material.dart';

import 'core/app_theme.dart';
import 'screens/screens.dart';

void main() => runApp(const KilviaApp());

class KilviaApp extends StatelessWidget {
  const KilviaApp({super.key});

  @override
  Widget build(BuildContext context) => MaterialApp(
    title: 'Kilvia',
    debugShowCheckedModeBanner: false,
    theme: AppTheme.theme,
    initialRoute: '/splash',
    routes: kilviaRoutes,
  );
}
