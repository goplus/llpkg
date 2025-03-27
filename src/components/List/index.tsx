import { useEffect, useState } from 'react';
import toast from 'react-hot-toast';
import { LLPkgStore } from '@/utils/parser/types';
import { getSearchParams, setSearchParams } from '@/utils/searchParams';
import DetailModal from '../Detail';
import Item from './Items';

interface ListProps {
    data?: LLPkgStore;
    titles: string[];
}

const List: React.FC<ListProps> = ({ data, titles }) => {
    const [modalOpen, setModalOpen] = useState(false);
    const [name, setDetailName] = useState<string>();
    const pkgQuery = getSearchParams('pkg');
    useEffect(() => {
        if (data && pkgQuery && pkgQuery in data) {
            setDetailName(pkgQuery);
            setModalOpen(true);
        } else if (data && pkgQuery && !(pkgQuery in data)) {
            toast.error(`Package "${pkgQuery}" from URL not found`);
            setSearchParams('pkg', '');
        }
    }, [pkgQuery, data]);
    return (
        <>
            <DetailModal
                modalOpen={modalOpen}
                setModalOpen={setModalOpen}
                data={data}
                name={name}
            />
            {data ? (
                <div className="grid grid-cols-1 gap-4 px-8 py-8 text-gray-600 md:px-32 md:py-16 lg:grid-cols-2">
                    {titles.map((name, index) => {
                        return (
                            <Item
                                key={name}
                                name={name}
                                data={data[name]}
                                index={index}
                                setInfo={setDetailName}
                                setModalOpen={setModalOpen}
                            />
                        );
                    })}
                </div>
            ) : (
                <p>Loading...</p>
            )}
        </>
    );
};

export default List;
