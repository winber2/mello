export function getURLParams(url) {
  const params = {};
  if (!url) return params;

  const kvPairs = url.slice(1).split('=');
  for (let i = 0; i < kvPairs.length; i += 2) {
    const key = kvPairs[i];
    if (i % 2 === 0 && kvPairs[i + 1]) {
      params[key] = kvPairs[i + 1];
    }
  }

  return params;
}
