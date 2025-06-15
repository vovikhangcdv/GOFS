import { useEffect, useState } from 'react';
import {
  Card,
  CardContent,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Box,
  CircularProgress,
  Alert,
  Chip,
  Paper,
  styled,
} from '@mui/material';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import CancelIcon from '@mui/icons-material/Cancel';
import axios from 'axios';

interface Rule {
  id: number;
  name: string;
  status: string;
  violations: number;
}

const StyledCard = styled(Card)<{ clickable?: boolean }>(({ theme, clickable }) => ({
  background: 'rgba(30, 41, 59, 0.8)',
  backdropFilter: 'blur(20px)',
  border: '1px solid rgba(100, 116, 139, 0.3)',
  borderRadius: 16,
  overflow: 'hidden',
  height: '100%',
  ...(clickable && {
    cursor: 'pointer',
    transition: 'all 0.2s ease-in-out',
    '&:hover': {
      transform: 'translateY(-2px)',
      boxShadow: '0 8px 32px rgba(59, 130, 246, 0.3)',
      border: '1px solid rgba(59, 130, 246, 0.5)',
    },
  }),
}));

const StyledTableContainer = styled(TableContainer)(({ theme }) => ({
  background: 'transparent',
  boxShadow: 'none',
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
        padding: theme.spacing(2),
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
          fontSize: '0.9rem',
          padding: theme.spacing(2),
        },
      },
    },
  },
}));

interface ComplianceRulesTableProps {
  onClick?: () => void;
}

export const ComplianceRulesTable = ({ onClick }: ComplianceRulesTableProps) => {
  const [rules, setRules] = useState<Rule[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchRules = async () => {
      try {
        setLoading(true);
        const res = await axios.get('/api/rules/violations');
        setRules(res.data);
      } catch (err) {
        setError('Failed to fetch compliance rules');
      } finally {
        setLoading(false);
      }
    };
    fetchRules();
  }, []);

  return (
    <StyledCard clickable={!!onClick} onClick={onClick}>
      <CardContent sx={{ p: { xs: 2, md: 3 } }}>
        <Box sx={{ mb: 2 }}>
          <Typography 
            variant="h6" 
            sx={{ 
              fontWeight: 600, 
              color: 'text.primary',
              fontSize: '1.1rem',
              display: 'flex',
              alignItems: 'center',
              gap: 1,
              justifyContent: 'space-between',
            }}
          >
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              📋 ACTIVE COMPLIANCE RULES
            </Box>
            {onClick && (
              <Typography 
                variant="caption" 
                sx={{ 
                  color: 'primary.main',
                  fontSize: '0.7rem',
                  fontWeight: 600,
                  opacity: 0.8,
                }}
              >
                Click to manage →
              </Typography>
            )}
          </Typography>
        </Box>
        
        {loading ? (
          <Box sx={{ display: 'flex', justifyContent: 'center', p: 3 }}>
            <CircularProgress 
              size={40}
              sx={{
                color: 'primary.main',
                '& .MuiCircularProgress-circle': {
                  strokeLinecap: 'round',
                },
              }}
            />
          </Box>
        ) : error ? (
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
        ) : (
          <StyledTableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Rule ID</TableCell>
                  <TableCell>Name</TableCell>
                  <TableCell>Status</TableCell>
                  <TableCell align="right">Violations/24h</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {rules.map((rule) => (
                  <TableRow key={rule.id}>
                    <TableCell sx={{ fontFamily: 'monospace', color: 'text.primary' }}>
                      {`R${String(rule.id).padStart(3, '0')}`}
                    </TableCell>
                    <TableCell sx={{ color: 'text.primary', textTransform: 'uppercase', fontWeight: 600 }}>
                      {rule.name}
                    </TableCell>
                    <TableCell>
                      {rule.status === 'active' ? (
                        <Chip
                          icon={<CheckCircleIcon sx={{ fontSize: '16px !important' }} />}
                          label="Active"
                          size="small"
                          sx={{
                            background: 'rgba(34, 197, 94, 0.2)',
                            color: '#22c55e',
                            border: '1px solid rgba(34, 197, 94, 0.3)',
                            '& .MuiChip-icon': {
                              color: '#22c55e',
                            },
                          }}
                        />
                      ) : (
                        <Chip
                          icon={<CancelIcon sx={{ fontSize: '16px !important' }} />}
                          label="Inactive"
                          size="small"
                          sx={{
                            background: 'rgba(156, 163, 175, 0.2)',
                            color: '#9ca3af',
                            border: '1px solid rgba(156, 163, 175, 0.3)',
                            '& .MuiChip-icon': {
                              color: '#9ca3af',
                            },
                          }}
                        />
                      )}
                    </TableCell>
                    <TableCell align="right" sx={{ color: 'text.primary', fontWeight: 600 }}>
                      {rule.violations}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </StyledTableContainer>
        )}
      </CardContent>
    </StyledCard>
  );
}; 