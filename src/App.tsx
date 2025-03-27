import { useEffect, useMemo, useState } from 'react';
import toast, { Toaster } from 'react-hot-toast';
import { Tooltip } from 'react-tooltip';
import List from '@/components/List';
import Pagination from '@/components/Pagination';
import Search from '@/components/Search';
import Header from '@/layout/Header';
import { getVersionData } from '@/utils/getLLPkgstore';
import { pkgnameParser } from '@/utils/parser';
import { LLPkgStore } from '@/utils/parser/types';
import { getSearchParams } from '@/utils/searchParams';
import { paginationSize } from '@/config/pagination';
import './App.css';

function App() {
    const [search, setSearch] = useState('');
    const [data, setData] = useState<LLPkgStore>();
    const [itemOffset, setItemOffset] = useState(0);
    const searchQuery = getSearchParams('search');
    const pageSize = paginationSize.pkg;
    const searchResult = useMemo(
        () => pkgnameParser(data, search, itemOffset, pageSize),
        // eslint-disable-next-line react-hooks/exhaustive-deps
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
                className="pb-12"
            />
        </>
    );
}

export default App;
