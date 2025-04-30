import { Link, useLocation } from 'react-router-dom';
import { route } from '../../../shared/constants/routes';
import { AppBar, Box, Button, Container, Toolbar } from '@mui/material';
import { AuthNavigation } from './auth-navigation';

export const Navigation = () => {
	const location = useLocation();

	return (
		<>
			<AppBar position="fixed">
				<Container maxWidth="md">
					<Toolbar sx={{ color: 'inherit' }} disableGutters>
						<Box sx={{ display: 'flex', gap: 2, flexGrow: 1 }}>
							<Button
								sx={{
									color: 'inherit',
									textDecoration:
										location.pathname === route.HOME
											? 'underline'
											: 'inherit',
								}}
								to={route.HOME}
								component={Link}
							>
								Home
							</Button>
						</Box>
						<AuthNavigation />
					</Toolbar>
				</Container>
			</AppBar>
			<Toolbar sx={{ mb: 4 }} />
		</>
	);
};
