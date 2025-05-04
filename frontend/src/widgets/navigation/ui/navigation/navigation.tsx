import { route } from '../../../../shared/constants/routes';
import { AppBar, Box, Container, Toolbar } from '@mui/material';
import { AuthNavigation } from './auth-navigation';
import NavLink from '../nav-link';

export const Navigation = () => {
	return (
		<>
			<AppBar position="fixed">
				<Container maxWidth="md">
					<Toolbar sx={{ color: 'inherit' }} disableGutters>
						<Box sx={{ display: 'flex', gap: 2, flexGrow: 1 }}>
							<NavLink to={route.HOME}>Home</NavLink>
						</Box>
						<AuthNavigation />
					</Toolbar>
				</Container>
			</AppBar>
			<Toolbar sx={{ mb: 4 }} />
		</>
	);
};
