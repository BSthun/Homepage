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
	<div class="gallery-view">
		{#each items as item, i}
			<GalleryItem item={item} />
		{/each}
	</div>
	<InfiniteScroll
		more={items.length !== count}
		on:load={() => {page++; fetch()}}
		threshold={700 > window.innerHeight ? 800 : window.innerHeight - 100}
	/>
</div>

<style lang="scss">
	.gallery-view {
		margin-top: 16px;
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		align-items: center;
		justify-content: center;
		gap: 6px;
	}
</style>