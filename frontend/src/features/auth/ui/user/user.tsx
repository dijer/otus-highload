import { CircularProgress } from '@mui/material';
import { useLazyGetUserQuery } from '../../model/auth.api';
import { useToast } from '../../../../app/ui/toast-provider/toast-provider';
import { useEffect } from 'react';
import UserInfo from '../user-info';

interface IUser {
	userId: number;
}

export const User = ({ userId }: IUser) => {
	const [getUser, { isLoading, data, error }] = useLazyGetUserQuery();
	const { showToast } = useToast();

	useEffect(() => {
		const getUserById = async () => {
			try {
				const res = await getUser(userId).unwrap();
				if (res.ok) {
					return;
				}
				throw new Error('failed get user');
			} catch {
				showToast({
					type: 'error',
					message: 'Server Error',
				});
			}
		};

		getUserById();
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [userId]);

	if (isLoading) {
		return <CircularProgress size={60} />;
	}

	if (error) {
		return <>Cant get User</>;
	}

	return <UserInfo {...data?.data} />;
};
