class StudentListItem {
  final int id;
  final String name;
  final int semesterId;
  final int departmentId;
  final int institutionId;
  StudentListItem({
    required this.id,
    required this.name,
    required this.semesterId,
    required this.departmentId,
    required this.institutionId,
  });
  factory StudentListItem.fromJson(Map<String, dynamic> j) => StudentListItem(
        id: j['id'],
        name: j['name'],
        semesterId: j['semester_id'],
        departmentId: j['department_id'],
        institutionId: j['institution_id'],
      );
}

class CreateStudentRequest {
  final String name;
  final String email;
  final String phone;
  final String password;
  final int semesterId;
  final int departmentId;
  final int institutionId;
  CreateStudentRequest({
    required this.name,
    required this.email,
    required this.phone,
    required this.password,
    required this.semesterId,
    required this.departmentId,
    required this.institutionId,
  });
  Map<String, dynamic> toJson() => {
        'name': name,
        'email': email,
        'phone': phone,
        'password': password,
        'semester_id': semesterId,
        'department_id': departmentId,
        'institution_id': institutionId,
      };
}
