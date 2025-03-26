export interface VersionData {
    [packageName: string]: {
        versions: {
            [version: string]: string[];
        };
    };
}
