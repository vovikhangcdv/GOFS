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
} from '@mui/material';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import axios from 'axios';

interface Rule {
  id: number;
  name: string;
  status: string;
  violations: number;
}

export const ComplianceRulesTable = () => {
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
    <Card sx={{ mt: 4 }}>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          📋 ACTIVE COMPLIANCE RULES
        </Typography>
        {loading ? (
          <Box sx={{ display: 'flex', justifyContent: 'center', p: 2 }}>
            <CircularProgress />
          </Box>
        ) : error ? (
          <Alert severity="error">{error}</Alert>
        ) : (
          <TableContainer component={Paper}>
            <Table size="small">
              <TableHead>
                <TableRow>
                  <TableCell>Rule ID</TableCell>
                  <TableCell>Name</TableCell>
                  <TableCell>Status</TableCell>
                  <TableCell>Violations/24h</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {rules.map((rule, idx) => (
                  <TableRow key={rule.id}>
                    <TableCell>{`R${String(rule.id).padStart(3, '0')}`}</TableCell>
                    <TableCell>{rule.name}</TableCell>
                    <TableCell>
                      {rule.status === 'active' ? (
                        <Chip
                          icon={<CheckCircleIcon style={{ color: '#43a047' }} />}
                          label="Active"
                          color="success"
                          size="small"
                        />
                      ) : (
                        <Chip label="Inactive" color="default" size="small" />
                      )}
                    </TableCell>
                    <TableCell>{rule.violations}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        )}
      </CardContent>
    </Card>
  );
}; 