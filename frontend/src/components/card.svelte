<script lang="ts">
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';

	const apiEndpoint: string = 'http://localhost:8080/price';

	let price = $state(0);
	let lastUpdated = $state('');

	async function getPrice() {
		const response = await fetch(apiEndpoint);

		const data = await response.json();

		price = data.price;
		lastUpdated = data.lastUpdated
			? new Date(data.lastUpdated).toLocaleString()
			: new Date().toLocaleString();
	}

	onMount(() => {
		getPrice();
	});
</script>

<div class="card-container">
	<h1 class="card-heading">XMR Price Tracker</h1>
	<div class="price-card">
		<div class="price-wrapper">
			<Icon icon="logos:monero" width="100" height="100" />
			<p class="price">{price}</p>
		</div>
		<p class="change">XMR/EUR</p>
		<p class="last-updated">Last updated: {lastUpdated || '—'}</p>
		<button class="refresh-btn" onclick={getPrice}>Refresh Price</button>
	</div>
</div>

<style>
	.card-container {
		display: flex;
		align-items: center;
		flex-direction: column;
		gap: clamp(2rem, 8vw, 10rem);
		width: 100%;
		padding: 0 1rem;
		box-sizing: border-box;
	}
	.card-heading {
		color: var(--orange);
		margin-top: clamp(2rem, 10vh, 5rem);
		font-size: clamp(2rem, 8vw, 3rem);
		text-shadow: 1px 10px 20px var(--orange);
		text-align: center;
	}
	.price-card {
		display: flex;
		flex-direction: column;
		align-items: center;
		width: min(40rem, 100%);
		min-height: 25rem;
		padding: 1rem;
		box-sizing: border-box;
		box-shadow:
			0px -3px 0px var(--orange),
			0px 3px 0px var(--orange),
			rgba(0, 0, 0, 0.56) 0px 22px 70px 4px;
		border-radius: 12px;
		background-color: var(--secondary-gray);
		color: var(--text-primary);
	}
	.price-wrapper {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: clamp(0.5rem, 2vw, 1rem);
		padding-top: 1rem;
		flex-wrap: wrap;
	}
	:global(.price-wrapper svg) {
		width: clamp(60px, 15vw, 100px);
		height: clamp(60px, 15vw, 100px);
	}
	.price {
		font-size: clamp(3rem, 14vw, 7rem);
	}
	.change {
		font-size: clamp(1.25rem, 5vw, 2rem);
	}
	.last-updated {
		font-size: clamp(1rem, 4vw, 1.5rem);
		text-align: center;
	}
	.refresh-btn {
		padding: clamp(0.75rem, 3vw, 1rem) clamp(3rem, 25vw, 12rem);
		color: var(--white);
		font-size: clamp(1rem, 4vw, 1.35rem);
		font-weight: bold;
		font-family: inherit;
		border: none;
		border-radius: 12px;
		margin-top: 2rem;
		cursor: pointer;
		background: var(--orange);
		box-shadow:
			inset 0 1px 0 rgba(255, 255, 255, 0.35),
			0 6px 0 var(--btn-edge),
			0 12px 12px rgba(0, 0, 0, 0.25);
		transition: 0.5s;
	}

	.refresh-btn:hover {
		color: black;
		transition: 0.5s;
	}

	.refresh-btn:active {
		transform: translateY(4px);
		box-shadow:
			inset 0 1px 0 rgba(255, 255, 255, 0.25),
			0 2px 0 var(--btn-edge),
			0 4px 8px rgba(0, 0, 0, 0.2);
	}

	.refresh-btn:focus-visible {
		outline: 3px solid rgba(255, 0, 0, 0.7);
		outline-offset: 2px;
	}

	@media (max-width: 480px) {
		.price-card {
			min-height: auto;
		}
		.refresh-btn {
			width: 100%;
		}
	}
</style>
