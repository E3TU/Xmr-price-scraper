<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import '@fontsource-variable/lexend/wght.css';

	let { children } = $props();

	import '../app.css'
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children()}
