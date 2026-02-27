import '../models/subject_model.dart';
import 'api_client.dart';

class SubjectsService {
  Future<List<SubjectListItem>> getSubjects({int? semesterId}) async {
    final params = <String, String>{};
    if (semesterId != null) params['semester_id'] = semesterId.toString();
    final data = await apiClient.get('/subjects', queryParams: params);
    final list = data['subjects'] as List<dynamic>;
    return list.map((e) => SubjectListItem.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<void> createSubject(CreateSubjectRequest req) async {
    await apiClient.post('/subjects', req.toJson());
  }
}
