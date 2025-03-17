import { VersionData } from '@/utils/parser/types';

interface VersionItemProps {
    ver: VersionData[string]['versions'][number];
    setVersion: (tag: string) => void;
}

const VersionItem: React.FC<VersionItemProps> = ({ ver, setVersion }) => {
    return (
        <div className="flex flex-col items-center gap-4 rounded-lg bg-gray-50 px-4 py-3 sm:flex-row">
            <span className="overflow-auto text-left text-lg font-bold sm:min-w-16">
                {ver.c}
            </span>

            <span className="flex flex-row flex-wrap gap-2 overflow-auto sm:flex-nowrap">
                {ver.go.map((ver, index) => (
                    <VersionTag key={index} tag={ver} onClick={setVersion} />
                ))}
            </span>
            <span className="ml-auto w-full self-center border-t border-gray-300 pt-2 pl-0 sm:w-auto sm:min-w-20 sm:border-t-0 sm:border-l sm:pl-2 md:pt-0">
                <VersionTag
                    tag={ver.go[ver.go.length - 1]}
                    onClick={setVersion}
                />
            </span>
        </div>
    );
};

export default VersionItem;

interface VersionTagProps {
    tag: string;
    className?: string;
    onClick?: (tag: string) => void;
}

const VersionTag: React.FC<VersionTagProps> = ({ tag, className, onClick }) => {
    return (
        <button
            onClick={() => onClick && onClick(tag)}
            className={[
                'cursor-pointer rounded-full bg-gradient-to-r from-blue-50 to-purple-50 px-3 py-1 text-sm text-blue-600 transition-colors duration-300 hover:from-blue-100 hover:to-purple-100',
                className,
            ].join(' ')}
        >
            {tag}
        </button>
    );
};
