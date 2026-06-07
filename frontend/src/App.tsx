import { useState } from 'preact/hooks';
import { Toaster } from 'react-hot-toast';
import { useAuth } from './contexts/AuthContext';
import Login from './pages/Login';
import Sidebar from './components/Layout/Sidebar';
import Header from './components/Layout/Header';
import Dashboard from './components/Dashboard/Dashboard';
import ClientsList from './features/clients/ClientsList';
import ProductsList from './features/products/ProductsList';
import MembershipsList from './features/memberships/MembershipsList';
import SubscriptionsList from './features/subscriptions/SubscriptionsList';
import SalesList from './features/sales/SalesList';

const sectionTitles: Record<string, string> = {
  dashboard: 'Dashboard',
  clients: 'Clientes',
  products: 'Productos',
  memberships: 'Membresias',
  subscriptions: 'Suscripciones',
  sales: 'Ventas',
};

export default function App() {
  const { isAuth, loading } = useAuth();
  const [activeSection, setActiveSection] = useState('dashboard');

  if (loading) {
    return (
      <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-900 via-gray-800 to-blue-900">
        <div class="text-white text-lg">Cargando...</div>
      </div>
    );
  }

  if (!isAuth) {
    return <Login />;
  }

  const renderSection = () => {
    switch (activeSection) {
      case 'dashboard': return <Dashboard />;
      case 'clients': return <ClientsList />;
      case 'products': return <ProductsList />;
      case 'memberships': return <MembershipsList />;
      case 'subscriptions': return <SubscriptionsList />;
      case 'sales': return <SalesList />;
      default: return <Dashboard />;
    }
  };

  return (
    <div class="flex h-screen bg-gray-100">
      <Sidebar activeSection={activeSection} onSectionChange={setActiveSection} />
      <div class="flex-1 flex flex-col overflow-hidden">
        <Header title={sectionTitles[activeSection] || 'Dashboard'} />
        <main class="flex-1 overflow-y-auto p-6">
          {renderSection()}
        </main>
      </div>
      <Toaster
        position="top-right"
        toastOptions={{
          duration: 3000,
          style: { background: '#363636', color: '#fff' },
          success: { duration: 3000, iconTheme: { primary: '#10B981', secondary: '#fff' } },
          error: { duration: 4000, iconTheme: { primary: '#EF4444', secondary: '#fff' } },
        }}
      />
    </div>
  );
}
