import { configureStore } from '@reduxjs/toolkit';
import { authReducer } from '../../features/auth/model/auth.slice';
import { authApi } from '../../features/auth/model/auth.api';

export const store = configureStore({
	reducer: {
		auth: authReducer,
		[authApi.reducerPath]: authApi.reducer,
	},
	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware().concat(authApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
