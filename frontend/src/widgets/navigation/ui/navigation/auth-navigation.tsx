import { Box } from '@mui/material';
import { route } from '../../../../shared/constants/routes';
import { useNavigate } from 'react-router-dom';
import { useAppSelector } from '../../../../app/hooks/use-app';
import { useToast } from '../../../../app/ui/toast-provider/toast-provider';
import { useLogoutMutation } from '../../../../features/auth/model/auth.api';
import NavLink from '../nav-link';

export const AuthNavigation = () => {
	const isAuth = useAppSelector((state) => state.auth.isAuth);
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
			<Box sx={{ display: 'flex', gap: 2 }}>
				<NavLink to={route.PROFILE}>Profile</NavLink>
				<NavLink onClick={handleLogoutHandler}>Logout</NavLink>
			</Box>
		);
	}

	return (
		<Box sx={{ display: 'flex', gap: 2 }}>
			<NavLink to={route.REGISTER}>Register</NavLink>
			<NavLink to={route.LOGIN}>Login</NavLink>
		</Box>
	);
};
