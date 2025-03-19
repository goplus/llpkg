import { useEffect, useMemo, useState } from 'react';
import toast, { Toaster } from 'react-hot-toast';
import { Tooltip } from 'react-tooltip';
import List from '@/components/List';
import Pagination from '@/components/Pagination';
import Search from '@/components/Search';
import Header from '@/layout/Header';
import { getVersionData } from '@/utils/getLLPkgstore';
import { titleParser } from '@/utils/parser';
import { VersionData } from '@/utils/parser/types';
import { getSearchParams } from '@/utils/searchParams';
import './App.css';

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
        getVersionData().then(
            (data) => {
                setData(data);
            },
            (error) => {
                toast.error(error.message);
            },
        );
    }, []);
    useEffect(() => {
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
            <Toaster />
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
