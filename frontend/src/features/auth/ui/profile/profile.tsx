import { CircularProgress } from '@mui/material';
import { useEffect } from 'react';
import { useGetProfileMutation } from '../../model/auth.api';
import { useToast } from '../../../../app/ui/toast-provider/toast-provider';
import UserInfo from '../user-info';

export const Profile = () => {
	const [getProfile, { isLoading, data, error }] = useGetProfileMutation();
	const { showToast } = useToast();

	useEffect(() => {
		const loadProfile = async () => {
			try {
				const res = await getProfile().unwrap();
				if (res.ok) {
					return;
				}
				throw new Error('failed get profile');
			} catch {
				showToast({
					type: 'error',
					message: 'Server Error',
				});
			}
		};

		loadProfile();
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, []);

	if (isLoading) {
		return <CircularProgress size={60} />;
	}

	if (error) {
		<>Cant get Profile</>;
	}

	return <UserInfo {...data?.data} />;
};
