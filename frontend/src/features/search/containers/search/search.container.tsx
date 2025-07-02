import { Stack } from '@mui/material';
import { useLazySearchQuery } from '../../model/search.api';
import SearchForm from '../../ui/search-form';
import { Users } from '../../ui/users';

export const SearchContainer = () => {
	const [search, { isUninitialized, isFetching }] = useLazySearchQuery();

	return (
		<Stack spacing={2}>
			<SearchForm search={search} isLoading={isFetching} />
			<Users isLoading={isFetching} isUninitialized={isUninitialized} />
		</Stack>
	);
};
