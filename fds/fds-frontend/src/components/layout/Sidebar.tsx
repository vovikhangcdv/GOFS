import { useState } from 'react';
import {
  Box,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  styled,
  useTheme,
  Typography,
  Divider,
} from '@mui/material';
import {
  Dashboard as DashboardIcon,
  SwapHoriz as TransactionsIcon,
  Settings as SettingsIcon,
  AccountBalance as WalletIcon,
  Block as BlockIcon,
  Security as SecurityIcon,
} from '@mui/icons-material';

const drawerWidth = 240;

const StyledDrawer = styled(Drawer)(({ theme }) => ({
  width: drawerWidth,
  flexShrink: 0,
  '& .MuiDrawer-paper': {
    width: drawerWidth,
    boxSizing: 'border-box',
    background: 'rgba(15, 23, 42, 0.95)',
    backdropFilter: 'blur(20px)',
    border: 'none',
    borderRight: 'none',
    marginTop: '64px', // Account for header height
    height: 'calc(100vh - 64px)',
  },
}));

const LogoContainer = styled(Box)(({ theme }) => ({
  padding: theme.spacing(3, 2),
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  background: 'rgba(30, 41, 59, 0.3)',
  borderBottom: '1px solid rgba(100, 116, 139, 0.2)',
  marginBottom: theme.spacing(2),
}));

const LogoIcon = styled(Box)(({ theme }) => ({
  width: 48,
  height: 48,
  borderRadius: 12,
  background: 'linear-gradient(135deg, #64748b, #475569)',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  color: 'white',
  fontWeight: 'bold',
  fontSize: '1.5rem',
  marginBottom: theme.spacing(1),
  boxShadow: '0 8px 24px rgba(100, 116, 139, 0.3)',
}));

const StyledListItemButton = styled(ListItemButton)(({ theme }) => ({
  margin: theme.spacing(0.5, 1),
  borderRadius: 12,
  padding: theme.spacing(1.5, 2),
  transition: 'all 0.3s ease',
  border: '1px solid transparent',
  '&:hover': {
    background: 'rgba(100, 116, 139, 0.15)',
    border: '1px solid rgba(100, 116, 139, 0.3)',
    transform: 'translateX(4px)',
  },
  '&.Mui-selected': {
    background: 'rgba(100, 116, 139, 0.25)',
    border: '1px solid rgba(100, 116, 139, 0.4)',
    color: theme.palette.primary.light,
    '&::before': {
      content: '""',
      position: 'absolute',
      left: 0,
      top: '50%',
      transform: 'translateY(-50%)',
      width: 4,
      height: '60%',
      background: 'linear-gradient(135deg, #64748b, #475569)',
      borderRadius: '0 2px 2px 0',
    },
    '&:hover': {
      background: 'rgba(100, 116, 139, 0.3)',
      transform: 'translateX(4px)',
    },
  },
}));

const MenuSection = styled(Box)(({ theme }) => ({
  padding: theme.spacing(0, 2, 2),
}));

const SectionTitle = styled(Typography)(({ theme }) => ({
  fontSize: '0.75rem',
  fontWeight: 600,
  color: theme.palette.text.secondary,
  textTransform: 'uppercase',
  letterSpacing: '0.1em',
  marginBottom: theme.spacing(1),
  marginLeft: theme.spacing(2),
}));

const StatusCard = styled(Box)(({ theme }) => ({
  padding: theme.spacing(2),
  borderRadius: 12,
  background: 'rgba(30, 41, 59, 0.5)',
  border: '1px solid rgba(100, 116, 139, 0.3)',
  textAlign: 'center',
}));

interface SidebarProps {
  onPageChange: (page: string) => void;
}

const menuItems = [
  { text: 'Dashboard', icon: <DashboardIcon />, page: 'dashboard' },
  { text: 'Suspicious Transactions', icon: <TransactionsIcon />, page: 'transactions' },
  { text: 'Wallet Analysis', icon: <WalletIcon />, page: 'wallet' },
  { text: 'Rules', icon: <SecurityIcon />, page: 'rules' },
  { text: 'Blacklist', icon: <BlockIcon />, page: 'blacklist' },
];

export const Sidebar = ({ onPageChange }: SidebarProps) => {
  const theme = useTheme();
  const [selected, setSelected] = useState(0);

  const handleItemClick = (index: number, page: string) => {
    setSelected(index);
    onPageChange(page);
  };

  return (
    <StyledDrawer variant="permanent">
      <LogoContainer>
        <LogoIcon>FDS</LogoIcon>
        <Typography 
          variant="subtitle1" 
          sx={{ 
            fontWeight: 600,
            color: 'text.primary',
            textAlign: 'center',
            lineHeight: 1.2,
            fontSize: '0.9rem',
          }}
        >

        </Typography>
      </LogoContainer>
      
      <MenuSection>
        <SectionTitle></SectionTitle>
        <List sx={{ padding: 0 }}>
          {menuItems.map((item, index) => (
            <ListItem key={item.text} disablePadding>
              <StyledListItemButton
                selected={selected === index}
                onClick={() => handleItemClick(index, item.page)}
              >
                <ListItemIcon
                  sx={{
                    color: selected === index ? 'primary.light' : 'text.secondary',
                    minWidth: 40,
                    transition: 'color 0.3s ease',
                  }}
                >
                  {item.icon}
                </ListItemIcon>
                <ListItemText 
                  primary={item.text}
                  primaryTypographyProps={{
                    fontSize: '0.9rem',
                    fontWeight: selected === index ? 600 : 500,
                  }}
                />
              </StyledListItemButton>
            </ListItem>
          ))}
        </List>
      </MenuSection>

      <Box sx={{ mt: 'auto', p: 2 }}>
        <Divider sx={{ mb: 2, borderColor: 'rgba(100, 116, 139, 0.2)' }} />
        <StatusCard>
          <Typography variant="caption" color="text.secondary" sx={{ fontSize: '0.7rem', fontWeight: 600 }}>
            System Status
          </Typography>
          <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', mt: 1 }}>
            <Box
              sx={{
                width: 8,
                height: 8,
                borderRadius: '50%',
                background: 'linear-gradient(135deg, #10b981, #34d399)',
                mr: 1,
                animation: 'pulse 2s infinite',
                '@keyframes pulse': {
                  '0%, 100%': { opacity: 1 },
                  '50%': { opacity: 0.5 },
                },
              }}
            />
            <Typography variant="caption" sx={{ color: 'success.light', fontWeight: 600, fontSize: '0.75rem' }}>
              Online
            </Typography>
          </Box>
        </StatusCard>
      </Box>
    </StyledDrawer>
  );
}; 