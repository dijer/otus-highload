import { Container, Paper, Typography } from '@mui/material';
import Profile from '../../features/auth/ui/profile';

export const ProfilePage = () => {
	return (
		<Container maxWidth="md">
			<Paper elevation={3} sx={{ p: 4 }}>
				<Typography variant="h4" component="h1" gutterBottom>
					Profile
				</Typography>

				<Profile />
			</Paper>
		</Container>
	);
};
