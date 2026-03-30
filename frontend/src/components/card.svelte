<div class="card-container">
	<h1 class="card-heading">XMR Price Tracker</h1>
	<div class="price-card"></div>
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
        text-shadow: 1px 1px 4px var(--orange);
	}
	.price-card {
		width: 40rem;
		height: 25rem;
		box-shadow:
			0px -3px 0px var(--orange),
			0px 3px 0px var(--orange),
               rgba(0, 0, 0, 0.56) 0px 22px 70px 4px;
        border-radius: 12px;
        background-color: var(--light-gray);
	}
</style>
