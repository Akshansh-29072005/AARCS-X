import '../models/auth_models.dart';
import 'api_client.dart';

class AuthService {
  Future<LoginResponse> login(LoginRequest req) async {
    final data = await apiClient.post('/auth/login', req.toJson());
    // Pass email so we can embed it in the UserResponse decoded from JWT
    return LoginResponse.fromJson(data, email: req.email);
  }

  Future<LoginResponse> register(RegisterRequest req) async {
    final data = await apiClient.post('/auth/register', req.toJson());
    return LoginResponse.fromJson(data, email: req.email);
  }
}
