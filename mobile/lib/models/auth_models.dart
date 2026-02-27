import 'dart:convert';

class LoginRequest {
  final String email;
  final String password;
  LoginRequest({required this.email, required this.password});
  Map<String, dynamic> toJson() => {'email': email, 'password': password};
}

class UserResponse {
  final int id;
  final String role;
  final String email;
  UserResponse({required this.id, required this.role, required this.email});

  factory UserResponse.fromJson(Map<String, dynamic> j) =>
      UserResponse(id: j['id'] ?? 0, role: j['role'] ?? '', email: j['email'] ?? '');

  /// Decode directly from a JWT payload (backend returns token-only responses).
  factory UserResponse.fromJwt(String token, String email) {
    try {
      final parts = token.split('.');
      if (parts.length != 3) throw const FormatException('Invalid JWT');
      // Base64url decode the payload
      String payload = parts[1];
      // Pad to make valid base64
      payload += '=' * ((4 - payload.length % 4) % 4);
      final decoded = utf8.decode(base64Url.decode(payload));
      final Map<String, dynamic> claims = jsonDecode(decoded);
      return UserResponse(
        id: (claims['sub'] as num?)?.toInt() ?? 0,
        role: (claims['role'] as String?) ?? 'institution',
        email: email,
      );
    } catch (_) {
      return UserResponse(id: 0, role: 'institution', email: email);
    }
  }
}

/// Backend returns {"token": "..."} for both login and register.
/// The user field is decoded from the JWT payload.
class LoginResponse {
  final String token;
  final UserResponse user;
  LoginResponse({required this.token, required this.user});

  factory LoginResponse.fromJson(Map<String, dynamic> j, {String email = ''}) {
    final token = j['token'] as String;
    // If the backend ever adds a user object, use it; otherwise decode JWT
    final userJson = j['user'];
    final user = userJson != null
        ? UserResponse.fromJson(userJson as Map<String, dynamic>)
        : UserResponse.fromJwt(token, email);
    return LoginResponse(token: token, user: user);
  }
}

class RegisterRequest {
  final String email;
  final String password;
  final String institutionName;
  final String institutionCode;
  RegisterRequest({
    required this.email,
    required this.password,
    required this.institutionName,
    required this.institutionCode,
  });
  Map<String, dynamic> toJson() => {
        'email': email,
        'password': password,
        'institution': institutionName,
        'institution_code': institutionCode,
      };
}
