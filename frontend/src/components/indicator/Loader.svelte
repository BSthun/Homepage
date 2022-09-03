<script lang="ts">
	import { getContext } from 'svelte';
	import CircularProgress from '../indicator/CircularProgress.svelte';
	
	const bind = getContext('bind');
	let load: boolean | string = false;
	let loadDesc: string = '';
	
	$: bind.update((value) => ({
		...value,
		setLoading: (m, n = null) => {
			load = m;
			loadDesc = n;
		},
	}));
</script>

{#if (load !== null && load !== false)}
	<div class="loader">
		{#if (load === true)}
			<CircularProgress />
		{/if}
		{#if (typeof load === "string")}
			<i class="las la-exclamation-circle" style="font-size: 36px"></i>
			<h3>{load}</h3>
			<p class="desc">{loadDesc}</p>
		{/if}
	</div>
{/if}

<style lang="scss">
	@import '../../styles/index';
	
	.loader {
		position: fixed;
		inset: 0;
		background-color: $color-grey-900;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		gap: 12px;
		z-index: 1000;
	}
	
	.desc {
		text-align: center;
		padding: 0 15px;
	}
</style>