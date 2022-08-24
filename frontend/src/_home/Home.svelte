<script>
	let count = 0;
	
	function handleClick() {
		count += 1;
	}
</script>

<div class="home">
	<button on:click={handleClick}>
		Clicked {count} {count === 1 ? 'time' : 'times'}
	</button>
</div>

<svelte:head>
	<title>Home - BSthun</title>
</svelte:head>

<style lang="scss">
	@import 'src/styles/_index.scss';
	
	.home {
		min-height: 100vh;
		background-color: $color-grey-900;
		padding-top: $size-navbar-height;
		
		h1 {
			padding: 2px;
		}
	}
</style>