import { useState, useEffect } from 'react';
import {
  Box,
  Card,
  CardContent,
  Grid,
  Typography,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  TrendingUp as TrendingUpIcon,
  TrendingDown as TrendingDownIcon,
  Warning as WarningIcon,
} from '@mui/icons-material';
import { getSuspiciousTransactions, getBlacklist } from '../../api';
import { SuspiciousTransfer, BlacklistedAddress } from '../../types';

export const DashboardStats = () => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [suspiciousTxs, setSuspiciousTxs] = useState<SuspiciousTransfer[]>([]);
  const [blacklist, setBlacklist] = useState<BlacklistedAddress[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const [suspiciousData, blacklistData] = await Promise.all([
          getSuspiciousTransactions(),
          getBlacklist(),
        ]);
        setSuspiciousTxs(suspiciousData);
        setBlacklist(blacklistData);
      } catch (err) {
        setError('Failed to fetch dashboard data');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

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
  );
}; 