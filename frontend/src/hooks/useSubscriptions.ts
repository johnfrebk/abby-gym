import { useState, useEffect } from 'preact/hooks';
import { Subscription, SubscriptionForm } from '../types';
import toast from 'react-hot-toast';
import { subscriptions } from '../services/api';

export function useSubscriptions() {
  const [subscriptionList, setSubscriptions] = useState<Subscription[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const getAll = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await subscriptions.getAll();
      setSubscriptions(data);
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
    } finally {
      setLoading(false);
    }
  };

  const create = async (subscriptionData: SubscriptionForm): Promise<boolean> => {
    try {
      await subscriptions.create(subscriptionData);
      toast.success('Suscripcion creada exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  const update = async (id: number, subscriptionData: Partial<Subscription>): Promise<boolean> => {
    try {
      await subscriptions.update(id, {
        client_id: subscriptionData.client_id!,
        membership_id: subscriptionData.membership_id!,
        start_date: subscriptionData.start_date!,
        end_date: subscriptionData.end_date!,
      });
      toast.success('Suscripcion actualizada exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  const remove = async (id: number): Promise<boolean> => {
    try {
      await subscriptions.remove(id);
      toast.success('Suscripcion eliminada exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  useEffect(() => {
    getAll();
  }, []);

  return {
    subscriptions: subscriptionList,
    loading,
    error,
    getAll,
    create,
    update,
    remove,
  };
}
