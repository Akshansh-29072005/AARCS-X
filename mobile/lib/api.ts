const BASE_URL = "http://192.168.31.102:8082";

export async function getStudents() {
  const res = await fetch(`${BASE_URL}/api/students`);
  if (!res.ok) {
    throw new Error("Failed to fetch students");
  }
  return res.json();
}
