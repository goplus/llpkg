import { useEffect, useMemo, useState } from 'react';
import AscendingImg from '@/assets/sortAscending.svg';
import DescendingImg from '@/assets/sortDescending.svg';
import Pagination from '@/components/Pagination';
import { versionParser } from '@/utils/parser';
import { LLPkgStore } from '@/utils/parser/types';
import { setSearchParams } from '@/utils/searchParams';
import { paginationSize } from '@/config/pagination';
import Modal from '../Modal';
import Title from './Title';
import VersionItem from './Items';

interface DetailModalProps {
    modalOpen: boolean;
    setModalOpen: (open: boolean) => void;
    name?: string;
    data?: LLPkgStore;
}

const DetailModal: React.FC<DetailModalProps> = ({
    modalOpen,
    setModalOpen,
    name,
    data,
}) => {
    const [originVersion, setOriginVersion] = useState('');
    const [mappedVersion, setMappedVersion] = useState('');
    const [itemOffset, setItemOffset] = useState(0);
    const [desc, setDesc] = useState(true);
    const [version, setVersion] = useState('latest');
    const pageSize = paginationSize.version;
    const searchResult = useMemo(
        () =>
            data &&
            name &&
            versionParser(
                data[name],
                originVersion,
                mappedVersion,
                itemOffset,
                pageSize,
                desc,
            ),
        // eslint-disable-next-line react-hooks/exhaustive-deps
        [data, originVersion, mappedVersion, desc, name, itemOffset],
    );

    useEffect(() => {
        setItemOffset(0);
        setDesc(true);
        setVersion('latest');
    }, [name]);

    useEffect(() => {
        setItemOffset(0);
    }, [originVersion, mappedVersion]);

    return (
        <Modal
            isOpen={modalOpen}
            onClose={() => {
                setModalOpen(false);
                setSearchParams('pkg', '');
            }}
        >
            <div className="flex h-full flex-col">
                <Title name={name} version={version} setVersion={setVersion} />
                <div className="relative h-full overflow-auto px-4 pb-3">
                    {data && name && searchResult ? (
                        <>
                            <div className="sticky top-1 mt-5 mb-1 flex flex-row justify-around text-base">
                                <div className="mb-2 flex w-full items-center gap-2 overflow-hidden rounded-lg border border-gray-300 bg-white/70 px-2 backdrop-blur-xs sm:w-fit">
                                    <input
                                        onChange={(e) =>
                                            setOriginVersion(e.target.value)
                                        }
                                        className="w-1/2 border-r border-gray-300 py-2 focus-visible:outline-0 sm:w-fit sm:px-1"
                                        placeholder="Original version"
                                    />
                                    <input
                                        onChange={(e) =>
                                            setMappedVersion(e.target.value)
                                        }
                                        className="w-1/2 py-2 focus-visible:outline-0 sm:w-fit sm:px-1"
                                        placeholder="Mapped version"
                                    />
                                </div>
                            </div>
                            <div className="mb-3 text-right">
                                <button
                                    className="btn-a cursor-pointer text-sm text-gray-500"
                                    onClick={() => setDesc(!desc)}
                                >
                                    <img
                                        src={
                                            desc ? DescendingImg : AscendingImg
                                        }
                                        className="inline-block w-3 align-middle"
                                    />
                                    <span className="align-middle">
                                        {desc ? 'Descending' : 'Ascending'}
                                    </span>
                                </button>
                            </div>
                            <div className="flex flex-col gap-3">
                                {searchResult.data.map((ver, index) => {
                                    return (
                                        <VersionItem
                                            key={index}
                                            cver={ver}
                                            gover={data[name].versions[ver]}
                                            setVersion={setVersion}
                                        />
                                    );
                                })}
                            </div>
                            <Pagination
                                className="mt-6"
                                setItemOffset={setItemOffset}
                                itemOffset={itemOffset}
                                itemCount={searchResult.totalCount}
                                pageSize={pageSize}
                            />
                            {searchResult.totalCount === 0 && (
                                <p>No results.</p>
                            )}
                        </>
                    ) : (
                        <p>No data.</p>
                    )}
                </div>
            </div>
        </Modal>
    );
};

export default DetailModal;
