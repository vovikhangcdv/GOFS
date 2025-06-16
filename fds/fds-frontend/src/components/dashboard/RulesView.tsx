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
  CircularProgress,
  Alert,
  Chip,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  TextField,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Grid,
  IconButton,
  Tooltip,
  styled,
} from '@mui/material';
import {
  Edit as EditIcon,
  CheckCircle as CheckCircleIcon,
  Cancel as CancelIcon,
  Security as SecurityIcon,
} from '@mui/icons-material';
import { getRules, updateRule } from '../../api';
import { Rule } from '../../types';

const StyledCard = styled(Card)(({ theme }) => ({
  background: 'rgba(30, 41, 59, 0.8)',
  backdropFilter: 'blur(20px)',
  border: '1px solid rgba(100, 116, 139, 0.3)',
  borderRadius: 16,
  overflow: 'hidden',
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

const StyledDialog = styled(Dialog)(({ theme }) => ({
  '& .MuiDialog-paper': {
    background: 'rgba(30, 41, 59, 0.95)',
    backdropFilter: 'blur(20px)',
    border: '1px solid rgba(100, 116, 139, 0.3)',
    borderRadius: 16,
    color: theme.palette.text.primary,
    minWidth: 600,
  },
}));

export const RulesView = () => {
  const [rules, setRules] = useState<Rule[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [editingRule, setEditingRule] = useState<Rule | null>(null);
  const [dialogOpen, setDialogOpen] = useState(false);
  const [saving, setSaving] = useState(false);

  const [editForm, setEditForm] = useState({
    name: '',
    description: '',
    status: '',
    parameters: '',
  });

  const [parsedParams, setParsedParams] = useState<Record<string, any>>({});

  const fetchRules = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await getRules();
      setRules(Array.isArray(data) ? data : []);
    } catch (err) {
      setError('Failed to fetch rules');
      console.error('Rules fetch error:', err);
      setRules([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchRules();
  }, []);

  const handleEditClick = (rule: Rule) => {
    setEditingRule(rule);
    setEditForm({
      name: rule.name,
      description: rule.description,
      status: rule.status,
      parameters: rule.parameters,
    });
    
    // Parse parameters for editing
    try {
      const params = JSON.parse(rule.parameters);
      setParsedParams(params);
    } catch {
      setParsedParams({});
    }
    
    setDialogOpen(true);
  };

  const handleCloseDialog = () => {
    setDialogOpen(false);
    setEditingRule(null);
    setEditForm({ name: '', description: '', status: '', parameters: '' });
    setParsedParams({});
  };

  const handleSave = async () => {
    if (!editingRule) return;

    try {
      setSaving(true);
      setError(null);

      // Convert parsed parameters back to JSON string
      const parametersJson = JSON.stringify(parsedParams);

      await updateRule({
        name: editForm.name,
        description: editForm.description,
        status: editForm.status,
        parameters: parametersJson,
      });
      await fetchRules(); // Refresh the list
      handleCloseDialog();
    } catch (err: any) {
      setError(err.response?.data?.error || 'Failed to update rule');
      console.error('Rule update error:', err);
    } finally {
      setSaving(false);
    }
  };

  const handleParamChange = (key: string, value: string | string[]) => {
    setParsedParams(prev => ({
      ...prev,
      [key]: value
    }));
  };

  const renderParameterEditor = () => {
    if (!editingRule) return null;

    const ruleId = `R${String(editingRule.id).padStart(3, '0')}`;
    
    switch (ruleId) {
      case 'R001': // large_transfer - threshold
        return (
          <Grid item xs={12}>
            <TextField
              fullWidth
              label="Threshold (eVND)"
              type="number"
              value={parsedParams.threshold || ''}
              onChange={(e) => handleParamChange('threshold', e.target.value)}
              helperText="Minimum transfer amount to trigger alert"
              sx={{
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
          </Grid>
        );

      case 'R002': // multiple_transfers - block_range, min_transfers
        return (
          <>
            <Grid item xs={6}>
              <TextField
                fullWidth
                label="Block Range"
                type="number"
                value={parsedParams.block_range || ''}
                onChange={(e) => handleParamChange('block_range', e.target.value)}
                helperText="Number of blocks to check"
                sx={{
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
            </Grid>
            <Grid item xs={6}>
              <TextField
                fullWidth
                label="Minimum Transfers"
                type="number"
                value={parsedParams.min_transfers || ''}
                onChange={(e) => handleParamChange('min_transfers', e.target.value)}
                helperText="Minimum number of transfers to trigger"
                sx={{
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
            </Grid>
          </>
        );

      case 'R003': // multiple_incoming_transfers - threshold, block_range
        return (
          <>
            <Grid item xs={6}>
              <TextField
                fullWidth
                label="Threshold (eVND)"
                type="number"
                value={parsedParams.threshold || ''}
                onChange={(e) => handleParamChange('threshold', e.target.value)}
                helperText="Minimum total amount"
                sx={{
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
            </Grid>
            <Grid item xs={6}>
              <TextField
                fullWidth
                label="Block Range"
                type="number"
                value={parsedParams.block_range || ''}
                onChange={(e) => handleParamChange('block_range', e.target.value)}
                helperText="Number of blocks to check"
                sx={{
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
            </Grid>
          </>
        );

      case 'R004': // addresses list
        return (
          <Grid item xs={12}>
            <TextField
              fullWidth
              label="Addresses List"
              multiline
              rows={4}
              value={Array.isArray(parsedParams.addresses) ? parsedParams.addresses.join('\n') : (parsedParams.addresses || '')}
              onChange={(e) => {
                const addresses = e.target.value.split('\n').filter(addr => addr.trim() !== '');
                handleParamChange('addresses', addresses);
              }}
              helperText="Enter one address per line"
              sx={{
                '& .MuiOutlinedInput-root': {
                  background: 'rgba(30, 41, 59, 0.7)',
                  fontFamily: 'monospace',
                  '& .MuiOutlinedInput-notchedOutline': {
                    borderColor: 'rgba(100, 116, 139, 0.3)',
                  },
                  '&:hover .MuiOutlinedInput-notchedOutline': {
                    borderColor: 'rgba(100, 116, 139, 0.5)',
                  },
                },
              }}
            />
          </Grid>
        );

      case 'R005': // insufficient_balance - check_blocks
        return (
          <Grid item xs={12}>
            <TextField
              fullWidth
              label="Check Blocks"
              type="number"
              value={parsedParams.check_blocks || ''}
              onChange={(e) => handleParamChange('check_blocks', e.target.value)}
              helperText="Number of blocks to check for balance validation"
              sx={{
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
          </Grid>
        );

      default:
        return (
          <Grid item xs={12}>
            <Alert severity="info" sx={{ background: 'rgba(59, 130, 246, 0.1)', border: '1px solid rgba(59, 130, 246, 0.3)' }}>
              No editable parameters available for this rule.
            </Alert>
          </Grid>
        );
    }
  };

  const formatParameters = (params: string, ruleId: number) => {
    try {
      const parsed = JSON.parse(params);
      const ruleIdStr = `R${String(ruleId).padStart(3, '0')}`;
      
      // Show only editable parameters based on rule ID
      let editableKeys: string[] = [];
      switch (ruleIdStr) {
        case 'R001': editableKeys = ['threshold']; break;
        case 'R002': editableKeys = ['block_range', 'min_transfers']; break;
        case 'R003': editableKeys = ['threshold', 'block_range']; break;
        case 'R004': editableKeys = ['addresses']; break;
        case 'R005': editableKeys = ['check_blocks']; break;
        default: editableKeys = Object.keys(parsed); break;
      }
      
      return editableKeys
        .filter(key => parsed[key] !== undefined)
        .map(key => {
          const value = Array.isArray(parsed[key]) ? `[${parsed[key].length} items]` : parsed[key];
          return `${key}: ${value}`;
        })
        .join(', ');
    } catch {
      return params;
    }
  };

  const getSeverityColor = (severity: string) => {
    switch (severity.toLowerCase()) {
      case 'high': return '#ef4444';
      case 'medium': return '#f59e0b';
      case 'low': return '#10b981';
      default: return '#6b7280';
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
          Loading rules...
        </Typography>
      </Box>
    );
  }

  if (error && rules.length === 0) {
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
    <Box sx={{ 
      p: { xs: 1, md: 2 }, 
      maxWidth: '100%', 
      height: '100%',
      width: '100%',
    }}>
      <Box sx={{ mb: 3 }}>
        <Typography 
          variant="h4" 
          sx={{ 
            mb: 2,
            fontWeight: 700,
            color: 'text.primary',
            fontSize: { xs: '1.75rem', md: '2.125rem' },
            display: 'flex',
            alignItems: 'center',
            gap: 2,
          }}
        >
          <SecurityIcon sx={{ fontSize: '2rem', color: 'primary.main' }} />
          Compliance Rules Management
        </Typography>
        <Typography variant="body1" color="text.secondary" sx={{ mb: 2 }}>
          Configure and manage fraud detection rules and their parameters
        </Typography>
      </Box>

      {error && (
        <Alert 
          severity="error" 
          sx={{
            mb: 2,
            background: 'rgba(239, 68, 68, 0.1)',
            border: '1px solid rgba(239, 68, 68, 0.3)',
            borderRadius: 2,
          }}
          onClose={() => setError(null)}
        >
          {error}
        </Alert>
      )}

      <StyledCard>
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
              }}
            >
              🛡️ Active Compliance Rules ({rules.length})
            </Typography>
          </Box>
          
          <StyledTableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Rule ID</TableCell>
                  <TableCell>Name</TableCell>
                  <TableCell>Description</TableCell>
                  <TableCell>Status</TableCell>
                  <TableCell>Severity</TableCell>
                  <TableCell>Parameters</TableCell>
                  <TableCell>Violations</TableCell>
                  <TableCell align="center">Actions</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {rules.map((rule) => (
                  <TableRow key={rule.id}>
                    <TableCell sx={{ fontFamily: 'monospace', color: 'text.primary' }}>
                      {`R${String(rule.id).padStart(3, '0')}`}
                    </TableCell>
                    <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>
                      {rule.name}
                    </TableCell>
                    <TableCell sx={{ color: 'text.secondary', maxWidth: 300 }}>
                      {rule.description}
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
                    <TableCell>
                      <Chip
                        label={rule.severity}
                        size="small"
                        sx={{
                          background: `${getSeverityColor(rule.severity)}20`,
                          color: getSeverityColor(rule.severity),
                          border: `1px solid ${getSeverityColor(rule.severity)}40`,
                          textTransform: 'capitalize',
                        }}
                      />
                    </TableCell>
                    <TableCell sx={{ 
                      color: 'text.secondary', 
                      fontFamily: 'monospace',
                      fontSize: '0.75rem',
                      maxWidth: 200,
                      overflow: 'hidden',
                      textOverflow: 'ellipsis',
                    }}>
                      {formatParameters(rule.parameters, rule.id)}
                    </TableCell>
                    <TableCell sx={{ color: 'text.primary', fontWeight: 600 }}>
                      {rule.violations}
                    </TableCell>
                    <TableCell align="center">
                      <Tooltip title="Edit Rule">
                        <IconButton
                          size="small"
                          sx={{ 
                            color: 'primary.light',
                            '&:hover': { background: 'rgba(59, 130, 246, 0.1)' }
                          }}
                          onClick={() => handleEditClick(rule)}
                        >
                          <EditIcon fontSize="small" />
                        </IconButton>
                      </Tooltip>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </StyledTableContainer>
        </CardContent>
      </StyledCard>

      {/* Edit Rule Dialog */}
      <StyledDialog
        open={dialogOpen}
        onClose={handleCloseDialog}
        maxWidth="md"
        fullWidth
      >
        <DialogTitle sx={{ 
          color: 'text.primary', 
          fontWeight: 600,
          borderBottom: '1px solid rgba(100, 116, 139, 0.2)',
        }}>
          Edit Rule {editingRule ? `R${String(editingRule.id).padStart(3, '0')}` : ''}: {editingRule?.name}
        </DialogTitle>
        <DialogContent sx={{ pt: 3 }}>
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="Description"
                multiline
                rows={3}
                value={editForm.description}
                onChange={(e) => setEditForm({ ...editForm, description: e.target.value })}
                sx={{
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
            </Grid>
            <Grid item xs={6}>
              <FormControl fullWidth>
                <InputLabel>Status</InputLabel>
                <Select
                  value={editForm.status}
                  label="Status"
                  onChange={(e) => setEditForm({ ...editForm, status: e.target.value })}
                  sx={{
                    background: 'rgba(30, 41, 59, 0.7)',
                    '& .MuiOutlinedInput-notchedOutline': {
                      borderColor: 'rgba(100, 116, 139, 0.3)',
                    },
                    '&:hover .MuiOutlinedInput-notchedOutline': {
                      borderColor: 'rgba(100, 116, 139, 0.5)',
                    },
                  }}
                >
                  <MenuItem value="active">Active</MenuItem>
                  <MenuItem value="inactive">Inactive</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <TextField
                fullWidth
                label="Rule Name"
                value={editForm.name}
                disabled
                sx={{
                  '& .MuiOutlinedInput-root': {
                    background: 'rgba(30, 41, 59, 0.5)',
                    '& .MuiOutlinedInput-notchedOutline': {
                      borderColor: 'rgba(100, 116, 139, 0.2)',
                    },
                  },
                }}
              />
            </Grid>
            {renderParameterEditor()}
          </Grid>
        </DialogContent>
        <DialogActions sx={{ 
          p: 3, 
          borderTop: '1px solid rgba(100, 116, 139, 0.2)',
          gap: 1,
        }}>
          <Button 
            onClick={handleCloseDialog}
            sx={{ color: 'text.secondary' }}
          >
            Cancel
          </Button>
          <Button 
            onClick={handleSave}
            variant="contained"
            disabled={saving}
            sx={{
              background: 'linear-gradient(135deg, #3b82f6, #1d4ed8)',
              '&:hover': {
                background: 'linear-gradient(135deg, #1d4ed8, #1e40af)',
              },
            }}
          >
            {saving ? <CircularProgress size={20} /> : 'Save Changes'}
          </Button>
        </DialogActions>
      </StyledDialog>
    </Box>
  );
}; 