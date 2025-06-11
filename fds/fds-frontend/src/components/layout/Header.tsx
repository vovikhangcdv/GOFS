import {
  AppBar,
  Box,
  Toolbar,
  Typography,
  styled,
} from '@mui/material';
import { Search as SearchIcon } from '@mui/icons-material';

const StyledAppBar = styled(AppBar)(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
  color: theme.palette.text.primary,
  boxShadow: 'none',
  borderBottom: `1px solid ${theme.palette.divider}`,
}));

const SearchBar = styled(Box)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  backgroundColor: theme.palette.background.default,
  borderRadius: theme.shape.borderRadius,
  padding: theme.spacing(0, 2),
  marginLeft: 'auto',
  width: 300,
}));

export const Header = () => {
  return (
    <StyledAppBar position="fixed">
      <Toolbar>
        <SearchBar>
          <SearchIcon sx={{ color: 'text.secondary', mr: 1 }} />
          <Typography variant="body2" color="text.secondary">
            Search transactions...
          </Typography>
        </SearchBar>
      </Toolbar>
    </StyledAppBar>
  );
}; 