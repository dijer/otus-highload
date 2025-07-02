import {
	Box,
	Grid,
	Paper,
	Table,
	TableBody,
	TableCell,
	TableContainer,
	TableHead,
	TableRow,
	Typography,
} from '@mui/material';
import { useAppSelector } from '../../../../app/hooks/use-app';

type TProps = {
	isLoading: boolean;
	isUninitialized: boolean;
};

export const Users = ({ isLoading, isUninitialized }: TProps) => {
	const users = useAppSelector((state) => state.search.users);

	console.log(users);

	if (isUninitialized) {
		return null;
	}

	if (isLoading) {
		return <Typography>Loading...</Typography>;
	}

	if (!users.length) {
		return <Typography>Users not found.</Typography>;
	}

	return (
		<Box>
			<Grid container spacing={2}>
				<Grid size={12}>
					<TableContainer component={Paper}>
						<Table aria-label="users table">
							<TableHead>
								<TableRow>
									<TableCell>
										<b>First Name</b>
									</TableCell>
									<TableCell>
										<b>Last Name</b>
									</TableCell>
								</TableRow>
							</TableHead>
							<TableBody>
								{users.map(({ firstName, lastName }, index) => (
									<TableRow key={index}>
										<TableCell>{firstName}</TableCell>
										<TableCell>{lastName}</TableCell>
									</TableRow>
								))}
							</TableBody>
						</Table>
					</TableContainer>
				</Grid>
			</Grid>
		</Box>
	);
};
