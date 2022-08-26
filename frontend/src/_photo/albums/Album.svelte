<script lang="ts">
	import { getContext, onMount } from 'svelte';
	import { useLocation, useParams } from 'svelte-navigator';
	import Container from '../../components/layout/Container.svelte';
	import { axios, caller } from '../../utils/api';
	import SectionItem from './SectionItem.svelte';
	
	const params = useParams();
	const location = useLocation();
	const bind = getContext('bind');
	
	let query = new URLSearchParams($location.search);
	
	let state: any = {
		album: {
			name: 'Album',
		},
	};
	
	const mount = () => {
		$bind.setLoading(true);
		caller(axios.get(`/photo/entity/album/detail`, {
			params: {
				album_slug: $params['album-slug'],
				token: query.get('token'),
			},
		}))
			.then((res) => {
				state = res.data;
				$bind.setLoading(false);
			})
			.catch((err) => {
				$bind.setLoading(err.message);
			})
	};
	
	onMount(mount);
</script>

<svelte:head>
	<title>{state.album.name} - BSthun</title>
</svelte:head>

<div class="album">
	<Container>
		<h1 class="title">Photo</h1>
		<div class="grid">
			{#each state.album.sections || [] as item, i}
				<SectionItem item={item} />
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
		
		grid-template-columns: repeat(2, 50% [col-start]);
		
		@include breakpoint('md', 'dn') {
			grid-template-columns: repeat(1, 100% [col-start]);
		}
	}
</style>