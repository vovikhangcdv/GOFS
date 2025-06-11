import {
  Card,
  CardContent,
  Typography,
  Box,
  Chip,
  styled,
} from '@mui/material';
import {
  TrendingUp as TrendingUpIcon,
  TrendingDown as TrendingDownIcon,
} from '@mui/icons-material';

const StyledCard = styled(Card)(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
  borderRadius: theme.shape.borderRadius * 2,
  boxShadow: '0 4px 12px rgba(0,0,0,0.05)',
  transition: 'transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out',
  '&:hover': {
    transform: 'translateY(-4px)',
    boxShadow: '0 8px 24px rgba(0,0,0,0.1)',
  },
}));

interface TransactionCardProps {
  hash: string;
  type: 'in' | 'out';
  amount: string;
  timestamp: string;
  status: 'completed' | 'pending' | 'failed';
  from: string;
  to: string;
}

export const TransactionCard = ({
  hash,
  type,
  amount,
  timestamp,
  status,
  from,
  to,
}: TransactionCardProps) => {
  const getStatusColor = (status: string) => {
    switch (status) {
      case 'completed':
        return 'success';
      case 'pending':
        return 'warning';
      case 'failed':
        return 'error';
      default:
        return 'default';
    }
  };

  return (
    <StyledCard>
      <CardContent>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 2 }}>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
            {type === 'in' ? (
              <TrendingUpIcon color="success" />
            ) : (
              <TrendingDownIcon color="error" />
            )}
            <Typography variant="h6" component="div">
              {amount}
            </Typography>
          </Box>
          <Chip
            label={status}
            color={getStatusColor(status)}
            size="small"
            sx={{ textTransform: 'capitalize' }}
          />
        </Box>
        
        <Typography variant="body2" color="text.secondary" gutterBottom>
          Hash: {hash.slice(0, 8)}...{hash.slice(-8)}
        </Typography>
        
        <Box sx={{ mt: 2 }}>
          <Typography variant="body2" color="text.secondary">
            From: {from.slice(0, 8)}...{from.slice(-8)}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            To: {to.slice(0, 8)}...{to.slice(-8)}
          </Typography>
        </Box>
        
        <Typography
          variant="caption"
          color="text.secondary"
          sx={{ display: 'block', mt: 2 }}
        >
          {new Date(timestamp).toLocaleString()}
        </Typography>
      </CardContent>
    </StyledCard>
  );
}; 