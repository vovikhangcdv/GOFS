import { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Button,
  CircularProgress,
  Alert,
  Snackbar,
  Alert as MuiAlert,
} from '@mui/material';
import { getBlacklist, unblacklistAddress, deleteBlacklistAddress } from '../../api';
import { BlacklistedAddress } from '../../types';

export const BlacklistView = () => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [blacklist, setBlacklist] = useState<BlacklistedAddress[]>([]);
  const [snackbar, setSnackbar] = useState<{ open: boolean; message: string; severity: 'success' | 'error' }>({ open: false, message: '', severity: 'success' });

  useEffect(() => {
    fetchBlacklist();
  }, []);

  const fetchBlacklist = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await getBlacklist();
      setBlacklist(data);
    } catch (err) {
      setError('Failed to fetch blacklist');
    } finally {
      setLoading(false);
    }
  };

  const handleUnblacklist = async (address: string) => {
    try {
      // First call unblacklist endpoint
      await unblacklistAddress(address);
      // Then call delete endpoint
      await deleteBlacklistAddress(address);
      setSnackbar({ open: true, message: 'Address unblacklisted successfully', severity: 'success' });
      // Refresh the blacklist after successful unblacklist
      await fetchBlacklist();
    } catch (err) {
      setSnackbar({ open: true, message: 'Failed to unblacklist address', severity: 'error' });
    }
  };

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
      <Typography variant="h6" sx={{ mb: 2 }}>
        Blacklisted Addresses
      </Typography>
      <TableContainer component={Paper}>
        <Table size="small">
          <TableHead>
            <TableRow>
              <TableCell>Address</TableCell>
              <TableCell>Reason</TableCell>
              <TableCell>Severity</TableCell>
              <TableCell>Block Number</TableCell>
              <TableCell>Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {blacklist.length === 0 ? (
              <TableRow>
                <TableCell colSpan={5} align="center">
                  No blacklisted addresses found.
                </TableCell>
              </TableRow>
            ) : (
              blacklist.map((item) => (
                <TableRow key={item.id}>
                  <TableCell>{item.address}</TableCell>
                  <TableCell>{item.reason}</TableCell>
                  <TableCell>{item.severity}</TableCell>
                  <TableCell>{item.blockNumber}</TableCell>
                  <TableCell>
                    <Button
                      variant="outlined"
                      color="primary"
                      size="small"
                      onClick={() => handleUnblacklist(item.address)}
                    >
                      Unblacklist
                    </Button>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </TableContainer>
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