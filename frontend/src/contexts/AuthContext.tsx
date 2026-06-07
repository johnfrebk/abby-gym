import { createContext } from 'preact';
import { useContext, useState, useEffect } from 'preact/hooks';
import { auth, setToken, clearToken, isAuthenticated } from '../services/api';

interface User {
  id: number;
  email: string;
  name: string;
}

interface AuthContextType {
  user: User | null;
  loading: boolean;
  login: (email: string, password: string) => Promise<void>;
  register: (email: string, password: string, name: string) => Promise<void>;
  logout: () => void;
  isAuth: boolean;
}

const AuthContext = createContext<AuthContextType>({
  user: null,
  loading: true,
  login: async () => {},
  register: async () => {},
  logout: () => {},
  isAuth: false,
});

export function AuthProvider({ children }: { children: preact.ComponentChildren }) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const stored = localStorage.getItem('abbygym_user');
    if (stored) {
      try { setUser(JSON.parse(stored)); } catch { clearToken(); localStorage.removeItem('abbygym_user'); }
    }
    setLoading(false);
  }, []);

  const login = async (email: string, password: string) => {
    const res = await auth.login(email, password);
    setToken(res.token);
    localStorage.setItem('abbygym_user', JSON.stringify(res.user));
    setUser(res.user);
  };

  const register = async (email: string, password: string, name: string) => {
    const res = await auth.register(email, password, name);
    setToken(res.token);
    localStorage.setItem('abbygym_user', JSON.stringify(res.user));
    setUser(res.user);
  };

  const logout = () => {
    clearToken();
    localStorage.removeItem('abbygym_user');
    setUser(null);
  };

  return (
    <AuthContext.Provider value={{ user, loading, login, register, logout, isAuth: !!user }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  return useContext(AuthContext);
}
