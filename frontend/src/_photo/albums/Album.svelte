<script lang="ts">
	import { getContext, onMount } from 'svelte'
	import { Link, navigate, useLocation, useParams } from 'svelte-navigator'
	import type { Writable } from 'svelte/store'
	import Container from '../../components/layout/Container.svelte'
	import { axios, caller } from '../../utils/api'
	import SectionItem from './SectionItem.svelte'

	const params = useParams()
	const location = useLocation()
	const bind: Writable<any> = getContext('bind')

	let query = new URLSearchParams($location.search)

	let state: any = {
		album: {
			name: 'Album',
		},
	}

	const sectionNav = (id) => {
		navigate('/photo/section/' + id)
	}

	const mount = () => {
		window.scrollTo({ top: 0, behavior: 'smooth' })
		$bind.setLoading(true)
		caller(
			axios.get(`/photo/entity/album/detail`, {
				params: {
					album_slug: $params['album-slug'],
					token: query.get('token'),
				},
			})
		)
			.then((res) => {
				state = res.data
				navigate('/photo/album/' + state.album.slug, { replace: true })
				$bind.setLoading(false)
			})
			.catch((err) => {
				$bind.setLoading(err.message)
			})
	}

	onMount(mount)
</script>

<svelte:head>
	<title>{state.album.name} - BSthun</title>
</svelte:head>

<div class="album">
	<Container>
		<Link to="/photo">
			<div class="r-1 bordered">
				<i class="las la-arrow-left" />
				<h4>All albums</h4>
			</div>
		</Link>
		<h1 class="title">{state.album.name}</h1>
		<div class="grid">
			{#each state.album.sections || [] as item, i}
				<SectionItem {item} nav={sectionNav} />
			{/each}
		</div>
	</Container>
</div>

<style lang="scss">
	@import '../../styles/index';

	.album {
		@include default-paging;
	}

	.title {
		border-bottom: white 1px solid;
		padding-bottom: 16px;
	}

	.grid {
		margin-top: 24px;
		display: grid;
		gap: 12px;
		grid-template-columns: repeat(auto-fill, minmax(360px, 1fr));
	}

	// Duplicated from SectionMeta.svelte
	.bordered {
		width: fit-content;
		line-height: 1.3;
		padding: 6px 12px;
		margin: 0 0 12px 0;
		border-radius: 16px;
		border: 1px solid white;
		transition: background-color 0.2s ease-in-out;

		&:hover {
			background-color: rgba(255, 255, 255, 0.1);
		}
	}
</style>
