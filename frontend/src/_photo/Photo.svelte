<script lang="ts">
	import { axios, caller } from '../utils/api'
	import { getContext, onMount } from 'svelte'
	import { navigate } from 'svelte-navigator'
	import type { Writable } from 'svelte/store'
	import Container from '../components/layout/Container.svelte'

	const bind: Writable<any> = getContext('bind')

	let state: any = {
		album: {
			name: 'Album',
		},
	}

	const mount = () => {
		$bind.setLoading(true)
		caller(axios.get(`/photo/entity/album/list`))
			.then((res) => {
				state = res.data
				$bind.setLoading(false)
			})
			.catch((err) => {
				$bind.setLoading(err.message)
			})
	}

	onMount(mount)
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
