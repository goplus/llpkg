import { useMemo } from 'react';
import { LLPkgStore } from '@/utils/parser/types';
import { setSearchParams } from '@/utils/searchParams';
import conanImg from '@/assets/conan.svg';
import githubImg from '@/assets/github.svg';
import { urls } from '@/config/urls';
import { versionParser } from '@/utils/parser';

interface ItemProps {
    name: string;
    data: LLPkgStore[string];
    index: number;
    setInfo: (data: string) => void;
    setModalOpen: (open: boolean) => void;
}

const Item: React.FC<ItemProps> = ({ name, data, setInfo, setModalOpen }) => {
    const versions = useMemo(() => {
        return versionParser(data, '', '', 0, 2);
    }, [data]);
    const openModal = () => {
        setInfo(name);
        setSearchParams('pkg', name);
        setModalOpen(true);
    };
    return (
        <button
            className="flex min-h-32 cursor-pointer flex-col items-start gap-2 overflow-clip rounded-xl border border-gray-300 bg-gray-50/60 px-8 py-4 text-left shadow-md transition-all duration-300 hover:bg-white hover:shadow-xl"
            onClick={openModal}
        >
            <div className="flex w-full flex-row items-center gap-2">
                <span className="w-full cursor-pointer text-2xl font-bold text-wrap text-gray-900">
                    {name}
                    <span className="ml-3 self-end text-sm font-normal text-gray-500">
                        {versions.totalCount} versions in total
                    </span>
                </span>
                <div
                    className="ml-auto hidden flex-row gap-2 sm:flex"
                    onClick={(e) => e.stopPropagation()}
                >
                    <a
                        href={`${urls.llpkg}/${name}`}
                        target="_blank"
                        rel="noreferrer"
                        className="btn-icon inline-block px-2"
                    >
                        <img src={githubImg} />
                    </a>
                    <a
                        href={`${urls.conan}/${name}`}
                        target="_blank"
                        rel="noreferrer"
                        className="btn-icon inline-block px-2"
                    >
                        <img src={conanImg} />
                    </a>
                </div>
            </div>
            <div className="w-full">
                {versions.data.map((cver, index) => {
                    const gover = data.versions[cver];
                    return (
                        <div
                            key={index}
                            className="flex flex-row items-center gap-2 overflow-hidden text-nowrap overflow-ellipsis whitespace-nowrap"
                        >
                            <span className="w-fit max-w-full min-w-16 overflow-hidden text-left text-lg leading-9 font-bold text-nowrap overflow-ellipsis whitespace-nowrap">
                                <span
                                    data-tooltip-id="default-tooltip"
                                    data-tooltip-content={cver}
                                    data-tooltip-place="top"
                                >
                                    {cver}
                                </span>
                            </span>
                            <span
                                className="w-fit max-w-full overflow-hidden leading-9 text-nowrap overflow-ellipsis whitespace-nowrap"
                                data-tooltip-id="default-tooltip"
                                data-tooltip-content={gover.join(' / ')}
                                data-tooltip-place="top"
                            >
                                <span>{gover.join(' / ')}</span>
                            </span>
                        </div>
                    );
                })}
            </div>
        </button>
    );
};

export default Item;
