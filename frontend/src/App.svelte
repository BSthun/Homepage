<script lang="ts">
	import { Route, Router } from "svelte-navigator";
	import Home from './_home/Home.svelte';
	import NotFound from './_splash/NotFound.svelte';
	import Navbar from './components/navbar/Navbar.svelte';
</script>

<svelte:head>
	<title>BSthun</title>
</svelte:head>

<div>
	<Router>
		<Navbar />
		<Route meta="{{ name: 'blog' }}" path="/">
			<Home />
		</Route>
		<Route path="/about">
			<Home />
		</Route>
		<Route>
			<NotFound />
		</Route>
	</Router>
</div>

<style global lang="scss">
	// Base StyleSheets
	@import 'src/styles/base';
	@import 'src/styles/font';
	@import 'src/styles/icon';
</style>