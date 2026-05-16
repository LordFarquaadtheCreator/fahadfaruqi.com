/**
 * Match a glob pattern against a list of candidates.
 * Supports * wildcard only. Case-insensitive.
 *
 * Examples:
 *   match('exp*',    ['experience','skills'])   → ['experience']
 *   match('*eng',    ['senior-eng','mid-eng'])  → ['senior-eng','mid-eng']
 *   match('mid*eng', ['mid-eng','senior-eng'])  → ['mid-eng']
 *   match('*',       ['a','b','c'])             → ['a','b','c']
 */
export function match(pattern: string, candidates: string[]): string[] {
  if (!pattern.includes('*')) {
    // Exact prefix match as fallback — makes 'exp' match 'experience'
    const lower = pattern.toLowerCase();
    return candidates.filter(c => c.toLowerCase().startsWith(lower));
  }

  // Convert glob to regex: escape everything, then restore *
  const escaped = pattern
    .toLowerCase()
    .replace(/[.+^${}()|[\]\\]/g, '\\$&')
    .replace(/\*/g, '.*');

  const re = new RegExp(`^${escaped}$`);
  return candidates.filter(c => re.test(c.toLowerCase()));
}

/**
 * Return the longest common prefix of an array of strings.
 * Used to fill in Tab completion when there are multiple candidates.
 */
export function longestCommonPrefix(strings: string[]): string {
  if (strings.length === 0) return '';
  if (strings.length === 1) return strings[0];

  let prefix = strings[0];
  for (let i = 1; i < strings.length; i++) {
    while (!strings[i].startsWith(prefix)) {
      prefix = prefix.slice(0, -1);
      if (!prefix) return '';
    }
  }
  return prefix;
}
