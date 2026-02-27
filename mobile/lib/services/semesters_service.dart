import '../models/semester_model.dart';
import 'api_client.dart';

class SemestersService {
  Future<List<SemesterListItem>> getSemesters({int? departmentId}) async {
    final params = <String, String>{};
    if (departmentId != null) params['department_id'] = departmentId.toString();
    final data = await apiClient.get('/semesters', queryParams: params);
    final list = data['semesters'] as List<dynamic>;
    return list.map((e) => SemesterListItem.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<void> createSemester(CreateSemesterRequest req) async {
    await apiClient.post('/semesters', req.toJson());
  }
}
