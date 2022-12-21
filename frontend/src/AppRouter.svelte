<script>
	import { Route, Router } from 'svelte-navigator'
	import Home from './_home/Home.svelte'
	import NotFound from './_page/NotFound.svelte'
	import SignalError from './_page/SignalError.svelte'
	import Album from './_photo/albums/Album.svelte'
	import Section from './_photo/albums/Section.svelte'
	import Photo from './_photo/Photo.svelte'
	import Navbar from './components/navbar/Navbar.svelte'
</script>

<Router>
	<Navbar />
	<Route path="/">
		<Home />
	</Route>
	<Route path="/photo">
		<Photo />
	</Route>
	<Route path="/photo/album/:album-slug">
		<Album />
	</Route>
	<Route path="/photo/section/:section-id">
		<Section />
	</Route>
	<Route path="/page/signal/err">
		<SignalError />
	</Route>
	<Route>
		<NotFound />
	</Route>
</Router>
