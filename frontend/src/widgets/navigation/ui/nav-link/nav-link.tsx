import { Link } from '@mui/material';
import { MouseEventHandler } from 'react';
import { useLocation } from 'react-router-dom';

interface IProps {
	to?: string;
	onClick?: MouseEventHandler;
}

export const NavLink: React.FC<React.PropsWithChildren<IProps>> = ({
	to,
	onClick,
	children,
}) => {
	const location = useLocation();
	const isActive = location.pathname === to;

	return (
		<Link
			sx={{
				color: 'inherit',
				cursor: 'pointer',
				textDecoration: isActive ? 'underline' : 'inherit',
			}}
			href={to}
			onClick={onClick}
		>
			{children}
		</Link>
	);
};
