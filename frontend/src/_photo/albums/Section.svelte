<script lang="ts">
	import { getContext, onMount } from 'svelte';
	import { useLocation, useParams } from 'svelte-navigator';
	import Container from '../../components/layout/Container.svelte';
	import { axios, caller } from '../../utils/api';
	import SectionGallery from './Gallery.svelte';
	import SectionMeta from './SectionMeta.svelte';
	
	const params = useParams();
	const location = useLocation();
	const bind = getContext('bind');
	
	let query = new URLSearchParams($location.search);
	
	let state: any = {
		section: {
			title: 'Album Section',
		},
		album: {},
	};
	
	const mount = () => {
		$bind.setLoading(true);
		caller(axios.get(`/photo/entity/section/detail`, {
			params: {
				section_id: $params['section-id'],
				token: query.get('token'),
			},
		}))
			.then((res) => {
				state = res.data;
				$bind.setLoading(false);
			})
			.catch((err) => {
				$bind.setLoading(err.message);
			});
	};
	
	onMount(mount);
</script>

<svelte:head>
	<title>{state.section.title} - BSthun</title>
</svelte:head>

<div class="section">
	<Container>
		<div class="title">
			<SectionMeta item={state} />
		</div>
		{#if (state.section.id !== undefined)}
			<SectionGallery id="{state.section.id}" count={state.section.photo_count} />
		{/if}
	</Container>
</div>

<style lang="scss">
	@import '../../styles/index';
	
	.section {
		@include default-paging;
	}
	
	.title {
		border-bottom: white 1px solid;
		padding-bottom: 16px;
		line-height: 1.6;
	}
</style>