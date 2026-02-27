import 'dart:convert';
import 'package:http/http.dart' as http;

class ApiException implements Exception {
  final String message;
  final int? statusCode;
  ApiException(this.message, {this.statusCode});
  @override
  String toString() => 'ApiException($statusCode): $message';
}

class ApiClient {
  static const String baseUrl = 'http://localhost:8000/api/v1';

  String? _token;

  void setToken(String? token) => _token = token;

  Map<String, String> get _headers => {
        'Content-Type': 'application/json',
        if (_token != null) 'Authorization': 'Bearer $_token',
      };

  Future<Map<String, dynamic>> get(String path,
      {Map<String, String>? queryParams}) async {
    final uri =
        Uri.parse('$baseUrl$path').replace(queryParameters: queryParams);
    final res = await http.get(uri, headers: _headers);
    return _handle(res);
  }

  Future<Map<String, dynamic>> post(String path,
      Map<String, dynamic> body) async {
    final uri = Uri.parse('$baseUrl$path');
    final res =
        await http.post(uri, headers: _headers, body: jsonEncode(body));
    return _handle(res);
  }

  Map<String, dynamic> _handle(http.Response res) {
    final body = jsonDecode(res.body) as Map<String, dynamic>;
    if (res.statusCode >= 200 && res.statusCode < 300) return body;
    final msg = body['error'] ?? body['message'] ?? 'Unknown error';
    throw ApiException(msg.toString(), statusCode: res.statusCode);
  }
}

// Singleton instance
final apiClient = ApiClient();
