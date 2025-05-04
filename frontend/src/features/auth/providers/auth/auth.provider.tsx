import { useEffect, useState } from 'react';
import { useAppSelector } from '../../../../app/hooks/use-app';
import { useCheckAuthMutation } from '../../model/auth.api';
import { useNavigate } from 'react-router-dom';
import { route } from '../../../../shared/constants/routes';

export const AuthProvider: React.FC<React.PropsWithChildren> = ({
	children,
}) => {
	const isAuth = useAppSelector((state) => state.auth.isAuth);
	const [checkAuth] = useCheckAuthMutation();
	const [isChecked, setChecked] = useState(false);
	const navigate = useNavigate();

	useEffect(() => {
		const check = async () => {
			try {
				const res = await checkAuth().unwrap();
				if (res.ok) {
					setChecked(true);
					return;
				}
				throw new Error('can check user');
			} catch {
				navigate(route.LOGIN);
			}
		};
		if (isAuth && !isChecked) {
			check();
		}
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, []);

	return <>{children}</>;
};
