<script lang="ts">
	import { getContext, onMount } from 'svelte'
	import type { Writable } from 'svelte/store'
	import Container from '../components/layout/Container.svelte'

	const bind: Writable<any> = getContext('bind')

	onMount(() => {
		$bind.setLoading(false)
	})
</script>

<svelte:head>
	<title>Photo - BSthun</title>
</svelte:head>

<div class="album">
	<Container>
		<h1 class="title">Photo</h1>
	</Container>
</div>

<style lang="scss">
	@import 'src/styles/_index.scss';

	.album {
		min-height: 100vh;
		background-color: $color-grey-900;
		padding-top: $size-navbar-height;
	}

	.title {
		border-bottom: white 1px solid;
		padding-bottom: 16px;
	}
</style>
