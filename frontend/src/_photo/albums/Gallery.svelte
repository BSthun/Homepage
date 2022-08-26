<script lang="ts">
	import { trackLog } from '../../utils/api/track';
	import { getContext, onMount, setContext } from 'svelte';
	import { writable } from 'svelte/store';
	import InfiniteScroll from '../../components/extension/InfiniteScroll.svelte';
	import { axios, caller } from '../../utils/api';
	import GalleryDetail from './GalleryDetail.svelte';
	import GalleryItem from './GalleryItem.svelte';
	
	export let id: number;
	export let count: number;
	
	const bind = getContext('bind');
	
	const gallery = writable<any>({
		count: count,
		page: 0,
		items: [],
		expand: null,
		fetching: false,
		fetch: () => {
			if ($gallery.fetching == true) return;
			$gallery.fetching = true;
			caller(axios.get(`/photo/entity/photo/list`, {
				params: {
					section_id: id,
					page_no: $gallery.page,
				},
			}))
				.then((res) => {
					$gallery.page++;
					$gallery.items = $gallery.items.concat(res.data.items);
				})
				.catch((err) => {
					$bind.openSnackbar(err.message);
				})
				.finally(() => {
					$gallery.fetching = false;
				});
		},
	});
	setContext('gallery', gallery);
	
	onMount($gallery.fetch);
</script>

<div class="gallery">
	<div class="gallery-view">
		{#each $gallery.items as item, i}
			<GalleryItem item={item} index="{i}" />
		{/each}
	</div>
	<InfiniteScroll
		more={$gallery.items.length !== count}
		on:load={$gallery.fetch()}
		threshold={700 > window.innerHeight ? 800 : window.innerHeight - 100}
	/>
	<GalleryDetail />
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