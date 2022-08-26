<script lang="ts">
	import { getContext, onMount } from 'svelte';
	import InfiniteScroll from '../../components/extension/InfiniteScroll.svelte';
	import { axios, caller } from '../../utils/api';
	import GalleryItem from './GalleryItem.svelte';
	
	const bind = getContext('bind');
	
	export let id: number;
	export let count: number;
	
	let page: number = 0;
	let items: [] = [];
	
	const fetch = () => {
		caller(axios.get(`/photo/entity/photo/list`, {
			params: {
				section_id: id,
				page_no: page,
			},
		}))
			.then((res) => {
				items = items.concat(res.data.items);
			})
			.catch((err) => {
				$bind.openSnackbar(err.message);
			});
	};
	
	onMount(fetch);
</script>

<div class="gallery">
	{#each items as item, i}
		<GalleryItem item={item} />
	{/each}
	<InfiniteScroll
		more={items.length !== count}
		on:load={() => {page++; fetch()}}
		threshold={200}
	/>
</div>