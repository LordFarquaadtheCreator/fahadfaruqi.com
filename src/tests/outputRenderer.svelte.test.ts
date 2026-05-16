import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import OutputRenderer from '$lib/components/OutputRenderer.svelte';

describe('OutputRenderer', () => {
  it('renders text output', () => {
    const { container } = render(OutputRenderer, {
      props: { output: { type: 'text', payload: 'hello world' } }
    });
    expect(container.textContent).toContain('hello world');
  });

  it('renders error output with red styling', () => {
    const { container } = render(OutputRenderer, {
      props: { output: { type: 'error', payload: 'not found' } }
    });
    const span = container.querySelector('.out-error');
    expect(span).toBeTruthy();
    expect(span!.textContent).toContain('not found');
  });

  it('renders success output', () => {
    const { container } = render(OutputRenderer, {
      props: { output: { type: 'success', payload: 'done' } }
    });
    expect(container.querySelector('.out-success')).toBeTruthy();
  });

  it('renders list output with items', () => {
    const { container } = render(OutputRenderer, {
      props: {
        output: {
          type: 'list',
          payload: [
            { id: 'a', primary: 'Item A', meta: '2024' },
            { id: 'b', primary: 'Item B', secondary: 'desc' }
          ]
        }
      }
    });
    expect(container.querySelectorAll('.list-row').length).toBe(2);
    expect(container.textContent).toContain('Item A');
    expect(container.textContent).toContain('2024');
    expect(container.textContent).toContain('desc');
  });
});
