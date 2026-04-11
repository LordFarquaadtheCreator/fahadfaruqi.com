<script lang="ts">
	import type { Component } from 'svelte';
	import IconAlternateEmail from '~icons/material-symbols/alternate-email';
	import IconArrowForwardIos from '~icons/material-symbols/arrow-forward-ios';
	import IconArrowOutward from '~icons/material-symbols/arrow-outward';
	import IconChevronRight from '~icons/material-symbols/chevron-right';
	import IconCode from '~icons/material-symbols/code';
	import IconDatabase from '~icons/material-symbols/database';
	import IconDataset from '~icons/material-symbols/dataset';
	import IconDescription from '~icons/material-symbols/description';
	import IconDownload from '~icons/material-symbols/download';
	import IconFolderOpen from '~icons/material-symbols/folder-open';
	import IconGridView from '~icons/material-symbols/grid-view';
	import IconHub from '~icons/material-symbols/hub';
	import IconMemory from '~icons/material-symbols/memory';
	import IconPerson from '~icons/material-symbols/person';
	import IconPersonSearch from '~icons/material-symbols/person-search';
	import IconQrCode from '~icons/material-symbols/qr-code';
	import IconRadar from '~icons/material-symbols/radar';
	import IconSatelliteAlt from '~icons/material-symbols/satellite-alt';
	import IconSchool from '~icons/material-symbols/school';
	import IconSend from '~icons/material-symbols/send';
	import IconSettingsInputComponent from '~icons/material-symbols/settings-input-component';
	import IconSignalCellular4Bar from '~icons/material-symbols/signal-cellular-4-bar';
	import IconTerminal from '~icons/material-symbols/terminal';

	const icons = {
		alternate_email: IconAlternateEmail,
		arrow_forward_ios: IconArrowForwardIos,
		arrow_outward: IconArrowOutward,
		chevron_right: IconChevronRight,
		code: IconCode,
		database: IconDatabase,
		dataset: IconDataset,
		description: IconDescription,
		download: IconDownload,
		folder_open: IconFolderOpen,
		grid_view: IconGridView,
		hub: IconHub,
		memory: IconMemory,
		person: IconPerson,
		person_search: IconPersonSearch,
		qr_code: IconQrCode,
		radar: IconRadar,
		satellite_alt: IconSatelliteAlt,
		school: IconSchool,
		send: IconSend,
		settings_input_component: IconSettingsInputComponent,
		signal_cellular_4_bar: IconSignalCellular4Bar,
		terminal: IconTerminal
	} satisfies Record<string, Component>;

	type IconName = keyof typeof icons;

	let {
		name,
		class: className = '',
		style = '',
		ariaLabel
	} = $props<{
		name: string;
		class?: string;
		style?: string;
		ariaLabel?: string;
	}>();

	const IconComponent = $derived(icons[name as IconName]);
</script>

{#if IconComponent}
	<IconComponent
		class={className}
		style={style}
		aria-label={ariaLabel}
		aria-hidden={ariaLabel ? undefined : 'true'}
	/>
{/if}
