// Single source of truth for VM state. Svelte components read from this.
// WASM Go backend pushes JSON via os.subscribe callback. No Svelte-derived state.

import type { RootVM } from './types';

type Subscriber = (vm: RootVM) => void;

let vm: RootVM | null = null;
const subscribers = new Set<Subscriber>();

export function initOS(): void {
  const tryInit = () => {
    const os = (window as any).os;
    if (!os) {
      setTimeout(tryInit, 50);
      return;
    }

    os.subscribe((json: string) => {
      try {
        vm = JSON.parse(json) as RootVM;
        for (const sub of subscribers) {
          sub(vm);
        }
      } catch (e) {
        console.error('Failed to parse VM JSON:', e);
      }
    });

    if (os.setViewport) {
      os.setViewport(window.innerWidth, window.innerHeight);
    }
  };
  tryInit();
}

export function getVM(): RootVM | null {
  return vm;
}

export function subscribe(fn: Subscriber): () => void {
  subscribers.add(fn);
  if (vm) fn(vm);
  return () => subscribers.delete(fn);
}

export function setViewport(w: number, h: number): void {
  (window as any).os?.setViewport?.(w, h);
}

// Typed accessor for os API
export function os(): any {
  return (window as any).os;
}
