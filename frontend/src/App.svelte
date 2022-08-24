<script lang="ts">
	import { Route, Router } from 'svelte-navigator';
	import Home from './_home/Home.svelte';
	import Photo from './_photo/Photo.svelte';
	import NotFound from './_splash/NotFound.svelte';
	import Navbar from './components/navbar/Navbar.svelte';
</script>

<svelte:head>
	<title>BSthun</title>
</svelte:head>

<div>
	<Router>
		<Navbar />
		<Route path="/">
			<Home />
		</Route>
		<Route path="/photo">
			<Photo />
		</Route>
		<Route>
			<NotFound />
		</Route>
	</Router>
</div>

<style global lang="scss">
	// Base StyleSheets
	@import 'styles/core/base';
	@import 'styles/core/font';
	@import 'styles/core/icon';
</style>