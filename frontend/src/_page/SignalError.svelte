<script lang="ts">
	import { axios, caller } from '../utils/api'
	import { getContext, onMount } from 'svelte'
	import { navigate, useLocation } from 'svelte-navigator'
	import Container from '../components/layout/Container.svelte'
	import NotFound from './NotFound.svelte'

	const location = useLocation()
	const bind = getContext('bind')

	let query = new URLSearchParams($location.search)

	onMount(() => {
		caller(
			axios.put(`/track/sig/err`, {
				title: query.get('e'),
			})
		)
			.then((res) => {
				navigate('/page/signal/err', { replace: true })
			})
			.finally(() => {
				$bind.setLoading(false)
			})
	})
</script>

<svelte:head>
	<title>Error - BSthun</title>
</svelte:head>

<NotFound
	title="An error occurred"
	description="Some system function redirected with error attachment named '{query.get('e')}'"
	code="ERR_REDIRECT_SIGNAL"
/>

<style lang="scss">
	@import 'src/styles/_index.scss';
</style>
