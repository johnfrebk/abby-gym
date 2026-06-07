const API_BASE = import.meta.env.VITE_API_URL || '';

class ApiError extends Error {
  status: number;
  constructor(message: string, status: number) {
    super(message);
    this.status = status;
  }
}

function getToken(): string | null {
  return localStorage.getItem('abbygym_token');
}

export function setToken(token: string) {
  localStorage.setItem('abbygym_token', token);
}

export function clearToken() {
  localStorage.removeItem('abbygym_token');
}

export function isAuthenticated(): boolean {
  return !!getToken();
}

async function request<T>(method: string, path: string, body?: unknown): Promise<T> {
  const headers: Record<string, string> = { 'Content-Type': 'application/json' };
  const token = getToken();
  if (token) headers['Authorization'] = `Bearer ${token}`;

  const res = await fetch(`${API_BASE}${path}`, {
    method,
    headers,
    body: body ? JSON.stringify(body) : undefined,
  });

  if (!res.ok) {
    const data = await res.json().catch(() => ({}));
    throw new ApiError(data.error || `Error ${res.status}`, res.status);
  }

  return res.json();
}

// Auth
export const auth = {
  login: (email: string, password: string) =>
    request<{ token: string; user: { id: number; email: string; name: string } }>('POST', '/api/auth/login', { email, password }),
  register: (email: string, password: string, name: string) =>
    request<{ token: string; user: { id: number; email: string; name: string } }>('POST', '/api/auth/register', { email, password, name }),
};

// Clients
export const clients = {
  getAll: () => request<any[]>('GET', '/api/clients'),
  getPaginated: (page = 1, pageSize = 20) =>
    request<any>('GET', `/api/clients/paginated?page=${page}&page_size=${pageSize}`),
  getById: (id: number) => request<any>('GET', `/api/clients/${id}`),
  create: (data: any) => request<any>('POST', '/api/clients', data),
  update: (id: number, data: any) => request<any>('PUT', `/api/clients/${id}`, data),
  remove: (id: number) => request<any>('DELETE', `/api/clients/${id}`),
};

// Products
export const products = {
  getAll: () => request<any[]>('GET', '/api/products'),
  create: (data: any) => request<any>('POST', '/api/products', data),
  update: (id: number, data: any) => request<any>('PUT', `/api/products/${id}`, data),
  remove: (id: number) => request<any>('DELETE', `/api/products/${id}`),
};

// Memberships
export const memberships = {
  getAll: () => request<any[]>('GET', '/api/memberships'),
  create: (data: any) => request<any>('POST', '/api/memberships', data),
  update: (id: number, data: any) => request<any>('PUT', `/api/memberships/${id}`, data),
  remove: (id: number) => request<any>('DELETE', `/api/memberships/${id}`),
};

// Subscriptions
export const subscriptions = {
  getAll: () => request<any[]>('GET', '/api/subscriptions'),
  create: (data: any) => request<any>('POST', '/api/subscriptions', data),
  update: (id: number, data: any) => request<any>('PUT', `/api/subscriptions/${id}`, data),
  remove: (id: number) => request<any>('DELETE', `/api/subscriptions/${id}`),
};

// Sales
export const sales = {
  getAll: () => request<any[]>('GET', '/api/sales'),
  create: (data: any) => request<any>('POST', '/api/sales', data),
  update: (id: number, data: any) => request<any>('PUT', `/api/sales/${id}`, data),
  remove: (id: number) => request<any>('DELETE', `/api/sales/${id}`),
};

// Dashboard
export const dashboard = {
  getStats: () => request<any>('GET', '/api/dashboard'),
  getActivities: () => request<any[]>('GET', '/api/activities'),
};

export { ApiError };
