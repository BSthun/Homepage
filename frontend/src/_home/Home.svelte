<script>
	let count = 0;
	
	function handleClick() {
		count += 1;
	}
</script>

<main>
	<div>
		<h1>Hello</h1>
		<button on:click={handleClick}>
			Clicked {count} {count === 1 ? 'time' : 'times'}
		</button>
	</div>
</main>

<style lang="scss">
	div {
		padding: 0;
		
		h1 {
			padding: 2px;
		}
	}
</style>