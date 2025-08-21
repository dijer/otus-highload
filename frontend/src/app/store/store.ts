import { configureStore } from '@reduxjs/toolkit';
import { authReducer } from '../../features/auth/model/auth.slice';
import { authApi } from '../../features/auth/model/auth.api';
import { searchApi } from '../../features/search/model/search.api';
import { searchReducer } from '../../features/search/model/search.slice';

export const store = configureStore({
	reducer: {
		auth: authReducer,
		[authApi.reducerPath]: authApi.reducer,

		search: searchReducer,
		[searchApi.reducerPath]: searchApi.reducer,
	},
	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware().concat(authApi.middleware, searchApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
