import { useState, useEffect } from 'preact/hooks';
import { DashboardStats, Activity } from '../types';
import toast from 'react-hot-toast';
import { dashboard } from '../services/api';

export function useDashboard() {
  const [stats, setStats] = useState<DashboardStats | null>(null);
  const [activities, setActivities] = useState<Activity[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const getStats = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await dashboard.getStats();
      setStats(data);
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
    } finally {
      setLoading(false);
    }
  };

  const getActivities = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await dashboard.getActivities();
      setActivities(data);
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getStats();
    getActivities();
  }, []);

  return {
    stats,
    activities,
    loading,
    error,
    getStats,
    getActivities,
  };
}
