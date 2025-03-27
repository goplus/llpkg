import { LLPkgStore } from './types';

/**
 * Pkg name parser for `llpkgstore.json`
 * @param data llpkgstore.json data
 * @param query query for package name
 * @param page current page
 * @param pageSize items per page
 * @returns pkg name list and total count under the query condition
 */
export const pkgnameParser = (
    data?: LLPkgStore,
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

/**
 * Verison mapping parser for `llpkgstore.json`
 * @param data version data only
 * @param queryOrigin query for origin C version
 * @param queryMapped query for mapped Go version
 * @param page current page
 * @param pageSize items per page
 * @param descending is descending order
 * @returns origin C version list and total count under the query condition
 */
export const versionParser = (
    data: LLPkgStore[string],
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
        !isEmpty(queryOrigin) || !isEmpty(queryMapped)
            ? Object.keys(versions).filter((key) => {
                  // Judge whether the key contains the queryOrigin
                  let result =
                      isEmpty(queryOrigin) || key.includes(queryOrigin);

                  // Judge whether the key contains the queryMapped
                  if (!isEmpty(queryMapped) && result) {
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

const isEmpty = (text: string): boolean => {
    return text === '';
};
