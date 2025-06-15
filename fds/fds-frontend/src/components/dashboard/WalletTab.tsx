import { useState, useEffect } from 'react';
import {
  Box,
  Card,
  CardContent,
  Typography,
  TextField,
  Button,
  CircularProgress,
  Alert,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Pagination,
  Stack,
  Snackbar,
  Alert as MuiAlert,
  Grid,
  Tabs,
  Tab,
  styled,
} from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import TableViewIcon from '@mui/icons-material/TableView';
import AccountTreeIcon from '@mui/icons-material/AccountTree';
import { getAddressTotals, getTransactionsByAddress, getBlacklist, blacklistAddress, getETHBalance, getEVNDBalance } from '../../api';
import { TransactionNetworkGraph } from '../transactions/TransactionNetworkGraph';

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
      id={`wallet-tabpanel-${index}`}
      {...other}
    >
      {value === index && <Box>{children}</Box>}
    </div>
  );
}

const StatsCard = styled(Card)(({ theme }) => ({
  background: 'rgba(30, 41, 59, 0.8)',
  backdropFilter: 'blur(20px)',
  border: '1px solid rgba(100, 116, 139, 0.3)',
  borderRadius: 16,
  transition: 'all 0.3s ease',
  '&:hover': {
    transform: 'translateY(-2px)',
    border: '1px solid rgba(100, 116, 139, 0.5)',
    boxShadow: '0 12px 40px rgba(0, 0, 0, 0.4)',
  },
}));

const SearchContainer = styled(Box)(({ theme }) => ({
  display: 'flex',
  gap: theme.spacing(2),
  marginBottom: theme.spacing(3),
  flexWrap: 'wrap',
  alignItems: 'flex-end',
}));

const formatAddress = (address: string | undefined | null): string => {
  if (!address || typeof address !== 'string' || address.length < 10) {
    return address || 'N/A';
  }
  return `${address.slice(0, 8)}...${address.slice(-8)}`;
};

interface WalletTabProps {
  searchAddress?: string;
}

export const WalletTab = ({ searchAddress }: WalletTabProps) => {
  const [address, setAddress] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [totals, setTotals] = useState<{ in: number; out: number } | null>(null);
  const [transactions, setTransactions] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const [tabValue, setTabValue] = useState(0);
  const txsPerPage = 10;
  const [blacklist, setBlacklist] = useState<any[]>([]);
  const [snackbar, setSnackbar] = useState<{ open: boolean; message: string; severity: 'success' | 'error' }>({ open: false, message: '', severity: 'success' });
  const [ethBalance, setEthBalance] = useState<{ balance: number; unit: string } | undefined>(undefined);
  const [evndBalance, setEvndBalance] = useState<{ balance: number; unit: string } | undefined>(undefined);

  useEffect(() => {
    fetchBlacklist();
  }, []);

  useEffect(() => {
    if (searchAddress) {
      setAddress(searchAddress);
      // Auto-search when address is provided
      handleSearchWithAddress(searchAddress);
    }
  }, [searchAddress]);

  const fetchBlacklist = async () => {
    try {
      const data = await getBlacklist();
      setBlacklist(Array.isArray(data) ? data : []);
    } catch (err) {
      // ignore
    }
  };

  const isBlacklisted = address && blacklist.some((b) => b.address?.toLowerCase() === address.toLowerCase());

  const handleBlacklist = async () => {
    if (!address) return;
    
    try {
      await blacklistAddress(address);
      setSnackbar({ open: true, message: 'Address blacklisted successfully', severity: 'success' });
      setBlacklist((prev) => [...prev, { address }]);
    } catch (err) {
      setSnackbar({ open: true, message: 'Failed to blacklist address', severity: 'error' });
    }
  };

  const handleSearchWithAddress = async (searchAddr: string) => {
    setError(null);
    setTotals(null);
    setTransactions([]);
    setEthBalance(undefined);
    setEvndBalance(undefined);
    setPage(1);
    setTabValue(0);
    
    if (!searchAddr) return;
    
    setLoading(true);
    try {
      const [totalsData, ethData, evndData] = await Promise.all([
        getAddressTotals(searchAddr),
        getETHBalance(searchAddr),
        getEVNDBalance(searchAddr)
      ]);
      
      setTotals({ in: totalsData.in, out: totalsData.out });
      setEthBalance(ethData);
      setEvndBalance(evndData);
      
      const txs = await getTransactionsByAddress(searchAddr);
      setTransactions(Array.isArray(txs) ? txs : []);
    } catch (err) {
      setError('Failed to fetch wallet data');
      setTotals(null);
      setTransactions([]);
      setEthBalance(undefined);
      setEvndBalance(undefined);
    } finally {
      setLoading(false);
    }
  };

  const handleSearch = async () => {
    await handleSearchWithAddress(address);
  };

  const handleTabChange = (_event: React.SyntheticEvent, newValue: number) => {
    setTabValue(newValue);
  };

  const paginatedTxs = transactions.slice((page - 1) * txsPerPage, page * txsPerPage);
  const pageCount = Math.ceil(transactions.length / txsPerPage);

  return (
    <Box sx={{ p: { xs: 2, md: 3 } }}>
      <Typography 
        variant="h4" 
        sx={{ 
          mb: 3,
          fontWeight: 700,
          color: 'text.primary',
          fontSize: { xs: '1.75rem', md: '2.125rem' },
        }}
      >
        Wallet Analysis
      </Typography>

      <SearchContainer>
        <TextField
          label="Search Address"
          variant="outlined"
          size="small"
          value={address}
          onChange={(e) => setAddress(e.target.value)}
          placeholder="0x..."
          sx={{ 
            width: { xs: '100%', sm: 400 },
            '& .MuiOutlinedInput-root': {
              background: 'rgba(30, 41, 59, 0.7)',
              '& .MuiOutlinedInput-notchedOutline': {
                borderColor: 'rgba(100, 116, 139, 0.3)',
              },
              '&:hover .MuiOutlinedInput-notchedOutline': {
                borderColor: 'rgba(100, 116, 139, 0.5)',
              },
            },
          }}
        />
        <Button
          variant="contained"
          startIcon={<SearchIcon />}
          onClick={handleSearch}
          disabled={loading || !address}
          sx={{
            background: 'linear-gradient(135deg, #64748b, #475569)',
            '&:hover': {
              background: 'linear-gradient(135deg, #475569, #334155)',
            },
          }}
        >
          Search
        </Button>
        {address && !isBlacklisted && (
          <Button
            variant="outlined"
            onClick={handleBlacklist}
            sx={{
              borderColor: 'error.main',
              color: 'error.main',
              '&:hover': {
                background: 'rgba(239, 68, 68, 0.1)',
                borderColor: 'error.light',
              },
            }}
          >
            Blacklist Address
          </Button>
        )}
        {isBlacklisted && (
          <Typography variant="body2" sx={{ color: 'error.main', alignSelf: 'center' }}>
            ⚠️ This address is blacklisted
          </Typography>
        )}
      </SearchContainer>

      {loading && (
        <Box sx={{ 
          display: 'flex', 
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          minHeight: 400,
          gap: 2,
        }}>
          <CircularProgress 
            size={60}
            sx={{
              color: 'primary.main',
              '& .MuiCircularProgress-circle': {
                strokeLinecap: 'round',
              },
            }}
          />
          <Typography variant="body1" color="text.secondary">
            Loading wallet data...
          </Typography>
        </Box>
      )}
      
      {error && (
        <Alert 
          severity="error"
          sx={{
            background: 'rgba(239, 68, 68, 0.1)',
            border: '1px solid rgba(239, 68, 68, 0.3)',
            borderRadius: 2,
          }}
        >
          {error}
        </Alert>
      )}
      
      {(totals || ethBalance || evndBalance) && (
        <Grid container spacing={3} sx={{ mb: 3 }}>
          {totals && (
            <Grid item xs={12} md={4}>
              <StatsCard>
                <CardContent sx={{ p: 3 }}>
                  <Typography variant="h6" sx={{ fontWeight: 600, mb: 2, color: 'text.primary' }}>
                    Total Amount
                  </Typography>
                  <Stack direction="row" spacing={4}>
                    <Box>
                      <Typography variant="caption" sx={{ color: 'success.light', fontSize: '0.75rem' }}>
                        In
                      </Typography>
                      <Typography variant="h6" sx={{ color: 'success.light', fontWeight: 600 }}>
                        {(totals.in / 1e18).toFixed(4)} eVND
                      </Typography>
                    </Box>
                    <Box>
                      <Typography variant="caption" sx={{ color: 'error.light', fontSize: '0.75rem' }}>
                        Out
                      </Typography>
                      <Typography variant="h6" sx={{ color: 'error.light', fontWeight: 600 }}>
                        {(totals.out / 1e18).toFixed(4)} eVND
                      </Typography>
                    </Box>
                  </Stack>
                </CardContent>
              </StatsCard>
            </Grid>
          )}
          
          {ethBalance && (
            <Grid item xs={12} md={4}>
              <StatsCard>
                <CardContent sx={{ p: 3 }}>
                  <Typography variant="h6" sx={{ fontWeight: 600, mb: 2, color: 'text.primary' }}>
                    ETH Balance
                  </Typography>
                  <Typography variant="h4" sx={{ color: 'info.light', fontWeight: 700 }}>
                    {Number(ethBalance.balance).toFixed(4)} {ethBalance.unit}
                  </Typography>
                </CardContent>
              </StatsCard>
            </Grid>
          )}
          
          {evndBalance && (
            <Grid item xs={12} md={4}>
              <StatsCard>
                <CardContent sx={{ p: 3 }}>
                  <Typography variant="h6" sx={{ fontWeight: 600, mb: 2, color: 'text.primary' }}>
                    eVND Balance
                  </Typography>
                  <Typography variant="h4" sx={{ color: 'warning.light', fontWeight: 700 }}>
                    {Number(evndBalance.balance).toFixed(4)} {evndBalance.unit}
                  </Typography>
                </CardContent>
              </StatsCard>
            </Grid>
          )}
        </Grid>
      )}

      {transactions.length > 0 && (
        <Card sx={{
          background: 'rgba(30, 41, 59, 0.8)',
          backdropFilter: 'blur(20px)',
          border: '1px solid rgba(100, 116, 139, 0.3)',
          borderRadius: 2,
        }}>
          <Box sx={{ borderBottom: '1px solid rgba(100, 116, 139, 0.2)' }}>
            <Tabs 
              value={tabValue} 
              onChange={handleTabChange}
              sx={{
                '& .MuiTab-root': {
                  color: 'text.secondary',
                  '&.Mui-selected': {
                    color: 'primary.light',
                  },
                },
                '& .MuiTabs-indicator': {
                  background: 'linear-gradient(135deg, #64748b, #475569)',
                },
              }}
            >
              <Tab 
                icon={<TableViewIcon />} 
                label="Table View" 
                iconPosition="start"
                sx={{ minHeight: 60 }}
              />
              <Tab 
                icon={<AccountTreeIcon />} 
                label="Network Graph" 
                iconPosition="start"
                sx={{ minHeight: 60 }}
              />
            </Tabs>
          </Box>

          <TabPanel value={tabValue} index={0}>
            <Box sx={{ p: 3 }}>
              <Typography variant="h6" sx={{ mb: 2, color: 'text.primary', fontWeight: 600 }}>
                Related Transactions ({transactions.length})
              </Typography>
              <TableContainer 
                component={Paper}
                sx={{
                  background: 'rgba(30, 41, 59, 0.6)',
                  border: '1px solid rgba(100, 116, 139, 0.2)',
                  borderRadius: 2,
                }}
              >
                <Table size="small">
                  <TableHead sx={{ background: 'rgba(100, 116, 139, 0.15)' }}>
                    <TableRow>
                      <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>Hash</TableCell>
                      <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>From</TableCell>
                      <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>To</TableCell>
                      <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>Amount</TableCell>
                      <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>Block Number</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {paginatedTxs.map((tx, idx) => (
                      <TableRow key={idx} sx={{ '&:hover': { background: 'rgba(100, 116, 139, 0.1)' } }}>
                        <TableCell sx={{ 
                          fontFamily: 'monospace', 
                          color: 'text.secondary',
                          fontSize: '0.8rem',
                        }}>
                          {formatAddress(tx.txHash)}
                        </TableCell>
                        <TableCell sx={{ 
                          fontFamily: 'monospace', 
                          color: 'text.secondary',
                          fontSize: '0.8rem',
                        }}>
                          {formatAddress(tx.from_address)}
                        </TableCell>
                        <TableCell sx={{ 
                          fontFamily: 'monospace', 
                          color: 'text.secondary',
                          fontSize: '0.8rem',
                        }}>
                          {formatAddress(tx.to_address)}
                        </TableCell>
                        <TableCell sx={{ color: 'text.primary' }}>
                          {typeof tx.amount !== 'undefined' && tx.amount !== null && !isNaN(Number(tx.amount)) 
                            ? `${(Number(tx.amount) / 1e18).toFixed(4)} eVND` 
                            : 'N/A'
                          }
                        </TableCell>
                        <TableCell sx={{ color: 'text.secondary' }}>
                          {tx.blockNumber ?? 'N/A'}
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
              {pageCount > 1 && (
                <Box sx={{ display: 'flex', justifyContent: 'center', mt: 3 }}>
                  <Pagination
                    count={pageCount}
                    page={page}
                    onChange={(event, value) => setPage(value)}
                    sx={{
                      '& .MuiPaginationItem-root': {
                        color: 'text.secondary',
                        '&.Mui-selected': {
                          background: 'linear-gradient(135deg, #64748b, #475569)',
                          color: 'white',
                        },
                      },
                    }}
                  />
                </Box>
              )}
            </Box>
          </TabPanel>

          <TabPanel value={tabValue} index={1}>
            <Box sx={{ p: 3 }}>
              <Typography variant="h6" sx={{ mb: 2, color: 'text.primary', fontWeight: 600 }}>
                Transaction Network Graph
              </Typography>
              <TransactionNetworkGraph
                targetAddress={address}
                transactions={transactions}
                ethBalance={ethBalance}
                evndBalance={evndBalance}
              />
            </Box>
          </TabPanel>
        </Card>
      )}

      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={() => setSnackbar({ ...snackbar, open: false })}
      >
        <MuiAlert
          onClose={() => setSnackbar({ ...snackbar, open: false })}
          severity={snackbar.severity}
          sx={{ 
            background: snackbar.severity === 'success' 
              ? 'rgba(34, 197, 94, 0.1)' 
              : 'rgba(239, 68, 68, 0.1)',
            border: `1px solid ${snackbar.severity === 'success' ? '#22c55e' : '#ef4444'}`,
          }}
        >
          {snackbar.message}
        </MuiAlert>
      </Snackbar>
    </Box>
  );
}; 