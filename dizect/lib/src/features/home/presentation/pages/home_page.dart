import 'package:dizect/src/core/constants/app_strings.dart';
import 'package:dizect/src/core/widgets/info_card.dart';
import 'package:flutter/material.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    final colorScheme = Theme.of(context).colorScheme;

    return Scaffold(
      appBar: AppBar(title: const Text(AppStrings.appName)),
      body: SafeArea(
        child: ListView(
          padding: const EdgeInsets.all(24),
          children: [
            Container(
              padding: const EdgeInsets.all(24),
              decoration: BoxDecoration(
                gradient: LinearGradient(
                  colors: [colorScheme.primary, colorScheme.primaryContainer],
                  begin: Alignment.topLeft,
                  end: Alignment.bottomRight,
                ),
                borderRadius: BorderRadius.circular(28),
              ),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    AppStrings.welcomeTitle,
                    style: Theme.of(
                      context,
                    ).textTheme.headlineLarge?.copyWith(color: Colors.white),
                  ),
                  const SizedBox(height: 12),
                  Text(
                    AppStrings.welcomeSubtitle,
                    style: Theme.of(context).textTheme.bodyMedium?.copyWith(
                      color: Colors.white.withValues(alpha: 0.92),
                    ),
                  ),
                ],
              ),
            ),
            const SizedBox(height: 24),
            const InfoCard(
              title: 'Node orchestration',
              subtitle:
                  'Organize distributed services, background jobs, and sync flows in one place.',
              icon: Icons.hub_rounded,
            ),
            const SizedBox(height: 16),
            const InfoCard(
              title: 'Storage visibility',
              subtitle:
                  'Keep room for observability screens, file tracking, and operational dashboards.',
              icon: Icons.storage_rounded,
            ),
            const SizedBox(height: 16),
            const InfoCard(
              title: 'Feature-first architecture',
              subtitle:
                  'Scale the project with isolated modules for data, domain, and presentation layers.',
              icon: Icons.grid_view_rounded,
            ),
          ],
        ),
      ),
    );
  }
}
