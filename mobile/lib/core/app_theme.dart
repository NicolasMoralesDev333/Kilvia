import 'package:flutter/material.dart';

abstract final class AppTheme {
  static const navy = Color(0xFF0B1D3A);
  static const blue = Color(0xFF1E6BFF);
  static const green = Color(0xFF22C06E);
  static const light = Color(0xFFF4F6F9);
  static const slate = Color(0xFF667085);
  static const border = Color(0xFFE2E8F0);
  static const white = Colors.white;

  static ThemeData get theme => ThemeData(
    fontFamily: 'Manrope',
    useMaterial3: true,
    scaffoldBackgroundColor: light,
    primaryColor: navy,
    colorScheme: ColorScheme.fromSeed(
      seedColor: blue,
      primary: navy,
      secondary: blue,
      surface: white,
      onSurface: navy,
    ),
    appBarTheme: const AppBarTheme(
      backgroundColor: navy,
      foregroundColor: white,
      elevation: 0,
      centerTitle: false,
      titleTextStyle: TextStyle(
        color: white,
        fontSize: 20,
        fontWeight: FontWeight.w800,
        fontFamily: 'Manrope',
      ),
    ),
    cardTheme: CardThemeData(
      color: white,
      elevation: 0,
      margin: EdgeInsets.zero,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(14),
        side: const BorderSide(color: border, width: 1),
      ),
      shadowColor: const Color(0x1A0B1D3A),
    ),
    inputDecorationTheme: InputDecorationTheme(
      filled: true,
      fillColor: white,
      contentPadding: const EdgeInsets.symmetric(horizontal: 14, vertical: 14),
      enabledBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(12),
        borderSide: const BorderSide(color: border),
      ),
      focusedBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(12),
        borderSide: const BorderSide(color: blue, width: 1.5),
      ),
      errorBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(12),
        borderSide: const BorderSide(color: Color(0xFFB42318)),
      ),
      labelStyle: const TextStyle(color: slate),
    ),
    elevatedButtonTheme: ElevatedButtonThemeData(
      style: ElevatedButton.styleFrom(
        backgroundColor: blue,
        foregroundColor: white,
        minimumSize: const Size.fromHeight(48),
        elevation: 0,
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
        textStyle: const TextStyle(fontWeight: FontWeight.w800),
      ),
    ),
    navigationBarTheme: NavigationBarThemeData(
      backgroundColor: white,
      indicatorColor: blue.withValues(alpha: 0.12),
      labelTextStyle: WidgetStateProperty.all(
        const TextStyle(fontSize: 12, fontWeight: FontWeight.w700),
      ),
      iconTheme: WidgetStateProperty.all(const IconThemeData(color: slate)),
    ),
    chipTheme: ChipThemeData(
      backgroundColor: const Color(0xFFF3F6FA),
      selectedColor: blue.withValues(alpha: 0.12),
      checkmarkColor: blue,
      labelStyle: const TextStyle(color: navy, fontWeight: FontWeight.w700),
      secondaryLabelStyle: const TextStyle(
        color: blue,
        fontWeight: FontWeight.w800,
      ),
      side: const BorderSide(color: border),
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(999)),
    ),
    textTheme: const TextTheme(
      headlineLarge: TextStyle(color: navy, fontWeight: FontWeight.w800),
      titleLarge: TextStyle(color: navy, fontWeight: FontWeight.w800),
      titleMedium: TextStyle(color: navy, fontWeight: FontWeight.w700),
      bodyLarge: TextStyle(color: navy),
      bodyMedium: TextStyle(color: navy),
      labelLarge: TextStyle(color: navy, fontWeight: FontWeight.w700),
    ),
  );
}
