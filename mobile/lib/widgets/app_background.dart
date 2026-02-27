import 'dart:ui';
import 'package:flutter/material.dart';
import '../theme/app_theme.dart';

/// Full-screen gradient background used by every screen.
class AppBackground extends StatelessWidget {
  final Widget child;
  const AppBackground({super.key, required this.child});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(gradient: AppGradients.background),
      child: child,
    );
  }
}

/// Decorative blurred orbs to create depth behind glass cards.
class GlowOrb extends StatelessWidget {
  final double size;
  final Color color;
  final Alignment alignment;

  const GlowOrb({
    super.key,
    this.size = 260,
    required this.color,
    required this.alignment,
  });

  @override
  Widget build(BuildContext context) {
    return Align(
      alignment: alignment,
      child: Container(
        width: size,
        height: size,
        decoration: BoxDecoration(
          shape: BoxShape.circle,
          color: color.withAlpha(60),
        ),
        child: BackdropFilter(
          filter: ImageFilter.blur(sigmaX: 80, sigmaY: 80),
          child: const SizedBox.expand(),
        ),
      ),
    );
  }
}
