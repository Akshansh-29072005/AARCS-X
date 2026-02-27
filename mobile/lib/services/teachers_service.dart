import '../models/teacher_model.dart';
import 'api_client.dart';

class TeachersService {
  Future<List<TeacherListItem>> getTeachers({
    int? departmentId,
    String? designation,
  }) async {
    final params = <String, String>{};
    if (departmentId != null) params['department_id'] = departmentId.toString();
    if (designation != null) params['designation'] = designation;

    final data = await apiClient.get('/teachers', queryParams: params);
    final list = data['teachers'] as List<dynamic>;
    return list.map((e) => TeacherListItem.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<void> createTeacher(CreateTeacherRequest req) async {
    await apiClient.post('/teachers', req.toJson());
  }
}
