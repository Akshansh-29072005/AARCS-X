import '../models/student_model.dart';
import 'api_client.dart';

class StudentsService {
  Future<List<StudentListItem>> getStudents({
    int? semesterId,
    int? departmentId,
    int? institutionId,
  }) async {
    final params = <String, String>{};
    if (semesterId != null) params['semester_id'] = semesterId.toString();
    if (departmentId != null) params['department_id'] = departmentId.toString();
    if (institutionId != null) params['institution_id'] = institutionId.toString();

    final data = await apiClient.get('/students', queryParams: params);
    final list = data['students'] as List<dynamic>;
    return list.map((e) => StudentListItem.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<void> createStudent(CreateStudentRequest req) async {
    await apiClient.post('/students', req.toJson());
  }
}
