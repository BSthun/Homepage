<script lang="ts">
	import { axios, caller } from '../utils/api'
	import { getContext, onMount } from 'svelte'
	import type { Writable } from 'svelte/store'
	import Container from '../components/layout/Container.svelte'
	import Graph from './components/Graph.svelte'
	import ViewLocker from './components/ViewLocker.svelte'

	const bind: Writable<any> = getContext('bind')

	let local: any = {
		year: new Date().getFullYear().toString(),
		yearOptions: Array.from(
			{
				length: new Date(new Date().getFullYear(), 0, 1).getFullYear() - 2022 + 1,
			},
			(_, i) => (i + 2022).toString()
		),
	}

	let remote: any = {
		days_graph: [],
	}

	const fetch = () => {
		caller(
			axios.get(`/diary/graph`, {
				params: {
					year: local.year,
				},
			})
		)
			.then((res) => {
				remote = res.data
				$bind.setLoading(false)
			})
			.catch((err) => {
				$bind.setLoading(err.message)
			})
	}

	const mount = () => {
		fetch()
	}

	$: {
		if (local.year) {
			fetch()
			const firstDay = new Date(local.year, 0, 1)
			local = {
				...local,
				year: firstDay.getFullYear().toString(),
				firstDayOffset: firstDay.getDay(),
			}
		}
	}

	onMount(mount)
</script>

<svelte:head>
	<title>Diary - BSthun</title>
</svelte:head>

<div class="diary">
	<Container>
		<h1 class="title">Diary</h1>

		<div class="lister">
			<Graph state={{ local, remote }} setYear={(year) => (local.year = year)} />
			<ViewLocker />
		</div>
	</Container>
</div>

<style lang="scss">
	@import '../styles/index';

	.diary {
		@include default-paging;
	}

	.title {
		border-bottom: white 1px solid;
		padding-bottom: 16px;
	}

	.lister {
		display: flex;
		flex-direction: column;
		align-items: center;
	}
</style>
