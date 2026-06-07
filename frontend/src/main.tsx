import { render } from 'preact';
import App from './App.tsx';
import { AuthProvider } from './contexts/AuthContext.tsx';
import './index.css';

render(
  <AuthProvider>
    <App />
  </AuthProvider>,
  document.getElementById('root')!
);
