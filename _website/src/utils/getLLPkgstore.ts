import { VersionData } from './parser/types';

export async function getVersionData(): Promise<VersionData> {
    const response = await fetch('./llpkgstore.json', {
        method: 'GET',
        headers: {
            'Cache-Control': 'no-cache',
        },
    });
    const data = await response.json();
    return data;
}
