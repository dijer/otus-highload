import { createSlice } from '@reduxjs/toolkit';
import { searchApi } from './search.api';
import { IUser } from '../../../entities/user';

interface ISearchState {
	total: number;
	users: IUser[];
}

const initialState: ISearchState = {
	total: 0,
	users: [],
};

const searchSlice = createSlice({
	name: 'search',
	initialState,
	reducers: {},
	extraReducers: (builder) => {
		builder.addMatcher(
			searchApi.endpoints.search.matchFulfilled,
			(state, action) => {
				state.total = action.payload.data.count;
				state.users = action.payload.data.users;
			}
		);
		builder.addMatcher(searchApi.endpoints.search.matchPending, (state) => {
			state.total = 0;
			state.users = [];
		});
	},
});

export const searchReducer = searchSlice.reducer;
export type SearchAction = ReturnType<
	(typeof searchSlice.actions)[keyof typeof searchSlice.actions]
>;
