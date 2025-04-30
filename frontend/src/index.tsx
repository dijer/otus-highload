import React from 'react';
import ReactDOM from 'react-dom/client';
import { Provider } from 'react-redux';
import { CssBaseline, ThemeProvider } from '@mui/material';
import { store } from './app/store';
import { BrowserRouter } from 'react-router-dom';
import App from './app';

import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import theme from './app/theme';
import ToastProvider from './app/ui/toast-provider';
import AuthProvider from './features/auth/providers/auth';

const root = ReactDOM.createRoot(
	document.getElementById('root') as HTMLElement
);
root.render(
	<React.StrictMode>
		<ThemeProvider theme={theme}>
			<CssBaseline />
			<ToastProvider>
				<BrowserRouter>
					<Provider store={store}>
						<AuthProvider>
							<App />
						</AuthProvider>
					</Provider>
				</BrowserRouter>
			</ToastProvider>
		</ThemeProvider>
	</React.StrictMode>
);
