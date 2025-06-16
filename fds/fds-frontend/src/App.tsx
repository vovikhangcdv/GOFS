import { useState } from 'react';
import {
  Container,
  TextField,
  Button,
  Typography,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Box,
  Tabs,
  Tab,
  Card,
  CardContent,
  AppBar,
  Toolbar,
  CircularProgress,
  Alert,
  Stack,
  GlobalStyles,
} from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import RefreshIcon from '@mui/icons-material/Refresh';
import { getAddressTotals, getSuspiciousTransactions, getBlacklist, getRelatedAddresses, getRelatedTransactionsOfSuspicious, getTransactionsByAddress } from './api';
import { AddressTotals, SuspiciousTransfer, BlacklistedAddress, RelatedAddresses } from './types';
import axios from 'axios';
import { Header } from './components/layout/Header';
import { Sidebar } from './components/layout/Sidebar';
import { TechBackground } from './components/layout/TechBackground';
import { TransactionCard } from './components/transactions/TransactionCard';
import { DashboardStats } from './components/dashboard/DashboardStats';
import { SuspiciousTransactionsList } from './components/transactions/SuspiciousTransactionsList';
import { WalletTab } from './components/dashboard/WalletTab';
import { BlacklistView } from './components/dashboard/BlacklistView';
import { RulesView } from './components/dashboard/RulesView';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;
  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
    </div>
  );
}

function formatAmount(amount: string | number): string {
  // Convert to number, divide by 1e18, and show up to 6 decimals
  return (Number(amount) / 1e18).toLocaleString(undefined, { maximumFractionDigits: 6 });
}

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

const mockTransactions = [
  {
    hash: '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef',
    type: 'in',
    amount: '1.5 ETH',
    timestamp: '2024-02-20T10:30:00Z',
    status: 'completed',
    from: '0xabcdef1234567890abcdef1234567890abcdef12',
    to: '0x1234567890abcdef1234567890abcdef12345678',
  },
  {
    hash: '0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890',
    type: 'out',
    amount: '0.5 ETH',
    timestamp: '2024-02-20T09:15:00Z',
    status: 'pending',
    from: '0x1234567890abcdef1234567890abcdef12345678',
    to: '0xabcdef1234567890abcdef1234567890abcdef12',
  },
];

function App() {
  const [currentPage, setCurrentPage] = useState('dashboard');
  const [searchAddress, setSearchAddress] = useState<string | undefined>(undefined);

  const handleSearchAddress = (address: string) => {
    setSearchAddress(address);
    setCurrentPage('wallet');
  };

  const handlePageChange = (page: string) => {
    // Clear search address when manually navigating to wallet (not through search)
    if (page === 'wallet' && currentPage !== 'wallet') {
      setSearchAddress(undefined);
    }
    setCurrentPage(page);
  };

  const renderContent = () => {
    switch (currentPage) {
      case 'dashboard':
        return <DashboardStats onPageChange={handlePageChange} />;
      case 'transactions':
        return <SuspiciousTransactionsList />;
      case 'wallet':
        return <WalletTab searchAddress={searchAddress} />;
      case 'rules':
        return <RulesView />;
      case 'blacklist':
        return <BlacklistView />;
      default:
        return <DashboardStats onPageChange={handlePageChange} />;
    }
  };

  return (
    <>
      <GlobalStyles
        styles={{
          'body': {
            margin: 0,
            fontFamily: '"Inter", "Roboto", "Helvetica", "Arial", sans-serif',
            background: '#0a0a0f',
          },
          '*::-webkit-scrollbar': {
            width: '8px',
          },
          '*::-webkit-scrollbar-track': {
            background: 'rgba(30, 41, 59, 0.3)',
            borderRadius: '4px',
          },
          '*::-webkit-scrollbar-thumb': {
            background: 'rgba(59, 130, 246, 0.6)',
            borderRadius: '4px',
          },
        }}
      />
      <Box sx={{ 
        display: 'flex',
        minHeight: '100vh',
        background: 'linear-gradient(135deg, #0a0a0f 0%, #1a1a2e 50%, #0f1419 100%)',
        position: 'relative',
        overflow: 'hidden',
      }}>
        <TechBackground />
        <Header onSearchAddress={handleSearchAddress} />
        <Sidebar onPageChange={handlePageChange} />
        <Box
          component="main"
          sx={{
            flexGrow: 1,
            p: 0.5,
            ml: '40px',
            mt: '64px',
            minHeight: 'calc(100vh - 64px)',
            width: 'calc(100vw - 238px)',
            position: 'relative',
          }}
        >
          {renderContent()}
        </Box>
      </Box>
    </>
  );
}

export default App; 