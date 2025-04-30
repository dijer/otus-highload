import { Container, Paper, Typography } from '@mui/material';
import LoginForm from '../../features/auth/ui/login-form';

export const LoginPage = () => {
	return (
		<Container maxWidth="md">
			<Paper elevation={3} sx={{ p: 4 }}>
				<Typography variant="h4" component="h1" gutterBottom>
					Login
				</Typography>

				<LoginForm />
			</Paper>
		</Container>
	);
};
