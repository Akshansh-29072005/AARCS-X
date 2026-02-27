import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../models/auth_models.dart';
import '../services/auth_service.dart';
import '../services/api_client.dart';

class AuthProvider extends ChangeNotifier {
  final _authService = AuthService();

  String? _token;
  UserResponse? _user;
  bool _loading = false;
  String? _error;

  String? get token => _token;
  UserResponse? get user => _user;
  bool get isAuthenticated => _token != null;
  bool get loading => _loading;
  String? get error => _error;

  AuthProvider() {
    _restoreSession();
  }

  Future<void> _restoreSession() async {
    final prefs = await SharedPreferences.getInstance();
    final savedToken = prefs.getString('auth_token');
    final savedEmail = prefs.getString('auth_email') ?? '';
    if (savedToken != null) {
      _token = savedToken;
      _user = UserResponse.fromJwt(savedToken, savedEmail);
      apiClient.setToken(savedToken);
      notifyListeners();
    }
  }

  Future<bool> login(String email, String password) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final res = await _authService.login(LoginRequest(email: email, password: password));
      await _saveSession(res);
      return true;
    } on ApiException catch (e) {
      _error = e.message;
      return false;
    } finally {
      _loading = false;
      notifyListeners();
    }
  }

  Future<bool> register({
    required String email,
    required String password,
    required String institutionName,
    required String institutionCode,
  }) async {
    _loading = true;
    _error = null;
    notifyListeners();
    try {
      final res = await _authService.register(RegisterRequest(
        email: email,
        password: password,
        institutionName: institutionName,
        institutionCode: institutionCode,
      ));
      await _saveSession(res);
      return true;
    } on ApiException catch (e) {
      _error = e.message;
      return false;
    } finally {
      _loading = false;
      notifyListeners();
    }
  }

  Future<void> _saveSession(LoginResponse res) async {
    _token = res.token;
    _user = res.user;
    apiClient.setToken(res.token);
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('auth_token', res.token);
    await prefs.setString('auth_email', res.user.email);
  }

  Future<void> logout() async {
    _token = null;
    _user = null;
    apiClient.setToken(null);
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove('auth_token');
    await prefs.remove('auth_email');
    notifyListeners();
  }
}
