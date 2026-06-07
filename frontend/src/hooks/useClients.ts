import { useState, useEffect } from 'preact/hooks';
import { Client, ClientForm, UpdateClientForm, PaginatedResult } from '../types';
import toast from 'react-hot-toast';
import { clients } from '../services/api';

export function useClients(pageSize = 20) {
  const [clientsList, setClients] = useState<Client[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [total, setTotal] = useState(0);

  const getAll = async () => {
    setLoading(true);
    setError(null);
    try {
      const result = await clients.getAll();
      setClients(result);
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
    } finally {
      setLoading(false);
    }
  };

  const fetchPage = async (p: number) => {
    setLoading(true);
    setError(null);
    try {
      const result = await clients.getPaginated(p, pageSize) as unknown as PaginatedResult;
      if (result && result.data) {
        setClients(result.data as Client[]);
        setPage(result.page);
        setTotalPages(result.total_pages);
        setTotal(result.total);
      }
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
    } finally {
      setLoading(false);
    }
  };

  const getById = async (id: number): Promise<Client | null> => {
    try {
      const data = await clients.getById(id);
      return data || null;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return null;
    }
  };

  const create = async (clientData: ClientForm): Promise<boolean> => {
    try {
      await clients.create({ ...clientData, registration_date: new Date().toISOString() });
      toast.success('Cliente creado exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  const update = async (id: number, clientData: UpdateClientForm): Promise<boolean> => {
    try {
      await clients.update(id, clientData);
      toast.success('Cliente actualizado exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  const remove = async (id: number): Promise<boolean> => {
    try {
      await clients.remove(id);
      toast.success('Cliente eliminado exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  useEffect(() => {
    fetchPage(1);
  }, []);

  return {
    clients: clientsList,
    loading,
    error,
    page,
    totalPages,
    total,
    getAll,
    fetchPage,
    getById,
    create,
    update,
    remove,
  };
}
