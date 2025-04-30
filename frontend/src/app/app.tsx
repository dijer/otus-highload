import { Route, Routes } from 'react-router-dom';
import HomePage from '../pages/home';
import NotFoundPage from '../pages/not-found';
import LoginPage from '../pages/login';
import RegisterPage from '../pages/register';
import Navigation from '../widgets/navigation';
import { route } from '../shared/constants/routes';
import UserPage from '../pages/user';
import ProfilePage from '../pages/profile';

export const App = () => {
	return (
		<>
			<Navigation />
			<Routes>
				<Route path={route.HOME} element={<HomePage />} />
				<Route path={route.REGISTER} element={<RegisterPage />} />
				<Route path={route.LOGIN} element={<LoginPage />} />
				<Route path={`${route.USER}/:id`} element={<UserPage />} />
				<Route path={route.PROFILE} element={<ProfilePage />} />
				<Route path="*" element={<NotFoundPage />} />
			</Routes>
		</>
	);
};
