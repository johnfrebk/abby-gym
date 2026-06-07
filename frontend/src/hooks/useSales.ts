import { useState, useEffect } from 'preact/hooks';
import { Sale, SaleForm } from '../types';
import toast from 'react-hot-toast';
import { sales } from '../services/api';

export function useSales() {
  const [salesList, setSales] = useState<Sale[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const getAll = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await sales.getAll();
      setSales(data);
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
    } finally {
      setLoading(false);
    }
  };

  const create = async (saleData: SaleForm): Promise<boolean> => {
    try {
      await sales.create(saleData);
      toast.success('Venta registrada exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  const update = async (id: number, saleData: Partial<Sale>): Promise<boolean> => {
    try {
      await sales.update(id, {
        client_id: saleData.client_id!,
        details: saleData.details!,
      });
      toast.success('Venta actualizada exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  const remove = async (id: number): Promise<boolean> => {
    try {
      await sales.remove(id);
      toast.success('Venta eliminada exitosamente');
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
    sales: salesList,
    loading,
    error,
    getAll,
    create,
    update,
    remove,
  };
}
