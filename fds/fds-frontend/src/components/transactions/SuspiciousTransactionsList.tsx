import { useState, useEffect } from 'react';
import {
  Box,
  Card,
  CardContent,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  TextField,
  MenuItem,
  CircularProgress,
  Alert,
  Chip,
} from '@mui/material';
import { getSuspiciousTransactions, getRelatedTransactionsOfSuspicious } from '../../api';
import { SuspiciousTransfer } from '../../types';

export const SuspiciousTransactionsList = () => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [transactions, setTransactions] = useState<SuspiciousTransfer[]>([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [severityFilter, setSeverityFilter] = useState('all');
  const [selectedTx, setSelectedTx] = useState<SuspiciousTransfer | null>(null);
  const [relatedTxs, setRelatedTxs] = useState<any[]>([]);

  useEffect(() => {
    fetchTransactions();
  }, []);

  const fetchTransactions = async () => {
    try {
      setLoading(true);
      const data = await getSuspiciousTransactions();
      setTransactions(data);
    } catch (err) {
      setError('Failed to fetch suspicious transactions');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleTransactionClick = async (tx: SuspiciousTransfer) => {
    setSelectedTx(tx);
    try {
      const related = await getRelatedTransactionsOfSuspicious(tx.txHash);
      setRelatedTxs(related);
    } catch (err) {
      console.error('Failed to fetch related transactions:', err);
    }
  };

  const filteredTransactions = transactions.filter((tx) => {
    const matchesSearch =
      tx.from_address.toLowerCase().includes(searchTerm.toLowerCase()) ||
      tx.to_address.toLowerCase().includes(searchTerm.toLowerCase()) ||
      tx.txHash.toLowerCase().includes(searchTerm.toLowerCase());
    const matchesSeverity = severityFilter === 'all' || tx.severity === severityFilter;
    return matchesSearch && matchesSeverity;
  });

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 3 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return <Alert severity="error">{error}</Alert>;
  }

  return (
    <Box>
      <Box sx={{ mb: 3, display: 'flex', gap: 2 }}>
        <TextField
          label="Search Transactions"
          variant="outlined"
          size="small"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          sx={{ width: 300 }}
        />
        <TextField
          select
          label="Severity"
          variant="outlined"
          size="small"
          value={severityFilter}
          onChange={(e) => setSeverityFilter(e.target.value)}
          sx={{ width: 150 }}
        >
          <MenuItem value="all">All</MenuItem>
          <MenuItem value="high">High</MenuItem>
          <MenuItem value="medium">Medium</MenuItem>
          <MenuItem value="low">Low</MenuItem>
        </TextField>
      </Box>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Hash</TableCell>
              <TableCell>From</TableCell>
              <TableCell>To</TableCell>
              <TableCell>Amount</TableCell>
              <TableCell>Severity</TableCell>
              <TableCell>Reason</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {filteredTransactions.map((tx) => (
              <TableRow
                key={tx.id}
                hover
                onClick={() => handleTransactionClick(tx)}
                sx={{ cursor: 'pointer' }}
              >
                <TableCell>{tx.txHash.slice(0, 8)}...{tx.txHash.slice(-8)}</TableCell>
                <TableCell>{tx.from_address.slice(0, 8)}...{tx.from_address.slice(-8)}</TableCell>
                <TableCell>{tx.to_address.slice(0, 8)}...{tx.to_address.slice(-8)}</TableCell>
                <TableCell>{(Number(tx.amount) / 1e18).toFixed(4)} ETH</TableCell>
                <TableCell>
                  <Chip
                    label={tx.severity}
                    color={
                      tx.severity === 'high'
                        ? 'error'
                        : tx.severity === 'medium'
                        ? 'warning'
                        : 'success'
                    }
                    size="small"
                  />
                </TableCell>
                <TableCell>{tx.reason}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      {selectedTx && (
        <Card sx={{ mt: 3 }}>
          <CardContent>
            <Typography variant="h6" gutterBottom>
              Transaction Details
            </Typography>
            <Typography variant="body2" paragraph>
              <strong>Hash:</strong> {selectedTx.txHash}
            </Typography>
            <Typography variant="body2" paragraph>
              <strong>From:</strong> {selectedTx.from_address}
            </Typography>
            <Typography variant="body2" paragraph>
              <strong>To:</strong> {selectedTx.to_address}
            </Typography>
            <Typography variant="body2" paragraph>
              <strong>Amount:</strong> {(Number(selectedTx.amount) / 1e18).toFixed(4)} ETH
            </Typography>
            <Typography variant="body2" paragraph>
              <strong>Reason:</strong> {selectedTx.reason}
            </Typography>
            <Typography variant="body2" paragraph>
              <strong>Details:</strong> {selectedTx.details}
            </Typography>

            {relatedTxs.length > 0 && (
              <>
                <Typography variant="h6" sx={{ mt: 2, mb: 1 }}>
                  Related Transactions
                </Typography>
                <TableContainer component={Paper}>
                  <Table size="small">
                    <TableHead>
                      <TableRow>
                        <TableCell>Hash</TableCell>
                        <TableCell>From</TableCell>
                        <TableCell>To</TableCell>
                        <TableCell>Amount</TableCell>
                        <TableCell>Status</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {relatedTxs.map((tx) => (
                        <TableRow key={tx.id}>
                          <TableCell>{tx.txHash.slice(0, 8)}...{tx.txHash.slice(-8)}</TableCell>
                          <TableCell>{tx.from_address.slice(0, 8)}...{tx.from_address.slice(-8)}</TableCell>
                          <TableCell>{tx.to_address.slice(0, 8)}...{tx.to_address.slice(-8)}</TableCell>
                          <TableCell>{(Number(tx.amount) / 1e18).toFixed(4)} ETH</TableCell>
                          <TableCell>
                            <Chip
                              label={tx.status}
                              color={tx.status === 'completed' ? 'success' : 'warning'}
                              size="small"
                            />
                          </TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
              </>
            )}
          </CardContent>
        </Card>
      )}
    </Box>
  );
}; 