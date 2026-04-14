import 'package:dizect/src/app/theme/app_theme.dart';
import 'package:dizect/src/features/home/presentation/pages/home_page.dart';
import 'package:flutter/material.dart';

class DistributedStorageApp extends StatelessWidget {
  const DistributedStorageApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Distributed Storage',
      theme: AppTheme.light(),
      home: const HomePage(),
    );
  }
}
