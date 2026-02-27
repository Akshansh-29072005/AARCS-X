import 'api_client.dart';

class SystemService {
  Future<Map<String, dynamic>> getHealth() async {
    return await apiClient.get('/system/health');
  }
}
