import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../../models/teacher_model.dart';
import '../../services/teachers_service.dart';
import '../../theme/app_theme.dart';
import '../../widgets/app_background.dart';
import '../../widgets/glass_list_tile.dart';

class TeachersScreen extends StatefulWidget {
  const TeachersScreen({super.key});

  @override
  State<TeachersScreen> createState() => _TeachersScreenState();
}

class _TeachersScreenState extends State<TeachersScreen> {
  final _service = TeachersService();
  List<TeacherListItem> _teachers = [];
  bool _loading = true;
  String? _error;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    setState(() { _loading = true; _error = null; });
    try {
      final list = await _service.getTeachers();
      setState(() { _teachers = list; _loading = false; });
    } catch (e) {
      setState(() { _error = e.toString(); _loading = false; });
    }
  }

  @override
  Widget build(BuildContext context) {
    return AppBackground(
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Teachers'),
          leading: IconButton(
              icon: const Icon(Icons.arrow_back_rounded),
              onPressed: () => context.go('/dashboard')),
        ),
        floatingActionButton: FloatingActionButton.extended(
          onPressed: () async {
            await context.push('/dashboard/teachers/add');
            _load();
          },
          icon: const Icon(Icons.add),
          label: const Text('Add Teacher'),
          backgroundColor: AppColors.primary,
        ),
        body: _loading
            ? const Center(child: CircularProgressIndicator(color: AppColors.primary))
            : _error != null
                ? Center(child: Text(_error!, style: const TextStyle(color: AppColors.error)))
                : _teachers.isEmpty
                    ? const Center(
                        child: Text('No teachers yet',
                            style: TextStyle(color: AppColors.textSecondary)))
                    : RefreshIndicator(
                        onRefresh: _load,
                        color: AppColors.primary,
                        child: ListView.separated(
                          padding: const EdgeInsets.fromLTRB(20, 16, 20, 100),
                          itemCount: _teachers.length,
                          separatorBuilder: (_, __) => const SizedBox(height: 10),
                          itemBuilder: (_, i) {
                            final t = _teachers[i];
                            return GlassListTile(
                              icon: Icons.school_rounded,
                              title: t.name,
                              subtitle: '${t.designation} · Dept ${t.departmentId}',
                              trailing: const Icon(Icons.chevron_right,
                                  color: AppColors.textHint, size: 18),
                            );
                          },
                        ),
                      ),
      ),
    );
  }
}
