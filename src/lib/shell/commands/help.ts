import type { OutputEntry } from '$lib/model/types';

export function help(): OutputEntry {
  return {
    type: 'text',
    payload: [
      'cd <path>     navigate (supports .., ~, relative, absolute)',
      'ls [path]     list directory contents',
      'ls <glob>     filter with wildcards  e.g. ls *eng',
      'cat <file>    read a record',
      'clear         clear the terminal',
      'theme <name>  set theme: gruvbox-{dark,hard,soft}, one-dark, dracula,',
      '              nord, monokai, solarized-dark, tokyo-night, cyberpunk,',
      '              synthwave, matrix, hotdogstand, campfire, vaporwave,',
      '              fallout, portal, ironman, bioshock, skyrim, metroid,',
      '              zelda, minecraft, deus-ex, mass-effect',
      '',
      'Tab           autocomplete paths and commands',
      '↑ ↓           cycle command history',
    ].join('\n'),
  };
}
