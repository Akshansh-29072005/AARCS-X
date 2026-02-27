class SemesterListItem {
  final int id;
  final int number;
  final int departmentId;
  SemesterListItem({required this.id, required this.number, required this.departmentId});
  factory SemesterListItem.fromJson(Map<String, dynamic> j) =>
      SemesterListItem(id: j['id'], number: j['number'], departmentId: j['department_id']);
}

class CreateSemesterRequest {
  final int number;
  final int departmentId;
  CreateSemesterRequest({required this.number, required this.departmentId});
  Map<String, dynamic> toJson() => {'number': number, 'department_id': departmentId};
}
