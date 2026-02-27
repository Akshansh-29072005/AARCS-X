class DepartmentListItem {
  final int id;
  final String name;
  final String code;
  final String headOfDepartment;
  final int institutionId;
  DepartmentListItem({
    required this.id,
    required this.name,
    required this.code,
    required this.headOfDepartment,
    required this.institutionId,
  });
  factory DepartmentListItem.fromJson(Map<String, dynamic> j) => DepartmentListItem(
        id: j['id'],
        name: j['name'],
        code: j['code'],
        headOfDepartment: j['head_of_department'],
        institutionId: j['institution_id'],
      );
}

class CreateDepartmentRequest {
  final String name;
  final String code;
  final String headOfDepartment;
  final int institutionId;
  CreateDepartmentRequest({
    required this.name,
    required this.code,
    required this.headOfDepartment,
    required this.institutionId,
  });
  Map<String, dynamic> toJson() => {
        'name': name,
        'code': code,
        'head_of_department': headOfDepartment,
        'institution_id': institutionId,
      };
}
