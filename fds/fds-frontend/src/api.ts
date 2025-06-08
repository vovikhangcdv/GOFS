import axios from 'axios';
import { AddressTotals, SuspiciousTransfer, BlacklistedAddress, RelatedAddresses } from './types';

// Create axios instance with base URL from environment variable
const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8888',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add response interceptor for debugging
api.interceptors.response.use(
  (response) => {
    console.log('API Response:', response.config.url, response.status);
    return response;
  },
  (error) => {
    console.error('API Error:', error.config?.url, error.response?.status, error.message);
    return Promise.reject(error);
  }
);

export const getAddressTotals = async (address: string): Promise<AddressTotals> => {
  try {
    const response = await api.get<AddressTotals>('/api/address/totals', {
      params: { address },
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching address totals:', error);
    throw error;
  }
};

export const getSuspiciousTransactions = async (): Promise<SuspiciousTransfer[]> => {
  try {
    const response = await axios.get('/api/suspicious');
    // Map PascalCase to camelCase
    return response.data.map((item: any) => ({
      id: item.ID,
      from_address: item.From,
      to_address: item.To,
      amount: item.Amount,
      txHash: item.TxHash,
      blockNumber: item.BlockNumber,
      timestamp: item.Timestamp,
      reason: item.Reason,
      severity: item.Severity,
      details: item.Details,
      isBlacklisted: item.IsBlacklisted,
      createdAt: item.CreatedAt,
      updatedAt: item.UpdatedAt,
    }));
  } catch (error) {
    console.error('Error fetching suspicious transactions:', error);
    throw error;
  }
};

export const getBlacklist = async (): Promise<BlacklistedAddress[]> => {
  try {
    const response = await axios.get('/api/blacklist');
    // Map PascalCase to camelCase
    return response.data.map((item: any) => ({
      id: item.ID,
      address: item.Address,
      reason: item.Reason,
      severity: item.Severity,
      blockNumber: item.BlockNumber,
      txHash: item.TxHash,
      createdAt: item.CreatedAt,
      updatedAt: item.UpdatedAt,
    }));
  } catch (error) {
    console.error('Error fetching blacklist:', error);
    throw error;
  }
};

export const getRelatedAddresses = async (address: string): Promise<RelatedAddresses> => {
  try {
    const response = await api.get<RelatedAddresses>('/api/address/related', {
      params: { address },
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching related addresses:', error);
    throw error;
  }
};

export const getTransactionsByAddress = async (address: string) => {
  try {
    const response = await axios.get('/api/address/transactions', { params: { address } });
    // Map PascalCase to camelCase for frontend
    return response.data.map((item: any) => ({
      id: item.ID,
      txHash: item.Hash,
      from_address: item.From,
      to_address: item.To,
      amount: item.Value,
      blockNumber: item.BlockNumber,
      timestamp: item.Timestamp,
      isAnalyzed: item.IsAnalyzed,
      isPending: item.IsPending,
      status: item.Status,
      createdAt: item.CreatedAt,
      updatedAt: item.UpdatedAt,
      deletedAt: item.DeletedAt,
    }));
  } catch (error) {
    console.error('Error fetching transactions by address:', error);
    throw error;
  }
};

export const getRelatedTransactionsOfSuspicious = async (txHash: string) => {
  try {
    const response = await axios.get('/api/suspicious/related', { params: { txHash } });
    // Map PascalCase to camelCase
    return response.data.map((item: any) => ({
      id: item.ID,
      from_address: item.From,
      to_address: item.To,
      amount: item.Value,
      txHash: item.Hash,
      blockNumber: item.BlockNumber,
      timestamp: item.Timestamp,
      isAnalyzed: item.IsAnalyzed,
      isPending: item.IsPending,
      status: item.Status,
      createdAt: item.CreatedAt,
      updatedAt: item.UpdatedAt,
    }));
  } catch (error) {
    console.error('Error fetching related transactions of suspicious:', error);
    throw error;
  }
};

export const blacklistAddress = async (address: string) => {
  const res = await axios.post('/api/blacklist', { addresses: [address] });
  return res.data;
};

export const unblacklistAddress = async (address: string) => {
  const res = await axios.post('/api/unblacklist', { addresses: [address] });
  return res.data;
};

export const deleteBlacklistAddress = async (address: string) => {
  const res = await axios.delete('/api/blacklist', { data: { address } });
  return res.data;
};

export const getETHBalance = async (address: string) => {
  try {
    const response = await axios.get('/api/balance/eth', { params: { address } });
    return response.data;
  } catch (error) {
    console.error('Error fetching ETH balance:', error);
    throw error;
  }
};

export const getEVNDBalance = async (address: string) => {
  try {
    const response = await axios.get('/api/balance/evnd', { params: { address } });
    return response.data;
  } catch (error) {
    console.error('Error fetching eVND balance:', error);
    throw error;
  }
};

function formatAmount(amount: string | number): string {
  // Convert to number, divide by 1e18, and show up to 6 decimals
  return (Number(amount) / 1e18).toLocaleString(undefined, { maximumFractionDigits: 6 });
} 