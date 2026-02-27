import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';
import '../../providers/auth_provider.dart';
import '../../services/system_service.dart';
import '../../services/students_service.dart';
import '../../services/teachers_service.dart';
import '../../theme/app_theme.dart';
import '../../widgets/app_background.dart';
import '../../widgets/glass_card.dart';
import '../../widgets/stat_card.dart';

class DashboardScreen extends StatefulWidget {
  const DashboardScreen({super.key});

  @override
  State<DashboardScreen> createState() => _DashboardScreenState();
}

class _DashboardScreenState extends State<DashboardScreen> {
  final _sysService = SystemService();
  final _studentsService = StudentsService();
  final _teachersService = TeachersService();

  bool _apiOnline = false;
  int _studentCount = 0;
  int _teacherCount = 0;
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    try {
      await _sysService.getHealth();
      setState(() => _apiOnline = true);
    } catch (_) {}
    try {
      final s = await _studentsService.getStudents();
      setState(() => _studentCount = s.length);
    } catch (_) {}
    try {
      final t = await _teachersService.getTeachers();
      setState(() => _teacherCount = t.length);
    } catch (_) {}
    setState(() => _loading = false);
  }

  @override
  Widget build(BuildContext context) {
    final auth = context.watch<AuthProvider>();
    return AppBackground(
      child: Scaffold(
        body: Stack(
          children: [
            const GlowOrb(
                color: AppColors.primary,
                alignment: Alignment.topRight,
                size: 280),
            SafeArea(
              child: RefreshIndicator(
                onRefresh: _load,
                color: AppColors.primary,
                child: CustomScrollView(
                  slivers: [
                    SliverToBoxAdapter(
                      child: Padding(
                        padding: const EdgeInsets.fromLTRB(24, 24, 24, 0),
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            // Header row
                            Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                Column(
                                  crossAxisAlignment: CrossAxisAlignment.start,
                                  children: [
                                    const Text('Dashboard',
                                        style: TextStyle(
                                            color: AppColors.textPrimary,
                                            fontSize: 26,
                                            fontWeight: FontWeight.w700)),
                                    Text(
                                        auth.user?.email ?? 'Admin',
                                        style: const TextStyle(
                                            color: AppColors.textSecondary,
                                            fontSize: 13)),
                                  ],
                                ),
                                IconButton(
                                  onPressed: () {
                                    auth.logout();
                                    context.go('/');
                                  },
                                  icon: const Icon(
                                      Icons.logout_rounded,
                                      color: AppColors.textSecondary),
                                  tooltip: 'Logout',
                                )
                              ],
                            ),
                            const SizedBox(height: 20),
                            // API Health card
                            GlassCard(
                              padding: const EdgeInsets.symmetric(
                                  horizontal: 18, vertical: 14),
                              borderRadius: 16,
                              child: Row(
                                children: [
                                  Container(
                                    width: 10,
                                    height: 10,
                                    decoration: BoxDecoration(
                                      shape: BoxShape.circle,
                                      color: _apiOnline
                                          ? AppColors.success
                                          : AppColors.error,
                                    ),
                                  ),
                                  const SizedBox(width: 10),
                                  Text(
                                    _apiOnline
                                        ? 'API Online'
                                        : 'API Offline',
                                    style: TextStyle(
                                        color: _apiOnline
                                            ? AppColors.success
                                            : AppColors.error,
                                        fontWeight: FontWeight.w600),
                                  ),
                                  const Spacer(),
                                  const Text('AARCS-X Backend',
                                      style: TextStyle(
                                          color: AppColors.textHint,
                                          fontSize: 12)),
                                ],
                              ),
                            ),
                            const SizedBox(height: 24),
                            const Text('Overview',
                                style: TextStyle(
                                    color: AppColors.textPrimary,
                                    fontSize: 18,
                                    fontWeight: FontWeight.w600)),
                            const SizedBox(height: 14),
                            _loading
                                ? const Center(
                                    child: CircularProgressIndicator(
                                        color: AppColors.primary))
                                : GridView.count(
                                    shrinkWrap: true,
                                    physics:
                                        const NeverScrollableScrollPhysics(),
                                    crossAxisCount: 2,
                                    mainAxisSpacing: 12,
                                    crossAxisSpacing: 12,
                                    childAspectRatio: 1.0,
                                    children: [
                                      StatCard(
                                          label: 'Students',
                                          value: '$_studentCount',
                                          icon: Icons.people_rounded,
                                          iconColor: AppColors.primary),
                                      StatCard(
                                          label: 'Teachers',
                                          value: '$_teacherCount',
                                          icon: Icons.school_rounded,
                                          iconColor: AppColors.primaryLight),
                                      StatCard(
                                          label: 'Departments',
                                          value: '—',
                                          icon: Icons.account_tree_rounded,
                                          iconColor: AppColors.success),
                                      StatCard(
                                          label: 'Subjects',
                                          value: '—',
                                          icon: Icons.book_rounded,
                                          iconColor: AppColors.warning),
                                    ],
                                  ),
                            const SizedBox(height: 28),
                            const Text('Manage',
                                style: TextStyle(
                                    color: AppColors.textPrimary,
                                    fontSize: 18,
                                    fontWeight: FontWeight.w600)),
                            const SizedBox(height: 14),
                          ],
                        ),
                      ),
                    ),
                    SliverPadding(
                      padding: const EdgeInsets.fromLTRB(24, 0, 24, 30),
                      sliver: SliverList(
                        delegate: SliverChildListDelegate([
                          _NavTile(
                              icon: Icons.people_rounded,
                              label: 'Students',
                              subtitle: 'View & add students',
                              onTap: () => context.go('/dashboard/students')),
                          const SizedBox(height: 10),
                          _NavTile(
                              icon: Icons.school_rounded,
                              label: 'Teachers',
                              subtitle: 'View & add teachers',
                              onTap: () => context.go('/dashboard/teachers')),
                          const SizedBox(height: 10),
                          _NavTile(
                              icon: Icons.business_rounded,
                              label: 'Institutions',
                              subtitle: 'Manage institutions',
                              onTap: () =>
                                  context.go('/dashboard/institutions')),
                          const SizedBox(height: 10),
                          _NavTile(
                              icon: Icons.account_tree_rounded,
                              label: 'Departments',
                              subtitle: 'Manage departments',
                              onTap: () =>
                                  context.go('/dashboard/departments')),
                          const SizedBox(height: 10),
                          _NavTile(
                              icon: Icons.calendar_month_rounded,
                              label: 'Semesters',
                              subtitle: 'Manage semesters',
                              onTap: () => context.go('/dashboard/semesters')),
                          const SizedBox(height: 10),
                          _NavTile(
                              icon: Icons.book_rounded,
                              label: 'Subjects',
                              subtitle: 'Manage subjects',
                              onTap: () => context.go('/dashboard/subjects')),
                        ]),
                      ),
                    ),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

class _NavTile extends StatelessWidget {
  final IconData icon;
  final String label;
  final String subtitle;
  final VoidCallback onTap;

  const _NavTile({
    required this.icon,
    required this.label,
    required this.subtitle,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      child: GlassCard(
        padding: const EdgeInsets.symmetric(horizontal: 18, vertical: 14),
        borderRadius: 16,
        child: Row(
          children: [
            Container(
              padding: const EdgeInsets.all(10),
              decoration: BoxDecoration(
                color: AppColors.primary.withAlpha(30),
                borderRadius: BorderRadius.circular(10),
              ),
              child: Icon(icon, color: AppColors.primary, size: 22),
            ),
            const SizedBox(width: 14),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(label,
                      style: const TextStyle(
                          color: AppColors.textPrimary,
                          fontWeight: FontWeight.w600,
                          fontSize: 15)),
                  Text(subtitle,
                      style: const TextStyle(
                          color: AppColors.textSecondary, fontSize: 12)),
                ],
              ),
            ),
            const Icon(Icons.chevron_right,
                color: AppColors.textHint, size: 20),
          ],
        ),
      ),
    );
  }
}
