import { Alert, AlertColor, Snackbar } from '@mui/material';
import { createContext, useContext, useState } from 'react';

interface IToastProps {
	message: string;
	type?: AlertColor;
}

interface IToastContext {
	showToast(props: IToastProps): void;
}

const ToastContext = createContext<IToastContext>({} as IToastContext);

export const ToastProvider: React.FC<React.PropsWithChildren> = ({
	children,
}) => {
	const [toast, setToast] = useState({
		message: '',
		type: 'info',
		isOpen: false,
	});
	const showToast = ({ message, type = 'info' }: IToastProps) => {
		setToast({
			message,
			type,
			isOpen: true,
		});
	};

	const setOpen = (isOpen: boolean) => {
		setToast((prev) => ({
			...prev,
			isOpen,
		}));
	};

	return (
		<ToastContext.Provider value={{ showToast }}>
			{children}
			<Snackbar
				open={toast.isOpen}
				autoHideDuration={3000}
				onClose={() => setOpen(false)}
				anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
			>
				<Alert
					onClose={() => setOpen(false)}
					severity={toast.type as AlertColor}
					sx={{ width: '100%' }}
				>
					{toast.message}
				</Alert>
			</Snackbar>
		</ToastContext.Provider>
	);
};

export const useToast = () => {
	return useContext(ToastContext);
};
