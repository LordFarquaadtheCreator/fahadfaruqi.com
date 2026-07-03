# src/lib/finder/

Read-only VM mirror. Svelte store subscribes to Go and renders.

## Files

- `finderStore.svelte.ts` — `$state` VM holder + `initFinder()` + `window.os` TS declarations

## Rule

Zero business logic. Zero `$derived` that computes view models.
Only trivial prop drilling and event forwarding.
