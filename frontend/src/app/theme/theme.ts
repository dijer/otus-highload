import { createTheme } from '@mui/material';
import { grey } from '@mui/material/colors';

export const theme = createTheme({
	palette: {
		mode: 'light',
		primary: {
			main: grey[900],
		},
	},
});
