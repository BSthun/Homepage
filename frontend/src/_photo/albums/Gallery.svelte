<script lang="ts">
	import { Link } from 'svelte-navigator'
	import CircularProgress from '../../components/indicator/CircularProgress.svelte'
	import { getContext, onMount, setContext } from 'svelte'
	import { writable, type Writable } from 'svelte/store'
	import InfiniteScroll from '../../components/extension/InfiniteScroll.svelte'
	import { axios, caller } from '../../utils/api'
	import GalleryDetail from './GalleryDetail.svelte'
	import GalleryItem from './GalleryItem.svelte'

	export let state: any

	const bind: Writable<any> = getContext('bind')

	const gallery = writable<any>({
		count: state.section.id,
		page: 0,
		items: [],
		expand: null,
		fetching: false,
		fetch: () => {
			if ($gallery.fetching == true) return
			$gallery.fetching = true
			caller<any>(
				axios.get(`/photo/entity/photo/list`, {
					params: {
						section_id: state.section.id,
						page_no: $gallery.page,
					},
				})
			)
				.then((res) => {
					$gallery.page++
					$gallery.items = $gallery.items.concat(res.data.items)
				})
				.catch((err) => {
					$bind.openSnackbar(err.message)
				})
				.finally(() => {
					$gallery.fetching = false
				})
		},
	})
	setContext('gallery', gallery)

	const mount = () => {
		window.scrollTo({ top: 0, behavior: 'smooth' })
		$gallery.fetch()
	}

	onMount(mount)
</script>

<div class="gallery">
	<div class="gallery-view">
		{#each $gallery.items as item, i}
			<GalleryItem {item} index={i} />
		{/each}
	</div>
	<div style="display: flex; justify-content: center; margin: 24px 0">
		{#if $gallery.fetching === true}
			<CircularProgress />
		{/if}
		{#if $gallery.items.length === state.section.photo_count}
			<h5 style="text-align: center; line-height: 2">
				This is the end of the section.<br />
				<Link to={'/photo/album/' + state.album.slug} style="text-decoration: underline">
					Back to parent album ({state.album.name})
				</Link><br />
				<Link to="/photo" style="text-decoration: underline">Browse for photo albums</Link>
			</h5>
		{/if}
		{#if $gallery.items.length !== state.section.photo_count && $gallery.fetching === false}
			<button on:click={$gallery.fetch} class="more">Load more</button>
		{/if}
	</div>
	<InfiniteScroll
		more={$gallery.items.length !== state.section.photo_count}
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

	.more {
		cursor: pointer;
		border-radius: 12px;
		padding: 8px 16px;
		border: none;
	}
</style>
