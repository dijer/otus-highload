import { Container, Paper, Typography } from '@mui/material';
import { useParams } from 'react-router-dom';
import User from '../../features/auth/ui/user';

export const UserPage = () => {
	const { id } = useParams<{ id: string }>();

	return (
		<Container maxWidth="md">
			<Paper elevation={3} sx={{ p: 4 }}>
				<Typography variant="h4" component="h1" gutterBottom>
					User
				</Typography>

				<User userId={Number(id)} />
			</Paper>
		</Container>
	);
};
