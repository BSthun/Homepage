<script lang="ts">
	import { onMount, setContext } from 'svelte';
	import { writable } from 'svelte/store';
	import AppRouter from './AppRouter.svelte';
	import CircularProgress from './components/indicator/CircularProgress.svelte';
	import SnackBar from './components/indicator/SnackBar.svelte';
	import Loader from './components/layout/Loader.svelte';
	import { axios, caller } from './utils/api';
	
	let state = writable<any>(null);
	let bind = writable<any>({});
	setContext('state', state);
	setContext('bind', bind);
	
	onMount(() => {
		$bind.setLoading(true)
		caller(axios.get('/account/state'))
			.then((res) => {
				state.set(res.data);
			})
			.catch((err) => {
				$bind.openSnackbar(err.message)
			});
	});
</script>

<svelte:head>
	<title>BSthun</title>
</svelte:head>

<div>
	{#if ($state !== null)}
		<AppRouter />
	{:else}
		<div class="loading">
			<CircularProgress />
		</div>
	{/if}
	<SnackBar />
	<Loader />
</div>

<style global lang="scss">
	// Base StyleSheets
	@import 'styles';
	@import 'node_modules/@smui/circular-progress/style';
	@import 'node_modules/@smui/snackbar/style';
</style>