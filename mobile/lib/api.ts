const API_BASE_URL = "http://192.168.31.102:8000/api"; // TODO: Update this when backend is deployed

// -----------------------------
// Types
// -----------------------------

export type UserRole = "student" | "faculty" | "institute";

export interface SignupPayload {
  first_name: string;
  last_name: string;
  email: string;
  phone: string;
  semester: number;
  branch: string;
  role: UserRole; // Added role
}

export interface TicketPayload {
  studentId: string;
  issueType: string;
  description: string;
}

// -----------------------------
// Validation Helpers
// -----------------------------

export function validateSignup(payload: SignupPayload): string | null {
  if (!payload.first_name || !payload.last_name) return "Name is required";
  if (!payload.email.includes("@")) return "Invalid email";
  if (payload.phone.length < 10) return "Invalid phone number";
  if (payload.role === "student" && (!payload.semester || !payload.branch)) {
    return "Semester and Branch are required for students";
  }
  return null;
}

export function validateTicket(payload: TicketPayload): string | null {
  if (!payload.description || payload.description.length < 10) {
    return "Description must be at least 10 characters";
  }
  if (!payload.issueType) return "Please select an issue type";
  return null;
}

// -----------------------------
// Helper (generic request)
// -----------------------------

async function request<T>(
  endpoint: string,
  options: RequestInit
): Promise<T> {
  // MOCK MODE: Return mock data if API is not reachable or for testing
  // Remove this block when backend is fully ready
  // console.log(`[Mock API] Request to ${endpoint}`, options);

  try {
    const res = await fetch(`${API_BASE_URL}${endpoint}`, {
      headers: {
        "Content-Type": "application/json",
      },
      ...options,
    });

    const data = await res.json();

    if (!res.ok) {
      throw new Error(data?.message || "Something went wrong");
    }

    return data as T;
  } catch (error) {
    console.error("API Request Error:", error);
    throw error;
  }
}

// -----------------------------
// Auth APIs
// -----------------------------

export async function signup(payload: SignupPayload): Promise<void> {
  const error = validateSignup(payload);
  if (error) throw new Error(error);

  return request<void>("/students", {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

// -----------------------------
// Student APIs
// -----------------------------

export async function submitTicket(payload: TicketPayload): Promise<void> {
  const error = validateTicket(payload);
  if (error) throw new Error(error);

  // TODO: Replace with actual backend endpoint
  console.log("Submitting ticket:", payload);
  return Promise.resolve();
}