import { Container, Paper, Typography } from '@mui/material';
import SearchContainer from '../../features/search/containers/search';

export const SearchPage = () => {
	return (
		<Container maxWidth="md">
			<Paper elevation={3} sx={{ p: 4 }}>
				<Typography variant="h4" component="h1" gutterBottom>
					Search Users
				</Typography>
				<SearchContainer />
			</Paper>
		</Container>
	);
};
