import { ChangeEvent, useState } from 'react';
import { Box, Button, Grid, TextField } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { IUserWithPassword } from '../../model/user.model';
import { useLoginMutation } from '../../model/auth.api';
import { useToast } from '../../../../app/ui/toast-provider/toast-provider';
import { route } from '../../../../shared/constants/routes';

export const LoginForm = () => {
	const [formData, setFormData] = useState<
		Partial<Partial<IUserWithPassword>>
	>({
		username: '',
		password: '',
	});
	const [login, { isLoading }] = useLoginMutation();
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

	const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();
		try {
			const res = await login(formData).unwrap();
			if (res.ok) {
				showToast({
					type: 'success',
					message: 'Success Login!',
				});
				navigate(route.PROFILE);
				return;
			}
			throw new Error('failed login');
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

				<Button
					type="submit"
					size="large"
					variant="contained"
					sx={{ mt: 2 }}
					disabled={isLoading}
				>
					Login
				</Button>
			</fieldset>
		</Box>
	);
};
