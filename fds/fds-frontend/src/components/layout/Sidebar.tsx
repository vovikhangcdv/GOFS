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
} from '@mui/material';
import {
  Dashboard as DashboardIcon,
  SwapHoriz as TransactionsIcon,
  Settings as SettingsIcon,
  AccountBalance as WalletIcon,
} from '@mui/icons-material';

const drawerWidth = 240;

const StyledDrawer = styled(Drawer)(({ theme }) => ({
  width: drawerWidth,
  flexShrink: 0,
  '& .MuiDrawer-paper': {
    width: drawerWidth,
    boxSizing: 'border-box',
    backgroundColor: theme.palette.background.default,
    borderRight: `1px solid ${theme.palette.divider}`,
  },
}));

interface SidebarProps {
  onPageChange: (page: string) => void;
}

const menuItems = [
  { text: 'Dashboard', icon: <DashboardIcon />, page: 'dashboard' },
  { text: 'Transactions', icon: <TransactionsIcon />, page: 'transactions' },
  { text: 'Wallet', icon: <WalletIcon />, page: 'wallet' },
  { text: 'Blacklist View', icon: <SettingsIcon />, page: 'blacklist' },
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
      <Box sx={{ p: 2, display: 'flex', justifyContent: 'center' }}>
        <img src="/icon.png" alt="Logo" style={{ height: 100 }} />
      </Box>
      <List>
        {menuItems.map((item, index) => (
          <ListItem key={item.text} disablePadding>
            <ListItemButton
              selected={selected === index}
              onClick={() => handleItemClick(index, item.page)}
              sx={{
                mx: 1,
                borderRadius: 1,
                '&.Mui-selected': {
                  backgroundColor: theme.palette.primary.main,
                  color: theme.palette.primary.contrastText,
                  '&:hover': {
                    backgroundColor: theme.palette.primary.dark,
                  },
                },
              }}
            >
              <ListItemIcon
                sx={{
                  color: selected === index ? 'inherit' : theme.palette.text.secondary,
                }}
              >
                {item.icon}
              </ListItemIcon>
              <ListItemText primary={item.text} />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </StyledDrawer>
  );
}; 