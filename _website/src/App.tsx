import { useEffect, useMemo, useState } from 'react';
import Header from './layout/header';
import List from './components/list/list';
import Search from './components/search';
import { titleParser } from './utils/parser/parser';
import { VersionData } from './utils/parser/types';
import Pagination from './components/pagination';
import { Tooltip } from 'react-tooltip';
import { getVersionData } from './utils/getLLPkgstore';
import './App.css';
import { getSearchParams } from './utils/searchParams';

function App() {
    const [search, setSearch] = useState('');
    const [data, setData] = useState<VersionData>();
    const [itemOffset, setItemOffset] = useState(0);
    const searchQuery = getSearchParams('search');
    const pageSize = 4;
    const searchResult = useMemo(
        () => titleParser(data, search, itemOffset, pageSize),
        [data, search, itemOffset],
    );
    useEffect(() => {
        getVersionData().then((data) => {
            setData(data);
        });
        if (searchQuery) {
            setSearch(searchQuery);
        }
    }, [searchQuery]);
    useEffect(() => {
        setItemOffset(0);
    }, [search]);
    return (
        <>
            <Tooltip id="default-tooltip" />
            <Header />
            <Search
                query={search}
                setSearch={setSearch}
                totalPackages={data ? Object.keys(data).length : 0}
                totalResults={searchResult.totalCount}
            />
            <List data={data} titles={searchResult.data} />
            <Pagination
                itemCount={searchResult.totalCount}
                pageSize={pageSize}
                setItemOffset={setItemOffset}
                itemOffset={itemOffset}
                className="pb-6"
            />
        </>
    );
}

export default App;