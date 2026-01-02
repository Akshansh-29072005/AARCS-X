const API_BASE_URL = "http://YOUR_BACKEND_IP:PORT/api"; 

// -----------------------------
// Types
// -----------------------------

export interface LoginPayload {
  email: string;
  password: string;
}

export interface SignupPayload {
  name: string;
  email: string;
  password: string;
}

export interface AuthResponse {
  success: boolean;
  message: string;
  token?: string;   // JWT (optional for now)
  user?: {
    id: string;
    name: string;
    email: string;
  };
}

// -----------------------------
// Helper (generic request)
// -----------------------------

async function request<T>(
  endpoint: string,
  options: RequestInit
): Promise<T> {
  const res = await fetch(`${API_BASE_URL}${endpoint}`, {
    headers: {
      "Content-Type": "application/json",
    },
    ...options,
  });

  const data = await res.json();

  if (!res.ok) {
    // backend-controlled error
    throw new Error(data?.message || "Something went wrong");
  }

  return data as T;
}

// -----------------------------
// Auth APIs
// -----------------------------

export async function login(payload: LoginPayload): Promise<AuthResponse> {
  return request<AuthResponse>("/auth/login", {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export async function signup(payload: SignupPayload): Promise<AuthResponse> {
  return request<AuthResponse>("/auth/signup", {
    method: "POST",
    body: JSON.stringify(payload),
  });
}
