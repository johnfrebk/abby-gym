import { useState, useEffect } from 'preact/hooks';
import { Membership, MembershipForm } from '../types';
import toast from 'react-hot-toast';
import { memberships } from '../services/api';

export function useMemberships() {
  const [membershipList, setMemberships] = useState<Membership[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const getAll = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await memberships.getAll();
      setMemberships(data);
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
    } finally {
      setLoading(false);
    }
  };

  const create = async (membershipData: MembershipForm): Promise<boolean> => {
    try {
      await memberships.create(membershipData);
      toast.success('Membresia creada exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  const update = async (id: number, membershipData: MembershipForm): Promise<boolean> => {
    try {
      await memberships.update(id, membershipData);
      toast.success('Membresia actualizada exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  const remove = async (id: number): Promise<boolean> => {
    try {
      await memberships.remove(id);
      toast.success('Membresia eliminada exitosamente');
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
    memberships: membershipList,
    loading,
    error,
    getAll,
    create,
    update,
    remove,
  };
}
