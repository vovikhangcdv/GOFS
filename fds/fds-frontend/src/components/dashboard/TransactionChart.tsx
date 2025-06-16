import { useState, useEffect } from 'react';
import {
  Box,
  Card,
  CardContent,
  Typography,
  CircularProgress,
  Alert,
  Tabs,
  Tab,
  styled,
} from '@mui/material';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  BarElement,
  Title,
  Tooltip,
  Legend,
  Filler,
} from 'chart.js';
import { Line, Bar } from 'react-chartjs-2';
import { getTransactionStats } from '../../api';
import { TransactionStats } from '../../types';

// Register Chart.js components
ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  BarElement,
  Title,
  Tooltip,
  Legend,
  Filler
);

const StyledCard = styled(Card)(({ theme }) => ({
  background: 'rgba(30, 41, 59, 0.8)',
  backdropFilter: 'blur(20px)',
  border: '1px solid rgba(100, 116, 139, 0.3)',
  borderRadius: 16,
  overflow: 'hidden',
  height: '100%',
}));

const StyledTabs = styled(Tabs)(({ theme }) => ({
  '& .MuiTabs-indicator': {
    backgroundColor: theme.palette.primary.main,
    height: 3,
    borderRadius: '3px 3px 0 0',
  },
  '& .MuiTab-root': {
    color: theme.palette.text.secondary,
    fontWeight: 600,
    fontSize: '0.85rem',
    textTransform: 'none',
    minHeight: 'auto',
    padding: '12px 16px',
    '&.Mui-selected': {
      color: theme.palette.primary.main,
    },
  },
}));

export const TransactionChart = () => {
  const [stats, setStats] = useState<TransactionStats | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [activeTab, setActiveTab] = useState(0);

  const fetchStats = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await getTransactionStats();
      setStats(data);
    } catch (err) {
      setError('Failed to fetch transaction statistics');
      console.error('Transaction stats fetch error:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchStats();
    // Refresh every 30 seconds
    const interval = setInterval(fetchStats, 30000);
    return () => clearInterval(interval);
  }, []);

  const handleTabChange = (event: React.SyntheticEvent, newValue: number) => {
    setActiveTab(newValue);
  };

  if (loading) {
    return (
      <StyledCard>
        <CardContent sx={{ 
          display: 'flex', 
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          minHeight: 300,
          gap: 2,
        }}>
          <CircularProgress 
            size={40}
            sx={{
              color: 'primary.main',
              '& .MuiCircularProgress-circle': {
                strokeLinecap: 'round',
              },
            }}
          />
          <Typography variant="body2" color="text.secondary">
            Loading transaction data...
          </Typography>
        </CardContent>
      </StyledCard>
    );
  }

  if (error || !stats) {
    return (
      <StyledCard>
        <CardContent>
          <Alert 
            severity="error" 
            sx={{
              background: 'rgba(239, 68, 68, 0.1)',
              border: '1px solid rgba(239, 68, 68, 0.3)',
              borderRadius: 2,
            }}
          >
            {error || 'No data available'}
          </Alert>
        </CardContent>
      </StyledCard>
    );
  }

  // Daily chart data
  const dailyChartData = {
    labels: stats.daily_stats.map(item => {
      const date = new Date(item.date);
      return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
    }),
    datasets: [
      {
        label: 'Transactions',
        data: stats.daily_stats.map(item => item.count),
        fill: true,
        borderColor: 'rgba(59, 130, 246, 1)',
        backgroundColor: 'rgba(59, 130, 246, 0.1)',
        borderWidth: 2,
        pointBackgroundColor: 'rgba(59, 130, 246, 1)',
        pointBorderColor: '#ffffff',
        pointBorderWidth: 2,
        pointRadius: 4,
        tension: 0.4,
      },
    ],
  };

  // Hourly chart data
  const hourlyChartData = {
    labels: stats.hourly_stats.map(item => item.hour),
    datasets: [
      {
        label: 'Transactions',
        data: stats.hourly_stats.map(item => item.count),
        backgroundColor: 'rgba(34, 197, 94, 0.8)',
        borderColor: 'rgba(34, 197, 94, 1)',
        borderWidth: 1,
        borderRadius: 4,
        borderSkipped: false,
      },
    ],
  };

  const chartOptions = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        display: false,
      },
      tooltip: {
        backgroundColor: 'rgba(30, 41, 59, 0.95)',
        titleColor: '#ffffff',
        bodyColor: '#ffffff',
        borderColor: 'rgba(100, 116, 139, 0.3)',
        borderWidth: 1,
        cornerRadius: 8,
        displayColors: false,
      },
    },
    scales: {
      x: {
        grid: {
          color: 'rgba(100, 116, 139, 0.2)',
          drawBorder: false,
        },
        ticks: {
          color: 'rgba(148, 163, 184, 0.8)',
          font: {
            size: 11,
          },
        },
      },
      y: {
        grid: {
          color: 'rgba(100, 116, 139, 0.2)',
          drawBorder: false,
        },
        ticks: {
          color: 'rgba(148, 163, 184, 0.8)',
          font: {
            size: 11,
          },
        },
        beginAtZero: true,
      },
    },
  };

  return (
    <StyledCard>
      <CardContent sx={{ p: { xs: 2, md: 3 }, height: '100%', display: 'flex', flexDirection: 'column' }}>
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
              mb: 1,
            }}
          >
            📊 Transaction Analytics
          </Typography>
          
          {/* Summary Stats */}
          <Box sx={{ display: 'flex', gap: 3, mb: 2 }}>
            <Box>
              <Typography variant="caption" color="text.secondary">
                Total Transactions
              </Typography>
              <Typography variant="h6" sx={{ color: 'primary.main', fontWeight: 700 }}>
                {stats.total_transactions.toLocaleString()}
              </Typography>
            </Box>
            <Box>
              <Typography variant="caption" color="text.secondary">
                Suspicious
              </Typography>
              <Typography variant="h6" sx={{ color: 'error.main', fontWeight: 700 }}>
                {stats.suspicious_transactions.toLocaleString()}
              </Typography>
            </Box>
          </Box>
        </Box>

        <StyledTabs 
          value={activeTab} 
          onChange={handleTabChange}
          sx={{ mb: 2, minHeight: 'auto' }}
        >
          <Tab label="Daily (7 days)" />
          <Tab label="Hourly (24h)" />
        </StyledTabs>
        
        <Box sx={{ flexGrow: 1, minHeight: 200 }}>
          {activeTab === 0 ? (
            <Line data={dailyChartData} options={chartOptions} />
          ) : (
            <Bar data={hourlyChartData} options={chartOptions} />
          )}
        </Box>
      </CardContent>
    </StyledCard>
  );
}; 