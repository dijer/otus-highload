import { Container, Paper, Typography } from '@mui/material';
import RegisterForm from '../../features/auth/ui/register-form';

export const RegisterPage = () => {
	return (
		<Container maxWidth="md">
			<Paper elevation={3} sx={{ p: 4 }}>
				<Typography variant="h4" component="h1" gutterBottom>
					Register
				</Typography>

				<RegisterForm />
			</Paper>
		</Container>
	);
};
