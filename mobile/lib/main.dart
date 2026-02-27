import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'providers/auth_provider.dart';
import 'router.dart';
import 'theme/app_theme.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  runApp(
    ChangeNotifierProvider(
      create: (_) => AuthProvider(),
      child: const AarcsApp(),
    ),
  );
}

class AarcsApp extends StatelessWidget {
  const AarcsApp({super.key});

  @override
  Widget build(BuildContext context) {
    final auth = context.watch<AuthProvider>();
    final router = buildRouter(auth);

    return MaterialApp.router(
      title: 'AARCS-X',
      debugShowCheckedModeBanner: false,
      theme: buildAppTheme(),
      routerConfig: router,
      builder: (context, child) {
        // Wrap every screen in the gradient background
        return Container(
          decoration: const BoxDecoration(gradient: AppGradients.background),
          child: child ?? const SizedBox.shrink(),
        );
      },
    );
  }
}
