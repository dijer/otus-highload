import { Container, Paper, Typography } from '@mui/material';

export const HomePage = () => {
	return (
		<Container maxWidth="md">
			<Paper elevation={3} sx={{ p: 4 }}>
				<Typography variant="h4" component="h1" gutterBottom>
					Home
				</Typography>
				Welcome!
			</Paper>
		</Container>
	);
};
