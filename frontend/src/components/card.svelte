<script>
	import Icon from '@iconify/svelte';
</script>

<div class="card-container">
	<h1 class="card-heading">XMR Price Tracker</h1>
	<div class="price-card">
		<div class="price-wrapper">
			<Icon icon="logos:monero" width="100" height="100" />
			<p class="price">200.43€</p>
		</div>
		<p class="change">XMR/EUR</p>
		<p class="last-updated">Last updated: Dec 31, 2026, 10:20</p>
		<button class="refresh-btn">Refresh Price</button>
	</div>
</div>

<style>
	.card-container {
		display: flex;
		align-items: center;
		flex-direction: column;
		gap: 10rem;
	}
	.card-heading {
		color: var(--orange);
		margin-top: 5rem;
		font-size: 3rem;
		text-shadow: 1px 10px 20px var(--orange);
	}
	.price-card {
		display: flex;
		flex-direction: column;
		/* justify-content: center; */
		align-items: center;
		width: 40rem;
		height: 25rem;
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
		gap: 1rem;
		padding-top: 1rem;
	}
	.price {
		font-size: 7rem;
	}
	.change {
		font-size: 2rem;
	}
	.last-updated {
		font-size: 1.5rem;
	}
	.refresh-btn {
		padding: 1rem 12rem;
		color: var(--white);
		font-size: 1.35rem;
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
		outline: 3px solid rgba(255, 255, 255, 0.7);
		outline-offset: 2px;
	}
</style>
