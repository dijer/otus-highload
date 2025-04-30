import { createSlice } from '@reduxjs/toolkit';
import { authApi } from './auth.api';
import { AUTH_LOCALSTORAGE_KEY } from '../constants';

interface IAuthState {
	isAuth: boolean;
}

const initialState: IAuthState = {
	isAuth: localStorage.getItem(AUTH_LOCALSTORAGE_KEY) === 'true',
};

const authSlice = createSlice({
	name: 'auth',
	initialState,
	reducers: {},
	extraReducers: (builder) => {
		builder.addMatcher(authApi.endpoints.login.matchFulfilled, (state) => {
			state.isAuth = true;
			localStorage.setItem(AUTH_LOCALSTORAGE_KEY, 'true');
		});
		builder.addMatcher(authApi.endpoints.logout.matchFulfilled, (state) => {
			state.isAuth = false;
			localStorage.removeItem(AUTH_LOCALSTORAGE_KEY);
		});
		builder.addMatcher(
			authApi.endpoints.checkAuth.matchRejected,
			(state) => {
				state.isAuth = false;
				localStorage.removeItem(AUTH_LOCALSTORAGE_KEY);
			}
		);
	},
});

export const authReducer = authSlice.reducer;
export type AuthAction = ReturnType<
	(typeof authSlice.actions)[keyof typeof authSlice.actions]
>;
