// This is a basic Flutter widget test.
//
// To perform an interaction with a widget in your test, use the WidgetTester
// utility in the flutter_test package. For example, you can send tap and scroll
// gestures. You can also use WidgetTester to find child widgets in the widget
// tree, read text, and verify that the values of widget properties are correct.

import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import 'package:mobile/core/app_theme.dart';
import 'package:mobile/screens/screens.dart';

void main() {
  Widget buildMobileApp(String initialRoute) => MediaQuery(
    data: const MediaQueryData(size: Size(390, 844), devicePixelRatio: 3),
    child: MaterialApp(
      title: 'Kilvia',
      debugShowCheckedModeBanner: false,
      theme: AppTheme.theme,
      initialRoute: initialRoute,
      routes: kilviaRoutes,
    ),
  );

  testWidgets('muestra el inicio de Kilvia', (WidgetTester tester) async {
    await tester.pumpWidget(buildMobileApp('/splash'));
    await tester.pumpAndSettle();
    expect(find.text('kilvia'), findsOneWidget);
    expect(find.text('Continuar'), findsOneWidget);
  });

  testWidgets('el splash tolera un viewport inicial mínimo', (
    WidgetTester tester,
  ) async {
    tester.view.physicalSize = const Size(1, 1);
    tester.view.devicePixelRatio = 1;
    addTearDown(tester.view.resetPhysicalSize);
    addTearDown(tester.view.resetDevicePixelRatio);

    await tester.pumpWidget(
      MaterialApp(
        debugShowCheckedModeBanner: false,
        initialRoute: '/splash',
        routes: kilviaRoutes,
      ),
    );
    await tester.pump();

    expect(tester.takeException(), isNull);
  });

  testWidgets('captura la home en vista mobile', (WidgetTester tester) async {
    await tester.pumpWidget(buildMobileApp('/home'));
    await tester.pumpAndSettle();

    await expectLater(
      find.byType(HomeScreen),
      matchesGoldenFile('goldens/home_mobile.png'),
    );
  });

  testWidgets('captura la pantalla de oportunidad en vista mobile', (
    WidgetTester tester,
  ) async {
    await tester.pumpWidget(buildMobileApp('/opportunities'));
    await tester.pumpAndSettle();

    await expectLater(
      find.byType(OpportunitiesScreen),
      matchesGoldenFile('goldens/opportunities_mobile.png'),
    );
  });

  testWidgets('captura el detalle del match en vista mobile', (
    WidgetTester tester,
  ) async {
    await tester.pumpWidget(buildMobileApp('/match-detail'));
    await tester.pumpAndSettle();

    await expectLater(
      find.byType(MatchDetailScreen),
      matchesGoldenFile('goldens/match_detail_mobile.png'),
    );
  });
}
