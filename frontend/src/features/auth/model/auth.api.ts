import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { IUserWithPassword } from './user.model';
import { API_URL } from '../../../shared/config/env';
import { IServerResponse } from '../../../shared/types/server';

export const authApi = createApi({
	reducerPath: 'authApi',
	baseQuery: fetchBaseQuery({ baseUrl: API_URL }),
	endpoints: (builder) => ({
		register: builder.mutation<IServerResponse, Partial<IUserWithPassword>>(
			{
				query: (credentials) => ({
					url: '/user/register',
					method: 'POST',
					body: credentials,
					credentials: 'include',
				}),
			}
		),
		login: builder.mutation<IServerResponse, Partial<IUserWithPassword>>({
			query: (credentials) => ({
				url: '/login',
				method: 'POST',
				body: credentials,
				credentials: 'include',
			}),
		}),
		getUser: builder.query({
			query: (userId: number) => `/user/get/${userId}`,
		}),
		logout: builder.mutation<IServerResponse, void>({
			query: () => ({
				url: '/user/logout',
				method: 'POST',
				credentials: 'include',
			}),
		}),
		checkAuth: builder.mutation<IServerResponse, void>({
			query: () => ({
				url: '/user/check',
				method: 'POST',
				credentials: 'include',
			}),
		}),
		getProfile: builder.mutation<IServerResponse, void>({
			query: () => ({
				url: '/user/profile',
				method: 'POST',
				credentials: 'include',
			}),
		}),
	}),
});

export const {
	useRegisterMutation,
	useLoginMutation,
	useLazyGetUserQuery,
	useLogoutMutation,
	useCheckAuthMutation,
	useGetProfileMutation,
} = authApi;
