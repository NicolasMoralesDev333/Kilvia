import 'package:flutter/material.dart';

import 'auth/auth_screens.dart';
import 'fleet/fleet_screens.dart';
import 'home/home_screens.dart';
import 'opportunities/opportunity_screens.dart';
import 'trips/trip_screens.dart';

export 'auth/auth_screens.dart';
export 'fleet/fleet_screens.dart';
export 'home/home_screens.dart';
export 'opportunities/opportunity_screens.dart';
export 'trips/trip_screens.dart';
export 'widgets/screen_components.dart';

Map<String, WidgetBuilder> get kilviaRoutes => {
  '/splash': (_) => const SplashScreen(),
  '/login': (_) => const LoginScreen(),
  '/register': (_) => const RegisterScreen(),
  '/home': (_) => const HomeScreen(),
  '/profile': (_) => const ProfileScreen(),
  '/company': (_) => const CompanySettingsScreen(),
  '/trips': (_) => const TripsListScreen(),
  '/trip-detail': (_) => const TripDetailScreen(),
  '/trip-create': (_) => const TripCreateScreen(),
  '/loads': (_) => const LoadsListScreen(),
  '/load-detail': (_) => const LoadDetailScreen(),
  '/load-create': (_) => const LoadCreateScreen(),
  '/opportunities': (_) => const OpportunitiesScreen(),
  '/match-detail': (_) => const MatchDetailScreen(),
  '/vehicles': (_) => const VehiclesListScreen(),
  '/drivers': (_) => const DriversListScreen(),
};
