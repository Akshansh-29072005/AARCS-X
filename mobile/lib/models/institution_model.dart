class InstitutionListItem {
  final int id;
  final String name;
  final String code;
  InstitutionListItem({required this.id, required this.name, required this.code});
  factory InstitutionListItem.fromJson(Map<String, dynamic> j) =>
      InstitutionListItem(id: j['id'], name: j['name'], code: j['code']);
}

class CreateInstitutionRequest {
  final String name;
  final String code;
  final String password;
  CreateInstitutionRequest({required this.name, required this.code, required this.password});
  Map<String, dynamic> toJson() => {'name': name, 'code': code, 'password': password};
}
