import 'package:go_router/go_router.dart';
import '../providers/auth_provider.dart';

// Auth screens
import '../screens/onboarding_screen.dart';
import '../screens/auth/login_screen.dart';
import '../screens/auth/register_screen.dart';

// Dashboard & tabs
import '../screens/dashboard/dashboard_screen.dart';

// Entity screens
import '../screens/students/students_screen.dart';
import '../screens/students/add_student_screen.dart';
import '../screens/teachers/teachers_screen.dart';
import '../screens/teachers/add_teacher_screen.dart';
import '../screens/institutions/institutions_screen.dart';
import '../screens/departments/departments_screen.dart';
import '../screens/semesters/semesters_screen.dart';
import '../screens/subjects/subjects_screen.dart';

GoRouter buildRouter(AuthProvider auth) {
  return GoRouter(
    refreshListenable: auth,
    redirect: (context, state) {
      final isAuth = auth.isAuthenticated;
      final onPublic = state.matchedLocation == '/' ||
          state.matchedLocation == '/login' ||
          state.matchedLocation == '/register';
      if (!isAuth && !onPublic) return '/login';
      if (isAuth && onPublic && state.matchedLocation != '/') return '/dashboard';
      return null;
    },
    routes: [
      GoRoute(path: '/', builder: (_, __) => const OnboardingScreen()),
      GoRoute(path: '/login', builder: (_, __) => const LoginScreen()),
      GoRoute(path: '/register', builder: (_, __) => const RegisterScreen()),
      GoRoute(
        path: '/dashboard',
        builder: (_, __) => const DashboardScreen(),
        routes: [
          GoRoute(path: 'students', builder: (_, __) => const StudentsScreen(),
              routes: [
                GoRoute(path: 'add', builder: (_, __) => const AddStudentScreen()),
              ]),
          GoRoute(path: 'teachers', builder: (_, __) => const TeachersScreen(),
              routes: [
                GoRoute(path: 'add', builder: (_, __) => const AddTeacherScreen()),
              ]),
          GoRoute(path: 'institutions', builder: (_, __) => const InstitutionsScreen()),
          GoRoute(path: 'departments', builder: (_, __) => const DepartmentsScreen()),
          GoRoute(path: 'semesters', builder: (_, __) => const SemestersScreen()),
          GoRoute(path: 'subjects', builder: (_, __) => const SubjectsScreen()),
        ],
      ),
    ],
  );
}
