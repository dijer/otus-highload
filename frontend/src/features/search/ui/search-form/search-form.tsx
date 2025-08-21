import { ChangeEvent, useState } from 'react';
import { Box, Button, Grid, TextField } from '@mui/material';
import { useToast } from '../../../../app/ui/toast-provider/toast-provider';
import { IUser } from '../../../../entities/user';

type TProps = {
	search: Function;
	isLoading: boolean;
};

export const SearchForm = ({ search, isLoading }: TProps) => {
	const [formData, setFormData] = useState<Partial<IUser>>({
		firstName: '',
		lastName: '',
	});

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
			const res = await search(formData).unwrap();
			if (res.ok) {
				showToast({
					type: 'success',
					message: 'Success search',
				});
				return;
			}
			throw new Error('failed search');
		} catch (e: any) {
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
							required
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
					Search
				</Button>
			</fieldset>
		</Box>
	);
};
