import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { API_URL } from '../../../shared/config/env';

export const searchApi = createApi({
	reducerPath: 'searchApi',
	baseQuery: fetchBaseQuery({ baseUrl: API_URL }),
	endpoints: (builder) => ({
		search: builder.query({
			query: ({ firstName, lastName }) => ({
				url: '/user/search',
				params: {
					firstname: firstName,
					lastname: lastName,
				},
			}),
		}),
	}),
});

export const { useLazySearchQuery } = searchApi;
