import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../../models/student_model.dart';
import '../../services/students_service.dart';
import '../../theme/app_theme.dart';
import '../../widgets/app_background.dart';
import '../../widgets/glass_list_tile.dart';

class StudentsScreen extends StatefulWidget {
  const StudentsScreen({super.key});

  @override
  State<StudentsScreen> createState() => _StudentsScreenState();
}

class _StudentsScreenState extends State<StudentsScreen> {
  final _service = StudentsService();
  List<StudentListItem> _students = [];
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
      final list = await _service.getStudents();
      setState(() { _students = list; _loading = false; });
    } catch (e) {
      setState(() { _error = e.toString(); _loading = false; });
    }
  }

  @override
  Widget build(BuildContext context) {
    return AppBackground(
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Students'),
          leading: IconButton(
              icon: const Icon(Icons.arrow_back_rounded),
              onPressed: () => context.go('/dashboard')),
        ),
        floatingActionButton: FloatingActionButton.extended(
          onPressed: () async {
            await context.push('/dashboard/students/add');
            _load();
          },
          icon: const Icon(Icons.add),
          label: const Text('Add Student'),
          backgroundColor: AppColors.primary,
        ),
        body: _loading
            ? const Center(child: CircularProgressIndicator(color: AppColors.primary))
            : _error != null
                ? Center(child: Text(_error!, style: const TextStyle(color: AppColors.error)))
                : _students.isEmpty
                    ? const Center(
                        child: Text('No students yet',
                            style: TextStyle(color: AppColors.textSecondary)))
                    : RefreshIndicator(
                        onRefresh: _load,
                        color: AppColors.primary,
                        child: ListView.separated(
                          padding: const EdgeInsets.fromLTRB(20, 16, 20, 100),
                          itemCount: _students.length,
                          separatorBuilder: (_, __) => const SizedBox(height: 10),
                          itemBuilder: (_, i) {
                            final s = _students[i];
                            return GlassListTile(
                              icon: Icons.person_rounded,
                              title: s.name,
                              subtitle: 'Semester ${s.semesterId} · Dept ${s.departmentId}',
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
