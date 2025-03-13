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

function App() {
    const [search, setSearch] = useState('');
    const [data, setData] = useState<VersionData>();
    const [itemOffset, setItemOffset] = useState(0);
    const pageSize = 4;
    const searchResult = useMemo(
        () => titleParser(data, search, itemOffset, pageSize),
        [data, search, itemOffset],
    );
    useEffect(() => {
        getVersionData().then((data) => {
            setData(data);
        });
    }, []);
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
                dataNumber={data ? Object.keys(data).length : 0}
                resultNumber={searchResult.totalCount}
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