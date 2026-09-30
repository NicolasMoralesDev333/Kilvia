import 'package:flutter/material.dart';

import '../../core/app_theme.dart';

class KilviaFrame extends StatelessWidget {
  const KilviaFrame({
    required this.title,
    required this.child,
    this.tab,
    super.key,
  });

  final String title;
  final Widget child;
  final int? tab;

  static const _tabs = [
    ('/home', 'Inicio', Icons.home_outlined),
    ('/trips', 'Viajes', Icons.local_shipping_outlined),
    ('/loads', 'Cargas', Icons.inventory_2_outlined),
    ('/opportunities', 'Oportunidades', Icons.tune_rounded),
    ('/profile', 'Más', Icons.more_horiz_rounded),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFF0B1D3A),
      body: SafeArea(
        child: Column(
          children: [
            Padding(
              padding: const EdgeInsets.fromLTRB(16, 8, 16, 8),
              child: Row(
                children: [
                  Container(
                    width: 42,
                    height: 42,
                    decoration: BoxDecoration(
                      color: const Color(0xFF1E6BFF),
                      borderRadius: BorderRadius.circular(12),
                    ),
                    child: const Icon(
                      Icons.home_rounded,
                      color: Colors.white,
                      size: 22,
                    ),
                  ),
                  const SizedBox(width: 12),
                  Expanded(
                    child: Container(
                      height: 42,
                      padding: const EdgeInsets.symmetric(horizontal: 12),
                      decoration: BoxDecoration(
                        color: const Color(0xFF2B2B2F),
                        borderRadius: BorderRadius.circular(18),
                      ),
                      child: Row(
                        children: [
                          const Icon(
                            Icons.warning_amber_rounded,
                            color: Colors.white,
                            size: 18,
                          ),
                          const SizedBox(width: 8),
                          Expanded(
                            child: FittedBox(
                              fit: BoxFit.scaleDown,
                              alignment: Alignment.centerLeft,
                              child: Text(
                                title,
                                maxLines: 1,
                                style: const TextStyle(
                                  color: Colors.white,
                                  fontSize: 16,
                                  fontWeight: FontWeight.w700,
                                ),
                              ),
                            ),
                          ),
                          if (tab != null) ...[
                            const SizedBox(width: 8),
                            Container(
                              width: 24,
                              height: 24,
                              alignment: Alignment.center,
                              decoration: BoxDecoration(
                                color: Colors.white,
                                borderRadius: BorderRadius.circular(8),
                              ),
                              child: Text(
                                '${tab! + 1}',
                                style: const TextStyle(
                                  color: Color(0xFF0B1D3A),
                                  fontSize: 12,
                                  fontWeight: FontWeight.w800,
                                ),
                              ),
                            ),
                          ],
                        ],
                      ),
                    ),
                  ),
                ],
              ),
            ),
            Expanded(
              child: Container(color: const Color(0xFF0B1D3A), child: child),
            ),
            if (tab != null)
              Container(
                padding: const EdgeInsets.only(top: 10, bottom: 6),
                decoration: const BoxDecoration(
                  color: Color(0xFF0B1D3A),
                  border: Border(top: BorderSide(color: Color(0xFF22466A))),
                ),
                child: LayoutBuilder(
                  builder: (context, constraints) {
                    final itemWidth = constraints.maxWidth / _tabs.length;
                    return Row(
                      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                      children: [
                        for (var i = 0; i < _tabs.length; i++)
                          SizedBox(
                            width: itemWidth,
                            child: GestureDetector(
                              onTap: () => Navigator.pushReplacementNamed(
                                context,
                                _tabs[i].$1,
                              ),
                              child: Column(
                                mainAxisSize: MainAxisSize.min,
                                children: [
                                  Icon(
                                    _tabs[i].$3,
                                    color: i == tab
                                        ? Colors.white
                                        : Colors.white.withValues(alpha: 0.7),
                                    size: 24,
                                  ),
                                  const SizedBox(height: 4),
                                  FittedBox(
                                    fit: BoxFit.scaleDown,
                                    child: Text(
                                      _tabs[i].$2,
                                      style: TextStyle(
                                        color: i == tab
                                            ? Colors.white
                                            : Colors.white.withValues(
                                                alpha: 0.72,
                                              ),
                                        fontSize: 11,
                                        fontWeight: FontWeight.w600,
                                      ),
                                    ),
                                  ),
                                ],
                              ),
                            ),
                          ),
                      ],
                    );
                  },
                ),
              ),
          ],
        ),
      ),
    );
  }
}

class StatusBadge extends StatelessWidget {
  const StatusBadge(this.label, {this.green = false, super.key});

  final String label;
  final bool green;

  @override
  Widget build(BuildContext context) => Container(
    padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 5),
    decoration: BoxDecoration(
      color: green
          ? AppTheme.green.withValues(alpha: 0.12)
          : AppTheme.blue.withValues(alpha: 0.1),
      borderRadius: BorderRadius.circular(999),
    ),
    child: Text(
      label,
      style: TextStyle(
        color: green ? const Color(0xFF178A50) : AppTheme.blue,
        fontSize: 11,
        fontWeight: FontWeight.w800,
      ),
    ),
  );
}

class RecordCard extends StatelessWidget {
  const RecordCard(this.item, this.route, {super.key});

  final Map<String, String> item;
  final String route;

  @override
  Widget build(BuildContext context) => Container(
    margin: const EdgeInsets.only(bottom: 10),
    child: Card(
      child: ListTile(
        contentPadding: const EdgeInsets.symmetric(
          horizontal: 16,
          vertical: 10,
        ),
        leading: Container(
          width: 42,
          height: 42,
          decoration: BoxDecoration(
            color: const Color(0xFFEAF0FF),
            borderRadius: BorderRadius.circular(12),
          ),
          child: const Icon(
            Icons.local_shipping_outlined,
            color: AppTheme.blue,
          ),
        ),
        title: Row(
          children: [
            Expanded(child: Text(item['title']!)),
            if (item.containsKey('status'))
              Padding(
                padding: const EdgeInsets.only(left: 8),
                child: StatusBadge(
                  item['status']!,
                  green:
                      item['status'] == 'Activo' ||
                      item['status'] == 'Disponible' ||
                      item['status'] == 'Publicada',
                ),
              ),
          ],
        ),
        subtitle: Padding(
          padding: const EdgeInsets.only(top: 5),
          child: Text(
            item['meta']!,
            style: const TextStyle(color: AppTheme.slate),
          ),
        ),
        trailing: const Icon(Icons.chevron_right, color: AppTheme.slate),
        onTap: () => Navigator.pushNamed(context, route),
      ),
    ),
  );
}

class KilviaFormScreen extends StatefulWidget {
  const KilviaFormScreen({
    required this.title,
    required this.fields,
    super.key,
  });

  final String title;
  final List<String> fields;

  @override
  State<KilviaFormScreen> createState() => _KilviaFormScreenState();
}

class _KilviaFormScreenState extends State<KilviaFormScreen> {
  final _key = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) => KilviaFrame(
    title: widget.title,
    child: Form(
      key: _key,
      child: ListView(
        padding: const EdgeInsets.all(16),
        children: [
          for (final field in widget.fields) ...[
            TextFormField(
              decoration: InputDecoration(labelText: field),
              validator: (value) => value == null || value.trim().isEmpty
                  ? 'Campo obligatorio'
                  : null,
            ),
            const SizedBox(height: 14),
          ],
          ElevatedButton(
            onPressed: () {
              if (_key.currentState!.validate()) {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('Registro guardado')),
                );
                Navigator.pop(context);
              }
            },
            child: const Text('Guardar'),
          ),
        ],
      ),
    ),
  );
}

class DetailsScreen extends StatelessWidget {
  const DetailsScreen({required this.title, required this.values, super.key});

  final String title;
  final List<(String, String)> values;

  @override
  Widget build(BuildContext context) => KilviaFrame(
    title: title,
    child: ListView(
      padding: const EdgeInsets.all(16),
      children: [
        Card(
          child: Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              children: [for (final row in values) InfoLine(row.$1, row.$2)],
            ),
          ),
        ),
        const SizedBox(height: 12),
        ElevatedButton(
          onPressed: () => Navigator.pop(context),
          child: const Text('Volver'),
        ),
      ],
    ),
  );
}

class InfoLine extends StatelessWidget {
  const InfoLine(this.label, this.value, {super.key});

  final String label;
  final String value;

  @override
  Widget build(BuildContext context) => Padding(
    padding: const EdgeInsets.symmetric(vertical: 8),
    child: Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Text(label, style: const TextStyle(color: AppTheme.slate)),
        Flexible(child: Text(value, textAlign: TextAlign.end)),
      ],
    ),
  );
}
