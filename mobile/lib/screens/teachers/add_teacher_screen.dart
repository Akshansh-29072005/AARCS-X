import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../../models/teacher_model.dart';
import '../../services/teachers_service.dart';
import '../../theme/app_theme.dart';
import '../../widgets/app_background.dart';
import '../../widgets/glass_card.dart';
import '../../widgets/app_text_field.dart';

class AddTeacherScreen extends StatefulWidget {
  const AddTeacherScreen({super.key});

  @override
  State<AddTeacherScreen> createState() => _AddTeacherScreenState();
}

class _AddTeacherScreenState extends State<AddTeacherScreen> {
  final _formKey = GlobalKey<FormState>();
  final _nameCtrl = TextEditingController();
  final _emailCtrl = TextEditingController();
  final _phoneCtrl = TextEditingController();
  final _pwCtrl = TextEditingController();
  final _deptCtrl = TextEditingController();
  final _desigCtrl = TextEditingController();
  bool _loading = false;

  @override
  void dispose() {
    for (final c in [_nameCtrl, _emailCtrl, _phoneCtrl, _pwCtrl, _deptCtrl, _desigCtrl]) {
      c.dispose();
    }
    super.dispose();
  }

  Future<void> _submit() async {
    if (!_formKey.currentState!.validate()) return;
    setState(() => _loading = true);
    try {
      await TeachersService().createTeacher(CreateTeacherRequest(
        name: _nameCtrl.text.trim(),
        email: _emailCtrl.text.trim(),
        phone: _phoneCtrl.text.trim(),
        password: _pwCtrl.text.trim(),
        departmentId: int.parse(_deptCtrl.text.trim()),
        designation: _desigCtrl.text.trim(),
      ));
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('Teacher added successfully!')));
        context.pop();
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context)
            .showSnackBar(SnackBar(content: Text(e.toString())));
      }
    } finally {
      if (mounted) setState(() => _loading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return AppBackground(
      child: Scaffold(
        appBar: AppBar(title: const Text('Add Teacher')),
        body: SingleChildScrollView(
          padding: const EdgeInsets.all(20),
          child: GlassCard(
            child: Form(
              key: _formKey,
              child: Column(
                children: [
                  AppTextField(label: 'Full Name', controller: _nameCtrl,
                      prefixIcon: const Icon(Icons.person_outline),
                      validator: (v) => v!.isEmpty ? 'Required' : null),
                  const SizedBox(height: 14),
                  AppTextField(label: 'Email', controller: _emailCtrl,
                      keyboardType: TextInputType.emailAddress,
                      prefixIcon: const Icon(Icons.email_outlined),
                      validator: (v) => !v!.contains('@') ? 'Invalid email' : null),
                  const SizedBox(height: 14),
                  AppTextField(label: 'Phone', controller: _phoneCtrl,
                      keyboardType: TextInputType.phone,
                      prefixIcon: const Icon(Icons.phone_outlined),
                      validator: (v) => v!.isEmpty ? 'Required' : null),
                  const SizedBox(height: 14),
                  AppTextField(label: 'Password', controller: _pwCtrl,
                      obscureText: true,
                      prefixIcon: const Icon(Icons.lock_outline),
                      validator: (v) => v!.length < 6 ? 'Min 6 chars' : null),
                  const SizedBox(height: 14),
                  AppTextField(label: 'Designation', controller: _desigCtrl,
                      prefixIcon: const Icon(Icons.badge_outlined),
                      validator: (v) => v!.isEmpty ? 'Required' : null),
                  const SizedBox(height: 14),
                  AppTextField(label: 'Department ID', controller: _deptCtrl,
                      keyboardType: TextInputType.number,
                      prefixIcon: const Icon(Icons.account_tree_outlined),
                      validator: (v) => int.tryParse(v ?? '') == null ? 'Number required' : null),
                  const SizedBox(height: 28),
                  _loading
                      ? const CircularProgressIndicator(color: AppColors.primary)
                      : Container(
                          decoration: BoxDecoration(
                            gradient: AppGradients.primaryButton,
                            borderRadius: BorderRadius.circular(14),
                            boxShadow: [BoxShadow(color: AppColors.primary.withAlpha(80),
                                blurRadius: 16, offset: const Offset(0, 5))],
                          ),
                          child: ElevatedButton(
                            onPressed: _submit,
                            style: ElevatedButton.styleFrom(
                              backgroundColor: Colors.transparent,
                              shadowColor: Colors.transparent,
                              minimumSize: const Size(double.infinity, 52),
                              shape: RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(14)),
                            ),
                            child: const Text('Add Teacher',
                                style: TextStyle(color: Colors.white,
                                    fontWeight: FontWeight.w600, fontSize: 16)),
                          ),
                        ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
