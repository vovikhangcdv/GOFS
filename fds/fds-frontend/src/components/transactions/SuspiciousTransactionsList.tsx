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
import { formatTokenAmount } from '../../utils/format';

// Helper function to safely truncate addresses
const formatAddress = (address: string | undefined | null): string => {
  if (!address || typeof address !== 'string' || address.length < 10) {
    return address || 'N/A';
  }
  return `${address.slice(0, 8)}...${address.slice(-8)}`;
};

// Helper function to safely truncate transaction hashes
const formatTxHash = (txHash: string | undefined | null): string => {
  if (!txHash || typeof txHash !== 'string' || txHash.length < 16) {
    return txHash || 'N/A';
  }
  return `${txHash.slice(0, 8)}...${txHash.slice(-8)}`;
};

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
      setError(null);
      const data = await getSuspiciousTransactions();
      // Ensure we have valid array with proper structure
      setTransactions(Array.isArray(data) ? data : []);
    } catch (err) {
      setError('Failed to fetch suspicious transactions');
      console.error('Suspicious transactions fetch error:', err);
      // Set empty array as fallback
      setTransactions([]);
    } finally {
      setLoading(false);
    }
  };

  const handleTransactionClick = async (tx: SuspiciousTransfer) => {
    if (!tx.txHash) return;
    
    setSelectedTx(tx);
    try {
      const related = await getRelatedTransactionsOfSuspicious(tx.txHash);
      setRelatedTxs(Array.isArray(related) ? related : []);
    } catch (err) {
      console.error('Failed to fetch related transactions:', err);
      setRelatedTxs([]);
    }
  };

  const filteredTransactions = transactions.filter((tx) => {
    if (!tx) return false;
    
    const searchLower = searchTerm.toLowerCase();
    const matchesSearch =
      (tx.from_address && tx.from_address.toLowerCase().includes(searchLower)) ||
      (tx.to_address && tx.to_address.toLowerCase().includes(searchLower)) ||
      (tx.txHash && tx.txHash.toLowerCase().includes(searchLower));
    const matchesSeverity = severityFilter === 'all' || tx.severity === severityFilter;
    return matchesSearch && matchesSeverity;
  });

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
          Loading suspicious transactions...
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
        Suspicious Transactions
      </Typography>
      
      <Box sx={{ mb: 3, display: 'flex', gap: 2, flexWrap: 'wrap' }}>
        <TextField
          label="Search Transactions"
          variant="outlined"
          size="small"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          sx={{ 
            width: { xs: '100%', sm: 300 },
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
        <TextField
          select
          label="Severity"
          variant="outlined"
          size="small"
          value={severityFilter}
          onChange={(e) => setSeverityFilter(e.target.value)}
          sx={{ 
            width: { xs: '100%', sm: 150 },
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
        >
          <MenuItem value="all">All</MenuItem>
          <MenuItem value="high">High</MenuItem>
          <MenuItem value="medium">Medium</MenuItem>
          <MenuItem value="low">Low</MenuItem>
        </TextField>
      </Box>

      <TableContainer 
        component={Paper}
        sx={{
          background: 'rgba(30, 41, 59, 0.8)',
          backdropFilter: 'blur(20px)',
          border: '1px solid rgba(100, 116, 139, 0.3)',
          borderRadius: 2,
        }}
      >
        <Table>
          <TableHead sx={{ background: 'rgba(100, 116, 139, 0.2)' }}>
            <TableRow>
              <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>Hash</TableCell>
              <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>From</TableCell>
              <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>To</TableCell>
              <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>Amount</TableCell>
              <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>Severity</TableCell>
              <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>Reason</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {filteredTransactions.map((tx) => (
              <>
                <TableRow
                  key={tx.id || Math.random()}
                  hover
                  onClick={() => handleTransactionClick(tx)}
                  sx={{ 
                    cursor: 'pointer',
                    background: selectedTx?.id === tx.id ? 'rgba(59, 130, 246, 0.1)' : 'transparent',
                    borderLeft: selectedTx?.id === tx.id ? '3px solid #3b82f6' : '3px solid transparent',
                    '&:hover': {
                      background: selectedTx?.id === tx.id 
                        ? 'rgba(59, 130, 246, 0.15)' 
                        : 'rgba(100, 116, 139, 0.1)',
                    },
                  }}
                >
                  <TableCell sx={{ 
                    fontFamily: 'monospace', 
                    color: 'text.secondary',
                    fontSize: '0.85rem',
                  }}>
                    {formatTxHash(tx.txHash)}
                  </TableCell>
                  <TableCell sx={{ 
                    fontFamily: 'monospace', 
                    color: 'text.secondary',
                    fontSize: '0.85rem',
                  }}>
                    {formatAddress(tx.from_address)}
                  </TableCell>
                  <TableCell sx={{ 
                    fontFamily: 'monospace', 
                    color: 'text.secondary',
                    fontSize: '0.85rem',
                  }}>
                    {formatAddress(tx.to_address)}
                  </TableCell>
                  <TableCell sx={{ color: 'text.primary' }}>
                    {formatTokenAmount(tx.amount, 18)} eVND
                  </TableCell>
                  <TableCell>
                    <Chip
                      label={tx.severity || 'Unknown'}
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
                  <TableCell sx={{ color: 'text.secondary' }}>
                    {tx.reason || 'N/A'}
                  </TableCell>
                </TableRow>
                
                {/* Transaction Details Row - Appears directly below clicked transaction */}
                {selectedTx?.id === tx.id && (
                  <TableRow>
                    <TableCell colSpan={6} sx={{ p: 0, border: 'none' }}>
                      <Box sx={{ 
                        background: 'rgba(59, 130, 246, 0.05)',
                        border: '1px solid rgba(59, 130, 246, 0.2)',
                        borderRadius: 1,
                        m: 1,
                      }}>
                        {/* Transaction Details Section */}
                        <Box sx={{ p: 3 }}>
                          <Typography variant="h6" gutterBottom sx={{ 
                            color: 'text.primary', 
                            fontWeight: 700,
                            fontSize: '1.1rem',
                            mb: 2,
                            display: 'flex',
                            alignItems: 'center',
                            gap: 1,
                          }}>
                            🔍 Transaction Details
                          </Typography>
                          
                          <Box sx={{ 
                            display: 'grid',
                            gridTemplateColumns: { xs: '1fr', md: '1fr 1fr' },
                            gap: 2,
                            mb: 2,
                          }}>
                            <Box>
                              <Typography variant="subtitle2" sx={{ 
                                color: 'text.secondary', 
                                fontSize: '0.8rem',
                                fontWeight: 600,
                                mb: 0.5,
                              }}>
                                Hash:
                              </Typography>
                              <Typography variant="body2" sx={{ 
                                color: 'text.primary',
                                fontFamily: 'monospace',
                                fontSize: '0.8rem',
                                wordBreak: 'break-all',
                              }}>
                                {selectedTx.txHash || 'N/A'}
                              </Typography>
                            </Box>
                            
                            <Box>
                              <Typography variant="subtitle2" sx={{ 
                                color: 'text.secondary', 
                                fontSize: '0.8rem',
                                fontWeight: 600,
                                mb: 0.5,
                              }}>
                                Amount:
                              </Typography>
                              <Typography variant="body2" sx={{ 
                                color: 'text.primary',
                                fontSize: '0.8rem',
                                fontWeight: 600,
                              }}>
                                {formatTokenAmount(selectedTx.amount, 18)} eVND
                              </Typography>
                            </Box>
                            
                            <Box>
                              <Typography variant="subtitle2" sx={{ 
                                color: 'text.secondary', 
                                fontSize: '0.8rem',
                                fontWeight: 600,
                                mb: 0.5,
                              }}>
                                From:
                              </Typography>
                              <Typography variant="body2" sx={{ 
                                color: 'text.primary',
                                fontFamily: 'monospace',
                                fontSize: '0.8rem',
                                wordBreak: 'break-all',
                              }}>
                                {selectedTx.from_address || 'N/A'}
                              </Typography>
                            </Box>
                            
                            <Box>
                              <Typography variant="subtitle2" sx={{ 
                                color: 'text.secondary', 
                                fontSize: '0.8rem',
                                fontWeight: 600,
                                mb: 0.5,
                              }}>
                                To:
                              </Typography>
                              <Typography variant="body2" sx={{ 
                                color: 'text.primary',
                                fontFamily: 'monospace',
                                fontSize: '0.8rem',
                                wordBreak: 'break-all',
                              }}>
                                {selectedTx.to_address || 'N/A'}
                              </Typography>
                            </Box>
                          </Box>
                          
                          <Box sx={{ mb: 2 }}>
                            <Typography variant="subtitle2" sx={{ 
                              color: 'text.secondary', 
                              fontSize: '0.8rem',
                              fontWeight: 600,
                              mb: 0.5,
                            }}>
                              Reason:
                            </Typography>
                            <Typography variant="body2" sx={{ 
                              color: 'text.primary',
                              fontSize: '0.8rem',
                            }}>
                              {selectedTx.reason || 'N/A'}
                            </Typography>
                          </Box>
                          
                          <Box sx={{ mb: 2 }}>
                            <Typography variant="subtitle2" sx={{ 
                              color: 'text.secondary', 
                              fontSize: '0.8rem',
                              fontWeight: 600,
                              mb: 0.5,
                            }}>
                              Details:
                            </Typography>
                            <Typography variant="body2" sx={{ 
                              color: 'text.primary',
                              fontSize: '0.75rem',
                              fontFamily: 'monospace',
                              background: 'rgba(100, 116, 139, 0.1)',
                              p: 1.5,
                              borderRadius: 1,
                              border: '1px solid rgba(100, 116, 139, 0.2)',
                              wordBreak: 'break-all',
                            }}>
                              {selectedTx.details || 'N/A'}
                            </Typography>
                          </Box>

                          {/* Related Transactions - Inline */}
                          {relatedTxs.length > 0 && (
                            <Box>
                              <Typography variant="subtitle2" sx={{ 
                                color: 'text.secondary', 
                                fontSize: '0.8rem',
                                fontWeight: 600,
                                mb: 1,
                              }}>
                                Related Transactions ({relatedTxs.length}):
                              </Typography>
                              <Box sx={{ 
                                maxHeight: 200, 
                                overflow: 'auto',
                                background: 'rgba(30, 41, 59, 0.3)',
                                borderRadius: 1,
                                border: '1px solid rgba(100, 116, 139, 0.2)',
                              }}>
                                <Table size="small">
                                  <TableHead sx={{ background: 'rgba(100, 116, 139, 0.1)' }}>
                                    <TableRow>
                                      <TableCell sx={{ color: 'text.primary', fontWeight: 600, fontSize: '0.7rem', p: 1 }}>Hash</TableCell>
                                      <TableCell sx={{ color: 'text.primary', fontWeight: 600, fontSize: '0.7rem', p: 1 }}>From</TableCell>
                                      <TableCell sx={{ color: 'text.primary', fontWeight: 600, fontSize: '0.7rem', p: 1 }}>To</TableCell>
                                      <TableCell sx={{ color: 'text.primary', fontWeight: 600, fontSize: '0.7rem', p: 1 }}>Amount</TableCell>
                                      <TableCell sx={{ color: 'text.primary', fontWeight: 600, fontSize: '0.7rem', p: 1 }}>Status</TableCell>
                                    </TableRow>
                                  </TableHead>
                                  <TableBody>
                                    {relatedTxs.map((relTx) => (
                                      <TableRow key={relTx.id || Math.random()}>
                                        <TableCell sx={{ 
                                          fontFamily: 'monospace', 
                                          color: 'text.secondary',
                                          fontSize: '0.7rem',
                                          p: 1,
                                        }}>
                                          {formatTxHash(relTx.txHash)}
                                        </TableCell>
                                        <TableCell sx={{ 
                                          fontFamily: 'monospace', 
                                          color: 'text.secondary',
                                          fontSize: '0.7rem',
                                          p: 1,
                                        }}>
                                          {formatAddress(relTx.from_address)}
                                        </TableCell>
                                        <TableCell sx={{ 
                                          fontFamily: 'monospace', 
                                          color: 'text.secondary',
                                          fontSize: '0.7rem',
                                          p: 1,
                                        }}>
                                          {formatAddress(relTx.to_address)}
                                        </TableCell>
                                        <TableCell sx={{ color: 'text.primary', fontSize: '0.7rem', p: 1 }}>
                                          {formatTokenAmount(relTx.amount, 18)} eVND
                                        </TableCell>
                                        <TableCell sx={{ p: 1 }}>
                                          <Chip
                                            label={relTx.status || 'confirmed'}
                                            color={relTx.status === 'completed' ? 'success' : 'warning'}
                                            size="small"
                                            sx={{ fontSize: '0.6rem', height: 20 }}
                                          />
                                        </TableCell>
                                      </TableRow>
                                    ))}
                                  </TableBody>
                                </Table>
                              </Box>
                            </Box>
                          )}
                        </Box>
                      </Box>
                    </TableCell>
                  </TableRow>
                )}
              </>
            ))}
            {filteredTransactions.length === 0 && (
              <TableRow>
                <TableCell colSpan={6} sx={{ textAlign: 'center', py: 4 }}>
                  <Typography variant="body2" color="text.secondary">
                    No suspicious transactions found
                  </Typography>
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </TableContainer>


    </Box>
  );
}; 