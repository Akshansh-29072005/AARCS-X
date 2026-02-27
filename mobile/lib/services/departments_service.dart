import '../models/department_model.dart';
import 'api_client.dart';

class DepartmentsService {
  Future<List<DepartmentListItem>> getDepartments({int? institutionId}) async {
    final params = <String, String>{};
    if (institutionId != null) params['institution_id'] = institutionId.toString();
    final data = await apiClient.get('/departments', queryParams: params);
    final list = data['departments'] as List<dynamic>;
    return list.map((e) => DepartmentListItem.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<void> createDepartment(CreateDepartmentRequest req) async {
    await apiClient.post('/departments', req.toJson());
  }
}
