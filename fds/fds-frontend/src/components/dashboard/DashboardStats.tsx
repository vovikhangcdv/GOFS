import { useState, useEffect } from 'react';
import {
  Box,
  Card,
  CardContent,
  Grid,
  Typography,
  CircularProgress,
  Alert,
  IconButton,
  Tooltip,
} from '@mui/material';
import {
  TrendingUp as TrendingUpIcon,
  TrendingDown as TrendingDownIcon,
  Warning as WarningIcon,
  Delete as DeleteIcon,
} from '@mui/icons-material';
import { getSuspiciousTransactions, getBlacklist } from '../../api';
import { SuspiciousTransfer, BlacklistedAddress, SuspiciousAddress, WhitelistAddress } from '../../types';
import { ComplianceRulesTable } from './ComplianceRulesTable';
import axios from 'axios';

export const DashboardStats = () => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [suspiciousTxs, setSuspiciousTxs] = useState<SuspiciousTransfer[]>([]);
  const [blacklist, setBlacklist] = useState<BlacklistedAddress[]>([]);
  const [whitelist, setWhitelist] = useState<WhitelistAddress[]>([]);
  const [suspiciousAddresses, setSuspiciousAddresses] = useState<SuspiciousAddress[]>([]);

  const fetchData = async () => {
    try {
      setLoading(true);
      const [suspiciousData, blacklistData, whitelistData, suspiciousAddrData] = await Promise.all([
        getSuspiciousTransactions(),
        getBlacklist(),
        axios.get('/api/whitelist-addresses').then(res => res.data),
        axios.get('/api/suspicious-addresses').then(res => res.data),
      ]);
      setSuspiciousTxs(suspiciousData);
      setBlacklist(blacklistData);
      setWhitelist(whitelistData);
      setSuspiciousAddresses(suspiciousAddrData);
    } catch (err) {
      setError('Failed to fetch dashboard data');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  const handleRemoveWhitelist = async (address: string) => {
    try {
      await axios.post('/api/whitelist-addresses/remove', { address });
      await fetchData(); // Refresh data after removal
    } catch (err) {
      console.error('Failed to remove whitelist address:', err);
      setError('Failed to remove whitelist address');
    }
  };

  const handleRemoveSuspicious = async (address: string) => {
    try {
      await axios.post('/api/suspicious-addresses/remove', { address });
      await fetchData(); // Refresh data after removal
    } catch (err) {
      console.error('Failed to remove suspicious address:', err);
      setError('Failed to remove suspicious address');
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

  const totalSuspiciousAmount = suspiciousTxs.reduce(
    (sum, tx) => sum + Number(tx.amount),
    0
  );

  return (
    <>
      <Grid container spacing={3}>
        <Grid item xs={12} md={4}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                <WarningIcon color="warning" sx={{ mr: 1 }} />
                <Typography variant="h6">Suspicious Transactions</Typography>
              </Box>
              <Typography variant="h4" gutterBottom>
                {suspiciousTxs.length}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Total Amount: {(totalSuspiciousAmount / 1e18).toFixed(4)} ETH
              </Typography>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={4}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                <TrendingUpIcon color="success" sx={{ mr: 1 }} />
                <Typography variant="h6">High Severity</Typography>
              </Box>
              <Typography variant="h4" gutterBottom>
                {suspiciousTxs.filter(tx => tx.severity === 'high').length}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Transactions requiring immediate attention
              </Typography>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={4}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                <TrendingDownIcon color="error" sx={{ mr: 1 }} />
                <Typography variant="h6">Blacklisted Addresses</Typography>
              </Box>
              <Typography variant="h4" gutterBottom>
                {blacklist.length}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Addresses currently blocked
              </Typography>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
      <ComplianceRulesTable />
      {/* Whitelist & Suspicious Addresses Row */}
      <Grid container spacing={3} sx={{ mt: 1 }}>
        <Grid item xs={12} md={6}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>Whitelist Addresses</Typography>
              <Box sx={{ maxHeight: 250, overflow: 'auto' }}>
                <table style={{ width: '100%' }}>
                  <thead>
                    <tr>
                      <th style={{ textAlign: 'left' }}>Address</th>
                      <th style={{ textAlign: 'left' }}>Reason</th>
                      <th style={{ textAlign: 'center' }}>Actions</th>
                    </tr>
                  </thead>
                  <tbody>
                    {whitelist.map(addr => (
                      <tr key={addr.id}>
                        <td style={{ fontFamily: 'monospace' }}>{addr.address}</td>
                        <td>{addr.reason}</td>
                        <td style={{ textAlign: 'center' }}>
                          <Tooltip title="Remove from whitelist">
                            <IconButton
                              size="small"
                              color="error"
                              onClick={() => handleRemoveWhitelist(addr.address)}
                            >
                              <DeleteIcon fontSize="small" />
                            </IconButton>
                          </Tooltip>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </Box>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={12} md={6}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>Suspicious Addresses</Typography>
              <Box sx={{ maxHeight: 250, overflow: 'auto' }}>
                <table style={{ width: '100%' }}>
                  <thead>
                    <tr>
                      <th style={{ textAlign: 'left' }}>Address</th>
                      <th style={{ textAlign: 'left' }}>Reason</th>
                      <th style={{ textAlign: 'center' }}>Actions</th>
                    </tr>
                  </thead>
                  <tbody>
                    {suspiciousAddresses.map(addr => (
                      <tr key={addr.id}>
                        <td style={{ fontFamily: 'monospace' }}>{addr.address}</td>
                        <td>{addr.reason}</td>
                        <td style={{ textAlign: 'center' }}>
                          <Tooltip title="Remove from suspicious list">
                            <IconButton
                              size="small"
                              color="error"
                              onClick={() => handleRemoveSuspicious(addr.address)}
                            >
                              <DeleteIcon fontSize="small" />
                            </IconButton>
                          </Tooltip>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </>
  );
}; 