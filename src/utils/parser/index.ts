import { VersionData } from './types';

export const titleParser = (
    data?: VersionData,
    query?: string,
    page: number = 0,
    pageSize: number = 10,
): { data: string[]; totalCount: number } => {
    if (!data) return { data: [], totalCount: 0 };
    query = query?.trim();

    const parsedData = query
        ? Object.keys(data).filter((key) => {
              return key.includes(query);
          })
        : Object.keys(data);

    const startIndex = page * pageSize;
    const endIndex = startIndex + pageSize;

    return {
        data: parsedData.slice(startIndex, endIndex),
        totalCount: parsedData.length,
    };
};

export const versionParser = (
    data: VersionData[string],
    queryOrigin: string = '',
    queryMapped: string = '',
    page: number = 0,
    pageSize: number = 10,
    descending: boolean = false,
): { data: string[]; totalCount: number } => {
    const versions = data.versions;
    if (!versions) return { data: [], totalCount: 0 };
    queryOrigin = queryOrigin?.trim();
    queryMapped = queryMapped?.trim();

    const filteredVersions =
        queryOrigin || queryMapped
            ? Object.keys(versions).filter((key) => {
                  let result = queryOrigin ? key.includes(queryOrigin) : true;
                  if (queryMapped) {
                      const mappedVersions = versions[key];
                      result = mappedVersions.some((ver) =>
                          ver.includes(queryMapped),
                      );
                  }
                  return result;
              })
            : Object.keys(versions);

    if (descending) {
        filteredVersions.reverse();
    }

    const startIndex = page * pageSize;
    const endIndex = startIndex + pageSize;

    return {
        data: filteredVersions.slice(startIndex, endIndex),
        totalCount: filteredVersions.length,
    };
};
