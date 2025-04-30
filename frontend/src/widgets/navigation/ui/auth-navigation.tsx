import { Box, Button } from '@mui/material';
import { route } from '../../../shared/constants/routes';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { useAppSelector } from '../../../app/hooks/use-app';
import { useToast } from '../../../app/ui/toast-provider/toast-provider';
import { useLogoutMutation } from '../../../features/auth/model/auth.api';

export const AuthNavigation = () => {
	const isAuth = useAppSelector((state) => state.auth.isAuth);
	const location = useLocation();
	const navigate = useNavigate();
	const { showToast } = useToast();
	const [logout] = useLogoutMutation();

	if (isAuth) {
		const handleLogoutHandler = async () => {
			try {
				const res = await logout().unwrap();
				if (res.ok) {
					navigate(route.LOGIN);
					showToast({
						type: 'success',
						message: 'Succefully logged out',
					});
					return;
				}
				throw new Error('failed logout');
			} catch {
				showToast({
					type: 'error',
					message: 'Server Error',
				});
			}
		};

		return (
			<>
				<Button
					sx={{
						color: 'inherit',
						textDecoration:
							location.pathname === route.PROFILE
								? 'underline'
								: 'inherit',
					}}
					to={route.PROFILE}
					component={Link}
				>
					Profile
				</Button>
				<Button
					sx={{
						color: 'inherit',
					}}
					onClick={handleLogoutHandler}
				>
					Logout
				</Button>
			</>
		);
	}

	return (
		<Box sx={{ display: 'flex', gap: 2 }}>
			<Button
				sx={{
					color: 'inherit',
					textDecoration:
						location.pathname === route.REGISTER
							? 'underline'
							: 'inherit',
				}}
				to={route.REGISTER}
				component={Link}
			>
				Register
			</Button>
			<Button
				sx={{
					color: 'inherit',
					textDecoration:
						location.pathname === route.LOGIN
							? 'underline'
							: 'inherit',
				}}
				to={route.LOGIN}
				component={Link}
			>
				Login
			</Button>
		</Box>
	);
};
