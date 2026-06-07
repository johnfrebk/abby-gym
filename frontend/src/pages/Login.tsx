import { useState } from 'preact/hooks';
import { useAuth } from '../contexts/AuthContext';

export default function Login() {
  const { login, register } = useAuth();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [name, setName] = useState('');
  const [isRegister, setIsRegister] = useState(false);
  const [error, setError] = useState('');
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: Event) => {
    e.preventDefault();
    setError('');
    setSubmitting(true);
    try {
      if (isRegister) {
        await register(email, password, name);
      } else {
        await login(email, password);
      }
    } catch (err: any) {
      setError(err.message || 'Error al iniciar sesion');
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-900 via-gray-800 to-blue-900">
      <div class="backdrop-blur-xl bg-white/10 border border-white/20 rounded-3xl shadow-2xl p-8 w-full max-w-md">
        <div class="text-center mb-8">
          <div class="w-16 h-16 mx-auto backdrop-blur-xl bg-white/20 rounded-2xl flex items-center justify-center shadow-xl border border-white/20 mb-4">
            <img src="/favicon.png" alt="AbbyGym" class="w-10 h-10 object-contain" />
          </div>
          <h1 class="text-2xl font-bold bg-gradient-to-r from-white via-gray-100 to-blue-200 bg-clip-text text-transparent">
            AbbyGym
          </h1>
          <p class="text-white/60 text-sm mt-1">Gestion de gimnasio</p>
        </div>

        <form onSubmit={handleSubmit} class="space-y-4">
          {isRegister && (
            <div>
              <label class="block text-sm font-medium text-white/80 mb-1">Nombre</label>
              <input
                type="text"
                value={name}
                onInput={(e) => setName((e.target as HTMLInputElement).value)}
                class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl text-white placeholder-white/40 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent backdrop-blur-sm"
                placeholder="Tu nombre"
                required
              />
            </div>
          )}

          <div>
            <label class="block text-sm font-medium text-white/80 mb-1">Email</label>
            <input
              type="email"
              value={email}
              onInput={(e) => setEmail((e.target as HTMLInputElement).value)}
              class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl text-white placeholder-white/40 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent backdrop-blur-sm"
              placeholder="admin@abbygym.com"
              required
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-white/80 mb-1">Password</label>
            <input
              type="password"
              value={password}
              onInput={(e) => setPassword((e.target as HTMLInputElement).value)}
              class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl text-white placeholder-white/40 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent backdrop-blur-sm"
              placeholder="••••••••"
              required
            />
          </div>

          {error && (
            <div class="bg-red-500/20 border border-red-500/50 text-red-200 px-4 py-3 rounded-xl text-sm">
              {error}
            </div>
          )}

          <button
            type="submit"
            disabled={submitting}
            class="w-full py-3 bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-500 hover:to-blue-600 text-white font-semibold rounded-xl shadow-lg shadow-blue-500/25 border border-white/20 transition-all duration-300 hover:scale-[1.02] disabled:opacity-50"
          >
            {submitting ? 'Cargando...' : isRegister ? 'Registrarse' : 'Iniciar Sesion'}
          </button>
        </form>

        <div class="mt-6 text-center">
          <button
            onClick={() => { setIsRegister(!isRegister); setError(''); }}
            class="text-white/60 hover:text-white text-sm transition-colors"
          >
            {isRegister ? 'Ya tienes cuenta? Inicia sesion' : 'No tienes cuenta? Registrate'}
          </button>
        </div>
      </div>
    </div>
  );
}
