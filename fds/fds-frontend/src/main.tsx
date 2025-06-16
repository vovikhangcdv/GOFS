import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { CssBaseline, ThemeProvider, createTheme } from '@mui/material';

const theme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#64748b', // Slate gray
      dark: '#475569',
      light: '#94a3b8',
    },
    secondary: {
      main: '#6b7280', // Gray
      dark: '#4b5563',
      light: '#9ca3af',
    },
    background: {
      default: 'linear-gradient(135deg, #0f172a 0%, #1e293b 100%)',
      paper: 'rgba(30, 41, 59, 0.8)',
    },
    text: {
      primary: '#f8fafc',
      secondary: '#94a3b8',
    },
    error: {
      main: '#ef4444',
      light: '#f87171',
    },
    warning: {
      main: '#f59e0b',
      light: '#fbbf24',
    },
    info: {
      main: '#06b6d4',
      light: '#22d3ee',
    },
    success: {
      main: '#10b981',
      light: '#34d399',
    },
    divider: 'rgba(148, 163, 184, 0.12)',
  },
  typography: {
    fontFamily: '"Inter", "SF Pro Display", -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif',
    h1: {
      fontWeight: 700,
      fontSize: '2.5rem',
      background: 'linear-gradient(135deg, #f8fafc, #e2e8f0)',
      WebkitBackgroundClip: 'text',
      WebkitTextFillColor: 'transparent',
    },
    h4: {
      fontWeight: 600,
      fontSize: '1.75rem',
    },
    h6: {
      fontWeight: 600,
      fontSize: '1.125rem',
    },
    body1: {
      fontSize: '0.95rem',
      lineHeight: 1.6,
    },
    body2: {
      fontSize: '0.875rem',
      lineHeight: 1.5,
    },
  },
  shape: {
    borderRadius: 12,
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        body: {
          background: 'linear-gradient(135deg, #0f172a 0%, #1e293b 50%, #334155 100%)',
          minHeight: '100vh',
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          background: 'rgba(30, 41, 59, 0.8)',
          backdropFilter: 'blur(20px)',
          border: '1px solid rgba(100, 116, 139, 0.2)',
          boxShadow: '0 8px 32px rgba(0, 0, 0, 0.4)',
          '&:hover': {
            border: '1px solid rgba(100, 116, 139, 0.3)',
            transform: 'translateY(-2px)',
            transition: 'all 0.3s ease',
            boxShadow: '0 12px 40px rgba(0, 0, 0, 0.5)',
          },
        },
      },
    },
    MuiButton: {
      styleOverrides: {
        root: {
          textTransform: 'none',
          fontWeight: 600,
          borderRadius: 8,
          padding: '10px 20px',
        },
        contained: {
          background: 'linear-gradient(135deg, #64748b, #475569)',
          boxShadow: '0 4px 14px rgba(100, 116, 139, 0.3)',
          '&:hover': {
            background: 'linear-gradient(135deg, #475569, #334155)',
            boxShadow: '0 6px 20px rgba(100, 116, 139, 0.4)',
          },
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundImage: 'none',
          background: 'rgba(30, 41, 59, 0.8)',
          backdropFilter: 'blur(20px)',
          border: '1px solid rgba(100, 116, 139, 0.2)',
        },
      },
    },
    MuiTableContainer: {
      styleOverrides: {
        root: {
          borderRadius: 12,
          border: '1px solid rgba(100, 116, 139, 0.2)',
        },
      },
    },
    MuiTableHead: {
      styleOverrides: {
        root: {
          background: 'linear-gradient(135deg, rgba(100, 116, 139, 0.15), rgba(71, 85, 105, 0.15))',
        },
      },
    },
  },
});

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <App />
    </ThemeProvider>
  </React.StrictMode>
); 