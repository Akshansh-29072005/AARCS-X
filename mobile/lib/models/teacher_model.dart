class TeacherListItem {
  final int id;
  final String name;
  final int departmentId;
  final String designation;
  TeacherListItem({
    required this.id,
    required this.name,
    required this.departmentId,
    required this.designation,
  });
  factory TeacherListItem.fromJson(Map<String, dynamic> j) => TeacherListItem(
        id: j['id'],
        name: j['name'],
        departmentId: j['department_id'],
        designation: j['designation'],
      );
}

class CreateTeacherRequest {
  final String name;
  final String email;
  final String phone;
  final String password;
  final int departmentId;
  final String designation;
  CreateTeacherRequest({
    required this.name,
    required this.email,
    required this.phone,
    required this.password,
    required this.departmentId,
    required this.designation,
  });
  Map<String, dynamic> toJson() => {
        'name': name,
        'email': email,
        'phone': phone,
        'password': password,
        'department_id': departmentId,
        'designation': designation,
      };
}
