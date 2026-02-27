class SubjectListItem {
  final int id;
  final String name;
  final String code;
  final int semesterId;
  SubjectListItem({
    required this.id,
    required this.name,
    required this.code,
    required this.semesterId,
  });
  factory SubjectListItem.fromJson(Map<String, dynamic> j) => SubjectListItem(
        id: j['id'],
        name: j['name'],
        code: j['code'],
        semesterId: j['semester_id'],
      );
}

class CreateSubjectRequest {
  final String name;
  final String code;
  final int semesterId;
  CreateSubjectRequest({required this.name, required this.code, required this.semesterId});
  Map<String, dynamic> toJson() => {'name': name, 'code': code, 'semester_id': semesterId};
}
