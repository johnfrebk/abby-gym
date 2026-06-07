import { useState, useEffect } from 'preact/hooks';
import { Product, ProductForm } from '../types';
import toast from 'react-hot-toast';
import { products } from '../services/api';

export function useProducts() {
  const [productList, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const getAll = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await products.getAll();
      setProducts(data);
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
    } finally {
      setLoading(false);
    }
  };

  const create = async (productData: ProductForm): Promise<boolean> => {
    try {
      await products.create(productData);
      toast.success('Producto creado exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  const update = async (id: number, productData: ProductForm): Promise<boolean> => {
    try {
      await products.update(id, productData);
      toast.success('Producto actualizado exitosamente');
      return true;
    } catch (err: any) {
      setError(err.message);
      toast.error(err.message);
      return false;
    }
  };

  const remove = async (id: number): Promise<boolean> => {
    try {
      await products.remove(id);
      toast.success('Producto eliminado exitosamente');
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
    products: productList,
    loading,
    error,
    getAll,
    create,
    update,
    remove,
  };
}
