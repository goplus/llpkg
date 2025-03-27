/**
 * to pluralize a word
 * @param word original word without "s"
 * @param count the item count
 * @returns word with "s" if count > 1
 */
export function pluralize(word: string, count: number): string {
    return count === 1 ? word : `${word}s`;
}
