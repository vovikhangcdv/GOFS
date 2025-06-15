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
  styled,
} from '@mui/material';
import {
  TrendingUp as TrendingUpIcon,
  TrendingDown as TrendingDownIcon,
  Warning as WarningIcon,
  Delete as DeleteIcon,
  Security as SecurityIcon,
  Assessment as AssessmentIcon,
} from '@mui/icons-material';
import { getSuspiciousTransactions, getBlacklist } from '../../api';
import { SuspiciousTransfer, BlacklistedAddress, SuspiciousAddress, WhitelistAddress } from '../../types';
import { ComplianceRulesTable } from './ComplianceRulesTable';
import { TransactionChart } from './TransactionChart';
import axios from 'axios';

interface DashboardStatsProps {
  onPageChange?: (page: string) => void;
}

const StatsCard = styled(Card)(({ theme }) => ({
  background: 'rgba(30, 41, 59, 0.8)',
  backdropFilter: 'blur(20px)',
  border: '1px solid rgba(100, 116, 139, 0.3)',
  borderRadius: 16,
  overflow: 'hidden',
  position: 'relative',
  transition: 'all 0.2s ease-in-out',
  cursor: 'pointer',
  '&:hover': {
    transform: 'translateY(-2px)',
    border: '1px solid rgba(59, 130, 246, 0.5)',
    boxShadow: '0 8px 32px rgba(59, 130, 246, 0.3)',
  },
  '&::before': {
    content: '""',
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    height: 3,
    background: 'linear-gradient(90deg, #64748b, #6b7280, #9ca3af)',
  },
}));

const IconContainer = styled(Box)(({ theme }) => ({
  width: 48,
  height: 48,
  borderRadius: 12,
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  marginBottom: theme.spacing(2),
  background: 'rgba(100, 116, 139, 0.15)',
  border: '1px solid rgba(100, 116, 139, 0.3)',
}));

const StatsValue = styled(Typography)(({ theme }) => ({
  fontSize: '2.5rem',
  fontWeight: 700,
  background: 'linear-gradient(135deg, #f8fafc, #e2e8f0)',
  WebkitBackgroundClip: 'text',
  WebkitTextFillColor: 'transparent',
  marginBottom: theme.spacing(1),
}));

const TableCard = styled(Card)(({ theme }) => ({
  background: 'rgba(30, 41, 59, 0.8)',
  backdropFilter: 'blur(20px)',
  border: '1px solid rgba(100, 116, 139, 0.3)',
  borderRadius: 16,
  overflow: 'hidden',
  '& .MuiTableContainer-root': {
    background: 'transparent',
    border: 'none',
  },
  '& .MuiTable-root': {
    '& .MuiTableHead-root': {
      background: 'rgba(100, 116, 139, 0.2)',
      '& .MuiTableCell-head': {
        color: theme.palette.text.primary,
        fontWeight: 600,
        fontSize: '0.85rem',
        textTransform: 'uppercase',
        letterSpacing: '0.05em',
        border: 'none',
      },
    },
    '& .MuiTableBody-root': {
      '& .MuiTableRow-root': {
        '&:hover': {
          background: 'rgba(100, 116, 139, 0.1)',
        },
        '& .MuiTableCell-body': {
          border: 'none',
          borderBottom: '1px solid rgba(100, 116, 139, 0.2)',
          color: theme.palette.text.secondary,
          fontSize: '0.8rem',
        },
      },
    },
  },
}));

export const DashboardStats = ({ onPageChange }: DashboardStatsProps) => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [suspiciousTxs, setSuspiciousTxs] = useState<SuspiciousTransfer[]>([]);
  const [blacklist, setBlacklist] = useState<BlacklistedAddress[]>([]);
  const [whitelist, setWhitelist] = useState<WhitelistAddress[]>([]);
  const [suspiciousAddresses, setSuspiciousAddresses] = useState<SuspiciousAddress[]>([]);

  const fetchData = async () => {
    try {
      setLoading(true);
      setError(null);
      const [suspiciousData, blacklistData, whitelistData, suspiciousAddrData] = await Promise.all([
        getSuspiciousTransactions(),
        getBlacklist(),
        axios.get('/api/whitelist-addresses').then(res => res.data),
        axios.get('/api/suspicious-addresses').then(res => res.data),
      ]);
      
      // Ensure we have valid arrays with proper structure
      setSuspiciousTxs(Array.isArray(suspiciousData) ? suspiciousData : []);
      setBlacklist(Array.isArray(blacklistData) ? blacklistData : []);
      setWhitelist(Array.isArray(whitelistData) ? whitelistData : []);
      setSuspiciousAddresses(Array.isArray(suspiciousAddrData) ? suspiciousAddrData : []);
    } catch (err) {
      setError('Failed to fetch dashboard data');
      console.error('Dashboard data fetch error:', err);
      // Set empty arrays as fallbacks
      setSuspiciousTxs([]);
      setBlacklist([]);
      setWhitelist([]);
      setSuspiciousAddresses([]);
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

  const handleCardClick = (page: string) => {
    if (onPageChange) {
      onPageChange(page);
    }
  };

  if (loading) {
    return (
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
          Loading dashboard data...
        </Typography>
      </Box>
    );
  }

  if (error) {
    return (
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
    );
  }

  const totalSuspiciousAmount = suspiciousTxs.reduce(
    (sum, tx) => {
      const amount = Number(tx.amount);
      return sum + (isNaN(amount) ? 0 : amount);
    },
    0
  );

  const highSeverityCount = suspiciousTxs.filter(tx => tx.severity === 'high').length;

  return (
    <Box sx={{ 
      p: { xs: 0.5, md: 1 }, 
      maxWidth: '100%', 
      height: '100%',
      width: '100%',
    }}>
      {/* Stats Overview */}
      <Box sx={{ mb: 3 }}>
        <Typography 
          variant="h4" 
          sx={{ 
            mb: 2,
            fontWeight: 700,
            color: 'text.primary',
            fontSize: { xs: '1.75rem', md: '2.125rem' },
          }}
        >
          FDS Dashboard
        </Typography>
        
        <Grid container spacing={3}>
          <Grid item xs={12} sm={6} md={4}>
            <StatsCard onClick={() => handleCardClick('transactions')}>
              <CardContent sx={{ p: 2 }}>
                <IconContainer>
                  <WarningIcon sx={{ color: 'warning.light', fontSize: 24 }} />
                </IconContainer>
                <Typography variant="h6" sx={{ fontWeight: 600, mb: 1, color: 'text.primary', fontSize: '1rem' }}>
                  Suspicious Transactions
                </Typography>
                <StatsValue sx={{ fontSize: '2rem' }}>{suspiciousTxs.length}</StatsValue>
                <Typography variant="body2" color="text.secondary" sx={{ mb: 1, fontSize: '0.8rem' }}>
                  Total Amount: {totalSuspiciousAmount > 0 ? (totalSuspiciousAmount / 1e18).toFixed(2) : 'NaN'} eVND
                </Typography>
                <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                  <Box sx={{ display: 'flex', alignItems: 'center' }}>
                    <TrendingUpIcon sx={{ color: 'warning.light', fontSize: 14, mr: 0.5 }} />
                    <Typography variant="caption" sx={{ color: 'warning.light', fontSize: '0.7rem' }}>
                      +12% from last week
                    </Typography>
                  </Box>
                  <Typography 
                    variant="caption" 
                    sx={{ 
                      color: 'primary.main',
                      fontSize: '0.65rem',
                      fontWeight: 600,
                      opacity: 0.8,
                    }}
                  >
                    Click to view →
                  </Typography>
                </Box>
              </CardContent>
            </StatsCard>
          </Grid>

          <Grid item xs={12} sm={6} md={4}>
            <StatsCard onClick={() => handleCardClick('transactions')}>
              <CardContent sx={{ p: 2 }}>
                <IconContainer>
                  <AssessmentIcon sx={{ color: 'error.light', fontSize: 24 }} />
                </IconContainer>
                <Typography variant="h6" sx={{ fontWeight: 600, mb: 1, color: 'text.primary', fontSize: '1rem' }}>
                  High Severity
                </Typography>
                <StatsValue sx={{ fontSize: '2rem' }}>{highSeverityCount}</StatsValue>
                <Typography variant="body2" color="text.secondary" sx={{ mb: 1, fontSize: '0.8rem' }}>
                  Transactions requiring immediate attention
                </Typography>
                <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                  <Box sx={{ display: 'flex', alignItems: 'center' }}>
                    <TrendingDownIcon sx={{ color: 'error.light', fontSize: 14, mr: 0.5 }} />
                    <Typography variant="caption" sx={{ color: 'error.light', fontSize: '0.7rem' }}>
                      Critical alerts
                    </Typography>
                  </Box>
                  <Typography 
                    variant="caption" 
                    sx={{ 
                      color: 'primary.main',
                      fontSize: '0.65rem',
                      fontWeight: 600,
                      opacity: 0.8,
                    }}
                  >
                    Click to view →
                  </Typography>
                </Box>
              </CardContent>
            </StatsCard>
          </Grid>

          <Grid item xs={12} sm={6} md={4}>
            <StatsCard onClick={() => handleCardClick('blacklist')}>
              <CardContent sx={{ p: 2 }}>
                <IconContainer>
                  <SecurityIcon sx={{ color: 'error.light', fontSize: 24 }} />
                </IconContainer>
                <Typography variant="h6" sx={{ fontWeight: 600, mb: 1, color: 'text.primary', fontSize: '1rem' }}>
                  Blacklisted Addresses
                </Typography>
                <StatsValue sx={{ fontSize: '2rem' }}>{blacklist.length}</StatsValue>
                <Typography variant="body2" color="text.secondary" sx={{ mb: 1, fontSize: '0.8rem' }}>
                  Addresses currently blocked
                </Typography>
                <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                  <Box sx={{ display: 'flex', alignItems: 'center' }}>
                    <TrendingUpIcon sx={{ color: 'warning.light', fontSize: 14, mr: 0.5 }} />
                    <Typography variant="caption" sx={{ color: 'warning.light', fontSize: '0.7rem' }}>
                      Active protection
                    </Typography>
                  </Box>
                  <Typography 
                    variant="caption" 
                    sx={{ 
                      color: 'primary.main',
                      fontSize: '0.65rem',
                      fontWeight: 600,
                      opacity: 0.8,
                    }}
                  >
                    Click to view →
                  </Typography>
                </Box>
              </CardContent>
            </StatsCard>
          </Grid>
        </Grid>
      </Box>

      {/* Compliance Rules and Transaction Analytics */}
      <Box sx={{ mb: 4 }}>
        <Grid container spacing={3}>
          <Grid item xs={12} lg={6}>
            <ComplianceRulesTable onClick={() => handleCardClick('rules')} />
          </Grid>
          <Grid item xs={12} lg={6}>
            <TransactionChart />
          </Grid>
        </Grid>
      </Box>

      {/* Address Management */}
      <Grid container spacing={3}>
        <Grid item xs={12} lg={6}>
          <TableCard sx={{ height: 'fit-content' }}>
            <CardContent sx={{ p: 0 }}>
              <Box sx={{ p: 2, pb: 1 }}>
                <Typography variant="h6" sx={{ fontWeight: 600, fontSize: '1.1rem' }}>
                  Whitelist Addresses
                </Typography>
                <Typography variant="body2" color="text.secondary" sx={{ mt: 0.5, fontSize: '0.85rem' }}>
                  Trusted addresses with verified status
                </Typography>
              </Box>
              <Box sx={{ maxHeight: 250, overflow: 'auto', px: 2, pb: 2 }}>
                {whitelist.length > 0 ? (
                  <Box component="table" sx={{ width: '100%', borderCollapse: 'collapse' }}>
                    <Box component="thead">
                      <Box component="tr">
                        <Box component="th" sx={{ textAlign: 'left', py: 1, fontWeight: 600, fontSize: '0.7rem', textTransform: 'uppercase', color: 'text.secondary' }}>Address</Box>
                        <Box component="th" sx={{ textAlign: 'left', py: 1, fontWeight: 600, fontSize: '0.7rem', textTransform: 'uppercase', color: 'text.secondary' }}>Reason</Box>
                        <Box component="th" sx={{ textAlign: 'center', py: 1, fontWeight: 600, fontSize: '0.7rem', textTransform: 'uppercase', color: 'text.secondary' }}>Actions</Box>
                      </Box>
                    </Box>
                    <Box component="tbody">
                      {whitelist.filter(addr => addr && addr.address).map(addr => (
                        <Box component="tr" key={addr.id} sx={{ '&:hover': { background: 'rgba(100, 116, 139, 0.1)' } }}>
                          <Box component="td" sx={{ py: 1.5, fontFamily: 'monospace', fontSize: '0.75rem', color: 'text.primary' }}>
                            {addr.address && addr.address.length > 10 
                              ? `${addr.address.slice(0, 6)}...${addr.address.slice(-4)}`
                              : addr.address || 'N/A'
                            }
                          </Box>
                          <Box component="td" sx={{ py: 1.5, fontSize: '0.75rem', color: 'text.secondary' }}>{addr.reason || 'N/A'}</Box>
                          <Box component="td" sx={{ py: 1.5, textAlign: 'center' }}>
                            <Tooltip title="Remove from whitelist">
                              <IconButton
                                size="small"
                                sx={{ 
                                  color: 'error.light',
                                  '&:hover': { background: 'rgba(239, 68, 68, 0.1)' }
                                }}
                                onClick={() => addr.address && handleRemoveWhitelist(addr.address)}
                              >
                                <DeleteIcon fontSize="small" />
                              </IconButton>
                            </Tooltip>
                          </Box>
                        </Box>
                      ))}
                    </Box>
                  </Box>
                ) : (
                  <Box sx={{ textAlign: 'center', py: 3 }}>
                    <Typography variant="body2" color="text.secondary" sx={{ fontSize: '0.85rem' }}>
                      No whitelist addresses found
                    </Typography>
                  </Box>
                )}
              </Box>
            </CardContent>
          </TableCard>
        </Grid>

        <Grid item xs={12} lg={6}>
          <TableCard sx={{ height: 'fit-content' }}>
            <CardContent sx={{ p: 0 }}>
              <Box sx={{ p: 2, pb: 1 }}>
                <Typography variant="h6" sx={{ fontWeight: 600, fontSize: '1.1rem' }}>
                  Suspicious Addresses
                </Typography>
                <Typography variant="body2" color="text.secondary" sx={{ mt: 0.5, fontSize: '0.85rem' }}>
                  Addresses flagged for suspicious activity
                </Typography>
              </Box>
              <Box sx={{ maxHeight: 250, overflow: 'auto', px: 2, pb: 2 }}>
                {suspiciousAddresses.length > 0 ? (
                  <Box component="table" sx={{ width: '100%', borderCollapse: 'collapse' }}>
                    <Box component="thead">
                      <Box component="tr">
                        <Box component="th" sx={{ textAlign: 'left', py: 1, fontWeight: 600, fontSize: '0.7rem', textTransform: 'uppercase', color: 'text.secondary' }}>Address</Box>
                        <Box component="th" sx={{ textAlign: 'left', py: 1, fontWeight: 600, fontSize: '0.7rem', textTransform: 'uppercase', color: 'text.secondary' }}>Reason</Box>
                        <Box component="th" sx={{ textAlign: 'center', py: 1, fontWeight: 600, fontSize: '0.7rem', textTransform: 'uppercase', color: 'text.secondary' }}>Actions</Box>
                      </Box>
                    </Box>
                    <Box component="tbody">
                      {suspiciousAddresses.filter(addr => addr && addr.address).map(addr => (
                        <Box component="tr" key={addr.id} sx={{ '&:hover': { background: 'rgba(100, 116, 139, 0.1)' } }}>
                          <Box component="td" sx={{ py: 1.5, fontFamily: 'monospace', fontSize: '0.75rem', color: 'text.primary' }}>
                            {addr.address && addr.address.length > 10 
                              ? `${addr.address.slice(0, 6)}...${addr.address.slice(-4)}`
                              : addr.address || 'N/A'
                            }
                          </Box>
                          <Box component="td" sx={{ py: 1.5, fontSize: '0.75rem', color: 'text.secondary' }}>{addr.reason || 'N/A'}</Box>
                          <Box component="td" sx={{ py: 1.5, textAlign: 'center' }}>
                            <Tooltip title="Remove from suspicious list">
                              <IconButton
                                size="small"
                                sx={{ 
                                  color: 'error.light',
                                  '&:hover': { background: 'rgba(239, 68, 68, 0.1)' }
                                }}
                                onClick={() => addr.address && handleRemoveSuspicious(addr.address)}
                              >
                                <DeleteIcon fontSize="small" />
                              </IconButton>
                            </Tooltip>
                          </Box>
                        </Box>
                      ))}
                    </Box>
                  </Box>
                ) : (
                  <Box sx={{ textAlign: 'center', py: 3 }}>
                    <Typography variant="body2" color="text.secondary" sx={{ fontSize: '0.85rem' }}>
                      No suspicious addresses found
                    </Typography>
                  </Box>
                )}
              </Box>
            </CardContent>
          </TableCard>
        </Grid>
      </Grid>
    </Box>
  );
}; 