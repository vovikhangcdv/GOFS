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
  CssBaseline,
  ThemeProvider,
  createTheme,
} from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import RefreshIcon from '@mui/icons-material/Refresh';
import { getAddressTotals, getSuspiciousTransactions, getBlacklist, getRelatedAddresses, getRelatedTransactionsOfSuspicious, getTransactionsByAddress } from './api';
import { AddressTotals, SuspiciousTransfer, BlacklistedAddress, RelatedAddresses } from './types';
import axios from 'axios';
import { Header } from './components/layout/Header';
import { Sidebar } from './components/layout/Sidebar';
import { TransactionCard } from './components/transactions/TransactionCard';
import { DashboardStats } from './components/dashboard/DashboardStats';
import { SuspiciousTransactionsList } from './components/transactions/SuspiciousTransactionsList';
import { WalletTab } from './components/dashboard/WalletTab';
import { BlacklistView } from './components/dashboard/BlacklistView';

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

const theme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: '#2196f3',
    },
    background: {
      default: '#f5f5f5',
      paper: '#ffffff',
    },
  },
  typography: {
    fontFamily: '"Inter", "Roboto", "Helvetica", "Arial", sans-serif',
  },
  shape: {
    borderRadius: 8,
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          textTransform: 'none',
        },
      },
    },
  },
});

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

  const renderContent = () => {
    switch (currentPage) {
      case 'dashboard':
        return <DashboardStats />;
      case 'transactions':
        return <SuspiciousTransactionsList />;
      case 'wallet':
        return <WalletTab />;
      case 'blacklist':
        return <BlacklistView />;
      default:
        return <DashboardStats />;
    }
  };

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Box sx={{ display: 'flex' }}>
        <Sidebar onPageChange={setCurrentPage} />
        <Box
          component="main"
          sx={{
            flexGrow: 1,
            p: 3,
            width: { sm: `calc(100% - 240px)` },
            mt: '64px',
          }}
        >
          <Header />
          {renderContent()}
        </Box>
      </Box>
    </ThemeProvider>
  );
}

export default App; 