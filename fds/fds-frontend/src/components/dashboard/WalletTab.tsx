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
} from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import { getAddressTotals, getTransactionsByAddress, getBlacklist, blacklistAddress, getETHBalance, getEVNDBalance } from '../../api';

export const WalletTab = () => {
  const [address, setAddress] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [totals, setTotals] = useState<{ in: number; out: number } | null>(null);
  const [transactions, setTransactions] = useState<any[]>([]);
  const [page, setPage] = useState(1);
  const txsPerPage = 10;
  const [blacklist, setBlacklist] = useState<any[]>([]);
  const [snackbar, setSnackbar] = useState<{ open: boolean; message: string; severity: 'success' | 'error' }>({ open: false, message: '', severity: 'success' });
  const [ethBalance, setEthBalance] = useState<{ balance: number; unit: string } | null>(null);
  const [evndBalance, setEvndBalance] = useState<{ balance: number; unit: string } | null>(null);

  useEffect(() => {
    fetchBlacklist();
  }, []);

  const fetchBlacklist = async () => {
    try {
      const data = await getBlacklist();
      setBlacklist(data);
    } catch (err) {
      // ignore
    }
  };

  const isBlacklisted = address && blacklist.some((b) => b.address?.toLowerCase() === address.toLowerCase());

  const handleBlacklist = async () => {
    try {
      await blacklistAddress(address);
      setSnackbar({ open: true, message: 'Address blacklisted successfully', severity: 'success' });
      setBlacklist((prev) => [...prev, { address }]);
    } catch (err) {
      setSnackbar({ open: true, message: 'Failed to blacklist address', severity: 'error' });
    }
  };

  const handleSearch = async () => {
    setError(null);
    setTotals(null);
    setTransactions([]);
    setEthBalance(null);
    setEvndBalance(null);
    setPage(1);
    if (!address) return;
    setLoading(true);
    try {
      const [totalsData, ethData, evndData] = await Promise.all([
        getAddressTotals(address),
        getETHBalance(address),
        getEVNDBalance(address)
      ]);
      setTotals({ in: totalsData.in, out: totalsData.out });
      setEthBalance(ethData);
      setEvndBalance(evndData);
      const txs = await getTransactionsByAddress(address);
      setTransactions(txs);
    } catch (err) {
      setError('Failed to fetch wallet data');
      setTotals(null);
      setTransactions([]);
      setEthBalance(null);
      setEvndBalance(null);
    } finally {
      setLoading(false);
    }
  };

  const paginatedTxs = transactions.slice((page - 1) * txsPerPage, page * txsPerPage);
  const pageCount = Math.ceil(transactions.length / txsPerPage);

  return (
    <Box>
      <Box sx={{ mb: 3, display: 'flex', gap: 2 }}>
        <TextField
          label="Search Address"
          variant="outlined"
          size="small"
          value={address}
          onChange={(e) => setAddress(e.target.value)}
          sx={{ width: 400 }}
        />
        <Button
          variant="contained"
          startIcon={<SearchIcon />}
          onClick={handleSearch}
          disabled={loading || !address}
        >
          Search
        </Button>
      </Box>
      {loading && (
        <Box sx={{ display: 'flex', justifyContent: 'center', p: 3 }}>
          <CircularProgress />
        </Box>
      )}
      {error && <Alert severity="error">{error}</Alert>}
      
      {(totals || ethBalance || evndBalance) && (
        <Grid container spacing={3} sx={{ mb: 3 }}>
          {totals && (
            <Grid item xs={12} md={4}>
              <Card>
                <CardContent>
                  <Typography variant="h6">Total Amount</Typography>
                  <Stack direction="row" spacing={4} sx={{ mt: 1 }}>
                    <Typography color="primary">In: <b>{(totals.in / 1e18).toFixed(4)} eVND</b></Typography>
                    <Typography color="secondary">Out: <b>{(totals.out / 1e18).toFixed(4)} eVND</b></Typography>
                  </Stack>
                </CardContent>
              </Card>
            </Grid>
          )}
          
          {ethBalance && (
            <Grid item xs={12} md={4}>
              <Card>
                <CardContent>
                  <Typography variant="h6">ETH Balance</Typography>
                  <Typography variant="h4" sx={{ mt: 1 }}>
                    {ethBalance.balance.toFixed(4)} {ethBalance.unit}
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          )}
          
          {evndBalance && (
            <Grid item xs={12} md={4}>
              <Card>
                <CardContent>
                  <Typography variant="h6">eVND Balance</Typography>
                  <Typography variant="h4" sx={{ mt: 1 }}>
                    {evndBalance.balance.toFixed(4)} {evndBalance.unit}
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          )}
        </Grid>
      )}

      {transactions.length > 0 && (
        <Box>
          <Typography variant="h6" sx={{ mb: 1 }}>Related Transactions</Typography>
          <TableContainer component={Paper}>
            <Table size="small">
              <TableHead>
                <TableRow>
                  <TableCell>Hash</TableCell>
                  <TableCell>From</TableCell>
                  <TableCell>To</TableCell>
                  <TableCell>Amount</TableCell>
                  <TableCell>Block Number</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {paginatedTxs.map((tx, idx) => (
                  <TableRow key={idx}>
                    <TableCell>{tx.txHash ? `${tx.txHash.slice(0, 8)}...${tx.txHash.slice(-8)}` : ''}</TableCell>
                    <TableCell>{tx.from_address ? `${tx.from_address.slice(0, 8)}...${tx.from_address.slice(-8)}` : ''}</TableCell>
                    <TableCell>{tx.to_address ? `${tx.to_address.slice(0, 8)}...${tx.to_address.slice(-8)}` : ''}</TableCell>
                    <TableCell>{typeof tx.amount !== 'undefined' && tx.amount !== null && !isNaN(Number(tx.amount)) ? `${(Number(tx.amount) / 1e18).toFixed(4)} eVND` : ''}</TableCell>
                    <TableCell>{tx.blockNumber ?? ''}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
          {pageCount > 1 && (
            <Box sx={{ display: 'flex', justifyContent: 'center', mt: 2 }}>
              <Pagination
                count={pageCount}
                page={page}
                onChange={(_, value) => setPage(value)}
                color="primary"
              />
            </Box>
          )}
        </Box>
      )}
      {address && !isBlacklisted && (
        <Button
          variant="contained"
          color="error"
          sx={{ mb: 2 }}
          onClick={handleBlacklist}
        >
          Blacklist Address
        </Button>
      )}
      <Snackbar
        open={snackbar.open}
        autoHideDuration={3000}
        onClose={() => setSnackbar({ ...snackbar, open: false })}
        anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
      >
        <MuiAlert onClose={() => setSnackbar({ ...snackbar, open: false })} severity={snackbar.severity} sx={{ width: '100%' }}>
          {snackbar.message}
        </MuiAlert>
      </Snackbar>
    </Box>
  );
}; 