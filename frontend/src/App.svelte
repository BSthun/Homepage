<script lang="ts">
	import { onMount, setContext } from 'svelte';
	import { Route, Router } from 'svelte-navigator';
	import { axios, caller } from '../utils/api';
	import Home from './_home/Home.svelte';
	import NotFound from './_page/NotFound.svelte';
	import Album from './_photo/Album.svelte';
	import Photo from './_photo/Photo.svelte';
	import CircularProgress from './components/indicator/CircularProgress.svelte';
	import Navbar from './components/navbar/Navbar.svelte';
	
	let state = null;
	
	onMount(() => {
		caller(axios.get("/account/state")).then((res) => {
			state = res.data
			setContext("state", res.data)
		});
	});
</script>

<svelte:head>
	<title>BSthun</title>
</svelte:head>

<div>
	{#if (state !== null)}
		<Router>
			<Navbar />
			<Route path="/">
				<Home />
			</Route>
			<Route path="/photo">
				<Photo />
			</Route>
			<Route path="/photo/album">
				<Album />
			</Route>
			<Route>
				<NotFound />
			</Route>
		</Router>
	{:else }
		<div class="loading">
			<CircularProgress />
		</div>
	{/if}
</div>

<style global lang="scss">
	// Base StyleSheets
	@import 'styles';
	@import 'node_modules/@smui/circular-progress/style';
	
	.loading {
		@include default-paging;
		display: flex;
		justify-content: center;
		align-items: center;
	}
</style>