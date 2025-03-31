import { LLPkgStore } from './parser/types';

class ContentError extends Error {
    constructor(message: string) {
        super(message);
        this.name = 'ContentError';
    }
}

export async function getVersionData(): Promise<LLPkgStore> {
    try {
        const response = await fetch('./llpkgstore.json', {
            method: 'GET',
            headers: {
                'Cache-Control': 'no-cache, no-store, must-revalidate',
            },
        });

        if (
            !response.headers.get('content-type')?.includes('application/json')
        ) {
            throw new ContentError('Invalid content type of llpkgstore.json');
        }

        try {
            const data = await response.json();
            return data;
        } catch (e) {
            if (e instanceof SyntaxError)
                throw new ContentError('Invalid llpkgstore.json');
            else throw new ContentError('Unknown error');
        }
    } catch (error) {
        if (error instanceof ContentError) {
            throw new Error(error.message);
        }
        throw new Error('Failed to fetch data');
    }
}
