import {
	Autocomplete,
	Box,
	Button,
	Grid,
	MenuItem,
	TextField,
} from '@mui/material';
import { ChangeEvent, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { IUserWithPassword } from '../../model/user.model';
import { useToast } from '../../../../app/ui/toast-provider/toast-provider';
import { useRegisterMutation } from '../../model/auth.api';
import { genders } from '../../../../shared/constants/genders';
import { route } from '../../../../shared/constants/routes';

export const RegisterForm = () => {
	const [formData, setFormData] = useState<
		Partial<Partial<IUserWithPassword>>
	>({
		firstName: '',
		lastName: '',
		birthday: undefined,
		city: '',
		gender: '',
		interests: [],
		username: '',
		password: '',
	});
	const [register, { isLoading }] = useRegisterMutation();
	const navigate = useNavigate();
	const { showToast } = useToast();

	const handleInputChange = ({
		target: { name, value },
	}: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
		setFormData((prev) => ({
			...prev,
			[name]: value,
		}));
	};

	const handleInterestsChange = (_: any, value: string[]) => {
		setFormData((prev) => ({
			...prev,
			interests: value,
		}));
	};

	const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();
		try {
			const res = await register(formData).unwrap();
			if (res.ok) {
				showToast({
					type: 'success',
					message: 'Success Registered!',
				});
				navigate(route.LOGIN);
				return;
			}
			throw new Error('failed register');
		} catch {
			showToast({
				type: 'error',
				message: 'Server Error',
			});
		}
	};

	return (
		<Box component="form" onSubmit={handleSubmit}>
			<fieldset disabled={isLoading}>
				<Grid container spacing={2}>
					<Grid size={6}>
						<TextField
							fullWidth
							name="firstName"
							label="First Name"
							value={formData.firstName}
							onChange={handleInputChange}
							margin="normal"
						/>
					</Grid>
					<Grid size={6}>
						<TextField
							fullWidth
							name="lastName"
							label="Last Name"
							value={formData.lastName}
							onChange={handleInputChange}
							margin="normal"
						/>
					</Grid>
				</Grid>
				<Grid container spacing={2}>
					<Grid size={6}>
						<TextField
							fullWidth
							name="username"
							label="Username"
							value={formData.username}
							onChange={handleInputChange}
							margin="normal"
							required
						/>
					</Grid>
					<Grid size={6}>
						<TextField
							fullWidth
							name="password"
							label="Password"
							type="password"
							value={formData.password}
							onChange={handleInputChange}
							margin="normal"
							required
						/>
					</Grid>
				</Grid>
				<Grid container spacing={2}>
					<Grid size={6}>
						<TextField
							fullWidth
							name="birthday"
							label="Birthday"
							type="date"
							slotProps={{ inputLabel: { shrink: true } }}
							value={formData.birthday}
							onChange={handleInputChange}
							margin="normal"
						/>
					</Grid>
					<Grid size={6}>
						<TextField
							fullWidth
							select
							name="gender"
							label="Gender"
							value={formData.gender}
							onChange={handleInputChange}
							margin="normal"
						>
							{Object.values(genders).map((option) => (
								<MenuItem key={option} value={option}>
									{option}
								</MenuItem>
							))}
						</TextField>
					</Grid>
				</Grid>
				<Grid container spacing={2}>
					<Grid size={6}>
						<Autocomplete
							multiple
							freeSolo
							options={['football', 'guitar']}
							value={formData.interests}
							onChange={handleInterestsChange}
							renderInput={(params) => (
								<TextField
									{...params}
									name="interests"
									label="Interests"
									placeholder="Add Interest"
									margin="normal"
								/>
							)}
						/>
					</Grid>
					<Grid size={6}>
						<TextField
							fullWidth
							name="city"
							label="City"
							value={formData.city}
							onChange={handleInputChange}
							margin="normal"
						/>
					</Grid>
				</Grid>

				<Button
					type="submit"
					size="large"
					variant="contained"
					sx={{ mt: 2 }}
					disabled={isLoading}
				>
					Register
				</Button>
			</fieldset>
		</Box>
	);
};
