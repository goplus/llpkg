/** Structure of the package.json file */
export interface LLPkgStore {
    [packageName: string]: {
        versions: {
            /** C version: Go version[] */
            [cVersion: string]: string[];
        };
    };
}
