<script>
	import { getContext, onMount } from 'svelte'

	const bind = getContext('bind')
	let count = 0

	function handleClick() {
		count += 1
	}

	onMount(() => {
		$bind.setLoading(false)
	})
</script>

<div class="home">
	<button on:click={handleClick}>
		Clicked {count}
		{count === 1 ? 'time' : 'times'}
	</button>
</div>

<svelte:head>
	<title>Home - BSthun</title>
</svelte:head>

<style lang="scss">
	@import 'src/styles/_index.scss';

	.home {
		@include default-paging;

		h1 {
			padding: 2px;
		}
	}
</style>
