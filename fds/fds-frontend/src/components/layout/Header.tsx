import { useState } from 'react';
import {
  AppBar,
  Box,
  Toolbar,
  Typography,
  styled,
  InputBase,
  IconButton,
  Badge,
} from '@mui/material';
import { 
  Search as SearchIcon,
  Notifications as NotificationsIcon,
  AccountCircle as AccountIcon,
} from '@mui/icons-material';

const StyledAppBar = styled(AppBar)(({ theme }) => ({
  background: 'rgba(15, 23, 42, 0.95)',
  backdropFilter: 'blur(20px)',
  borderBottom: '1px solid rgba(100, 116, 139, 0.3)',
  boxShadow: '0 8px 32px rgba(0, 0, 0, 0.4)',
}));

const SearchContainer = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  background: 'rgba(30, 41, 59, 0.8)',
  backdropFilter: 'blur(15px)',
  border: '1px solid rgba(100, 116, 139, 0.4)',
  borderRadius: 12,
  padding: theme.spacing(0.5, 2),
  marginLeft: 'auto',
  marginRight: theme.spacing(2),
  width: 320,
  transition: 'all 0.3s ease',
  '&:hover': {
    border: '1px solid rgba(59, 130, 246, 0.6)',
    background: 'rgba(30, 41, 59, 0.95)',
    boxShadow: '0 4px 20px rgba(59, 130, 246, 0.2)',
  },
  '&:focus-within': {
    border: '1px solid rgba(59, 130, 246, 0.8)',
    boxShadow: '0 0 0 3px rgba(59, 130, 246, 0.2)',
  },
}));

const SearchInput = styled(InputBase)(({ theme }) => ({
  marginLeft: theme.spacing(1),
  flex: 1,
  color: theme.palette.text.primary,
  '& .MuiInputBase-input': {
    fontSize: '0.9rem',
    '&::placeholder': {
      color: theme.palette.text.secondary,
      opacity: 0.8,
    },
  },
}));

const TitleBox = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  gap: theme.spacing(2),
}));

const LogoIcon = styled(Box)(({ theme }) => ({
  width: 32,
  height: 32,
  borderRadius: 8,
  background: 'linear-gradient(135deg, #64748b, #475569)',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  color: 'white',
  fontWeight: 'bold',
  fontSize: '1.2rem',
  boxShadow: '0 4px 12px rgba(100, 116, 139, 0.3)',
}));

const ActionButton = styled(IconButton)(({ theme }) => ({
  background: 'rgba(30, 41, 59, 0.7)',
  border: '1px solid rgba(100, 116, 139, 0.3)',
  color: theme.palette.text.primary,
  transition: 'all 0.3s ease',
  '&:hover': {
    background: 'rgba(30, 41, 59, 0.9)',
    border: '1px solid rgba(59, 130, 246, 0.5)',
    boxShadow: '0 4px 16px rgba(59, 130, 246, 0.25)',
  },
}));

interface HeaderProps {
  onSearchAddress?: (address: string) => void;
}

export const Header = ({ onSearchAddress }: HeaderProps) => {
  const [searchTerm, setSearchTerm] = useState('');

  const handleSearchSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    if (searchTerm.trim() && onSearchAddress) {
      onSearchAddress(searchTerm.trim());
      setSearchTerm('');
    }
  };

  const handleKeyPress = (event: React.KeyboardEvent) => {
    if (event.key === 'Enter') {
      handleSearchSubmit(event);
    }
  };

  return (
    <StyledAppBar position="fixed" sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
      <Toolbar sx={{ justifyContent: 'space-between', minHeight: '64px !important' }}>
        <Box sx={{ ml: '240px' }} />
        
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
          <SearchContainer component="form" onSubmit={handleSearchSubmit}>
            <SearchIcon 
              sx={{ 
                color: 'text.secondary', 
                fontSize: 20,
                cursor: searchTerm.trim() ? 'pointer' : 'default',
                '&:hover': {
                  color: searchTerm.trim() ? 'primary.main' : 'text.secondary',
                }
              }} 
              onClick={searchTerm.trim() ? handleSearchSubmit : undefined}
            />
            <SearchInput
              placeholder="Search transactions, addresses..."
              inputProps={{ 'aria-label': 'search' }}
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              onKeyPress={handleKeyPress}
            />
          </SearchContainer>
          
          <ActionButton>
            <Badge badgeContent={3} color="error">
              <NotificationsIcon fontSize="small" />
            </Badge>
          </ActionButton>
          
          <ActionButton>
            <AccountIcon fontSize="small" />
          </ActionButton>
        </Box>
      </Toolbar>
    </StyledAppBar>
  );
}; 