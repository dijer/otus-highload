import { Autocomplete, Box, Grid, TextField } from '@mui/material';
import { IUser } from '../../../../entities/user';

interface IProps extends IUser {}

export const UserInfo = ({
	username,
	firstName,
	lastName,
	birthday,
	city,
	gender,
	interests,
}: IProps) => {
	return (
		<Box>
			<Grid container spacing={2}>
				<Grid size={6}>
					<TextField
						disabled
						fullWidth
						name="firstName"
						label="First Name"
						value={firstName}
						slotProps={{ inputLabel: { shrink: true } }}
						margin="normal"
					/>
				</Grid>
				<Grid size={6}>
					<TextField
						disabled
						fullWidth
						name="lastName"
						label="Last Name"
						value={lastName}
						slotProps={{ inputLabel: { shrink: true } }}
						margin="normal"
					/>
				</Grid>
			</Grid>
			<Grid container spacing={2}>
				<Grid size={6}>
					<TextField
						disabled
						fullWidth
						name="username"
						label="Username"
						value={username}
						slotProps={{ inputLabel: { shrink: true } }}
						margin="normal"
						required
					/>
				</Grid>
			</Grid>
			<Grid container spacing={2}>
				<Grid size={6}>
					<TextField
						disabled
						fullWidth
						name="birthday"
						label="Birthday"
						type="date"
						slotProps={{ inputLabel: { shrink: true } }}
						value={birthday}
						margin="normal"
					/>
				</Grid>
				<Grid size={6}>
					<TextField
						disabled
						fullWidth
						name="gender"
						label="Gender"
						value={gender}
						slotProps={{ inputLabel: { shrink: true } }}
						margin="normal"
					/>
				</Grid>
			</Grid>
			<Grid container spacing={2}>
				<Grid size={6}>
					<Autocomplete
						disabled
						multiple
						freeSolo
						options={['football', 'guitar']}
						value={interests}
						renderInput={(params) => (
							<TextField
								{...params}
								name="interests"
								label="Interests"
								placeholder="Add Interest"
								slotProps={{ inputLabel: { shrink: true } }}
								margin="normal"
							/>
						)}
					/>
				</Grid>
				<Grid size={6}>
					<TextField
						disabled
						fullWidth
						name="city"
						label="City"
						value={city}
						slotProps={{ inputLabel: { shrink: true } }}
						margin="normal"
					/>
				</Grid>
			</Grid>
		</Box>
	);
};
