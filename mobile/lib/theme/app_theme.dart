import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

// ─── Colour palette ───────────────────────────────────────────────────────────
class AppColors {
  // Backgrounds — deep purple-ish dark gradient stops
  static const Color bg1 = Color(0xFF0D0B1E); // near-black purple
  static const Color bg2 = Color(0xFF1A1333); // dark violet

  // Primary — light purple family
  static const Color primary = Color(0xFFB48EFF);       // soft lavender-violet
  static const Color primaryLight = Color(0xFFD4B8FF);  // lighter lilac
  static const Color primaryDark = Color(0xFF7C5CBF);   // deeper violet

  // Glass card surface
  static const Color glassWhite = Color(0x1AFFFFFF);    // 10% white
  static const Color glassBorder = Color(0x33FFFFFF);   // 20% white border
  static const Color glassDark  = Color(0x26000000);    // subtle shadow tint

  // Text
  static const Color textPrimary   = Color(0xFFF0EBFF);
  static const Color textSecondary = Color(0xFFB0A8CC);
  static const Color textHint      = Color(0xFF7A7090);

  // Status
  static const Color success = Color(0xFF7EEBB5);
  static const Color error   = Color(0xFFFF7070);
  static const Color warning = Color(0xFFFFD080);
}

// ─── Gradients ────────────────────────────────────────────────────────────────
class AppGradients {
  static const LinearGradient background = LinearGradient(
    begin: Alignment.topLeft,
    end: Alignment.bottomRight,
    colors: [AppColors.bg1, AppColors.bg2, Color(0xFF251848)],
    stops: [0.0, 0.5, 1.0],
  );

  static const LinearGradient primaryButton = LinearGradient(
    colors: [AppColors.primary, AppColors.primaryDark],
    begin: Alignment.centerLeft,
    end: Alignment.centerRight,
  );

  static const LinearGradient card = LinearGradient(
    begin: Alignment.topLeft,
    end: Alignment.bottomRight,
    colors: [Color(0x33B48EFF), Color(0x0DFFFFFF)],
  );
}

// ─── ThemeData ────────────────────────────────────────────────────────────────
ThemeData buildAppTheme() {
  final base = ThemeData.dark();
  final textTheme = GoogleFonts.outfitTextTheme(base.textTheme).copyWith(
    displayLarge: GoogleFonts.outfit(
        color: AppColors.textPrimary, fontWeight: FontWeight.w700, fontSize: 34),
    headlineMedium: GoogleFonts.outfit(
        color: AppColors.textPrimary, fontWeight: FontWeight.w600, fontSize: 22),
    titleLarge: GoogleFonts.outfit(
        color: AppColors.textPrimary, fontWeight: FontWeight.w600, fontSize: 18),
    bodyLarge: GoogleFonts.outfit(color: AppColors.textPrimary, fontSize: 16),
    bodyMedium: GoogleFonts.outfit(color: AppColors.textSecondary, fontSize: 14),
    labelLarge: GoogleFonts.outfit(
        color: AppColors.textPrimary, fontWeight: FontWeight.w600, fontSize: 16),
  );

  return base.copyWith(
    scaffoldBackgroundColor: Colors.transparent,
    colorScheme: const ColorScheme.dark(
      primary: AppColors.primary,
      secondary: AppColors.primaryLight,
      surface: AppColors.glassWhite,
      error: AppColors.error,
    ),
    textTheme: textTheme,
    appBarTheme: AppBarTheme(
      backgroundColor: Colors.transparent,
      elevation: 0,
      titleTextStyle: GoogleFonts.outfit(
          color: AppColors.textPrimary, fontWeight: FontWeight.w600, fontSize: 18),
      iconTheme: const IconThemeData(color: AppColors.textPrimary),
    ),
    inputDecorationTheme: InputDecorationTheme(
      filled: true,
      fillColor: AppColors.glassWhite,
      hintStyle: GoogleFonts.outfit(color: AppColors.textHint),
      labelStyle: GoogleFonts.outfit(color: AppColors.textSecondary),
      enabledBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(14),
        borderSide: const BorderSide(color: AppColors.glassBorder, width: 1),
      ),
      focusedBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(14),
        borderSide: const BorderSide(color: AppColors.primary, width: 1.5),
      ),
      errorBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(14),
        borderSide: const BorderSide(color: AppColors.error, width: 1),
      ),
      focusedErrorBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(14),
        borderSide: const BorderSide(color: AppColors.error, width: 1.5),
      ),
      contentPadding: const EdgeInsets.symmetric(horizontal: 18, vertical: 16),
    ),
    elevatedButtonTheme: ElevatedButtonThemeData(
      style: ElevatedButton.styleFrom(
        backgroundColor: AppColors.primary,
        foregroundColor: Colors.white,
        minimumSize: const Size(double.infinity, 52),
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(14)),
        textStyle:
            GoogleFonts.outfit(fontWeight: FontWeight.w600, fontSize: 16),
        elevation: 0,
      ),
    ),
    floatingActionButtonTheme: const FloatingActionButtonThemeData(
      backgroundColor: AppColors.primary,
      foregroundColor: Colors.white,
    ),
    chipTheme: ChipThemeData(
      backgroundColor: AppColors.glassWhite,
      labelStyle: GoogleFonts.outfit(color: AppColors.textPrimary),
      side: const BorderSide(color: AppColors.glassBorder),
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
    ),
    snackBarTheme: SnackBarThemeData(
      backgroundColor: AppColors.bg2,
      contentTextStyle: GoogleFonts.outfit(color: AppColors.textPrimary),
      behavior: SnackBarBehavior.floating,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
    ),
    bottomNavigationBarTheme: const BottomNavigationBarThemeData(
      backgroundColor: Color(0xCC0D0B1E),
      selectedItemColor: AppColors.primary,
      unselectedItemColor: AppColors.textHint,
      type: BottomNavigationBarType.fixed,
      elevation: 0,
    ),
  );
}
