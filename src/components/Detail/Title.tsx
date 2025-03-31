import { useMemo, useRef } from 'react';
import toast from 'react-hot-toast';
import { CSSTransition } from 'react-transition-group';
import clipboardImg from '@/assets/clipboard.svg';
import conanImg from '@/assets/conan.svg';
import githubImg from '@/assets/github.svg';
import rollbackImg from '@/assets/rollback.svg';
import { urls } from '@/config/urls';

interface TitleProps {
    name?: string;
    latestVersion?: string;
    version: string;
    setVersion: (tag: string) => void;
}

const Title: React.FC<TitleProps> = ({
    name,
    version,
    latestVersion,
    setVersion,
}) => {
    const rollbackRef = useRef<HTMLButtonElement>(null);
    const majorVersion = useMemo(() => {
        const currentVer = version === 'latest' ? latestVersion : version;
        return parseInt(
            currentVer?.replace(/[^0-9]/, '').split('.')[0] ?? '1',
            10,
        );
    }, [version, latestVersion]);
    const copyToClipboard = (text: string) => {
        if (navigator.clipboard)
            navigator.clipboard.writeText(text).then(
                () => {
                    toast.success('Copied to clipboard');
                },
                () => {
                    toast.error('Failed to copy');
                },
            );
        else toast.error('Failed to copy');
    };
    return (
        <div className="flex flex-col items-start gap-4 border-b border-gray-300 px-4 pb-3 sm:h-1/6 sm:flex-row">
            <div className="flex flex-col text-left">
                <h1 className="px-2 text-3xl font-bold">{name}</h1>
                <button
                    className="btn-button cursor-pointer px-2 py-0.5 text-xs font-light text-gray-400 hover:text-gray-500"
                    onClick={() =>
                        copyToClipboard(
                            `${urls.modulePath}/${name}${majorVersion > 1 ? `/v${majorVersion}` : ''}`,
                        )
                    }
                >
                    {urls.modulePath}/{name}
                    {majorVersion > 1 ? `/v${majorVersion}` : ''}
                </button>
            </div>
            <div className="self-center sm:ml-auto">
                <div className="flex flex-row items-center overflow-hidden rounded-lg border border-gray-300">
                    <p className="px-4 py-2 text-gray-700">
                        llgo get {name}@{version}
                    </p>
                    <CSSTransition
                        in={version !== 'latest'}
                        timeout={300}
                        classNames={{
                            enter: 'opacity-0',
                            enterActive:
                                'opacity-100 transition-all duration-300',
                            exitActive: 'opacity-0 transition-all duration-300',
                        }}
                        unmountOnExit
                        nodeRef={rollbackRef}
                    >
                        <button
                            className="btn-button mr-2 h-fit cursor-pointer p-1.5 text-white opacity-0 transition duration-200 hover:scale-110 active:scale-90"
                            onClick={() => setVersion('latest')}
                            ref={rollbackRef}
                        >
                            <img src={rollbackImg} className="h-4 w-4" />
                        </button>
                    </CSSTransition>
                    <button
                        className="btn-button mr-2 h-fit cursor-pointer p-1.5 text-white transition duration-200 hover:scale-110 active:scale-90"
                        onClick={() =>
                            copyToClipboard(`llgo get ${name}@${version}`)
                        }
                    >
                        <img src={clipboardImg} className="h-4 w-4" />
                    </button>
                </div>
                <div className="mt-2 text-center sm:text-right">
                    <a
                        href={`${urls.llpkg}/${name}`}
                        target="_blank"
                        rel="noreferrer"
                        className="btn-icon mx-0.5 inline-block"
                    >
                        <img className="mx-1 h-4 w-4" src={githubImg} />
                    </a>
                    <a
                        href={`${urls.conan}/${name}`}
                        target="_blank"
                        rel="noreferrer"
                        className="btn-icon mx-0.5 inline-block"
                    >
                        <img className="mx-1 h-4 w-4" src={conanImg} />
                    </a>
                </div>
            </div>
        </div>
    );
};

export default Title;
