export function setSearchParams(param: string, value: string) {
    const url = new URL(window.location.toString());
    if (value === '') url.searchParams.delete(param);
    else url.searchParams.set(param, value);
    history.replaceState(null, '', url);
}

export function getSearchParams(param: string): string | null {
    const url = new URL(window.location.toString());
    return url.searchParams.get(param);
}
