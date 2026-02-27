import '../models/institution_model.dart';
import 'api_client.dart';

class InstitutionsService {
  Future<List<InstitutionListItem>> getInstitutions({String? name, String? code}) async {
    final params = <String, String>{};
    if (name != null) params['name'] = name;
    if (code != null) params['code'] = code;
    final data = await apiClient.get('/institutions', queryParams: params);
    final list = data['institutions'] as List<dynamic>;
    return list.map((e) => InstitutionListItem.fromJson(e as Map<String, dynamic>)).toList();
  }

  Future<void> createInstitution(CreateInstitutionRequest req) async {
    await apiClient.post('/institutions', req.toJson());
  }
}
