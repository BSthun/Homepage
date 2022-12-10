<script lang="ts">
	import moment from 'moment'
	import { trackLog } from '../../utils/api/track'
	import { getContext, onDestroy, onMount } from 'svelte'
	import SectionMeta from './SectionMeta.svelte'

	export let item: Object
	export let nav: Function

	const thumbnails = item.thumbnail_url.split(',')
	let thumbnail = 0
	let interval: number

	onMount(() => {
		interval = setInterval(() => {
			thumbnail = (thumbnail + 1) % thumbnails.length
		}, 1000)
	})

	onDestroy(() => {
		clearInterval(interval)
	})
</script>

<div class="album-item" on:click={() => nav(item.id)}>
	<div class="cover" style={`background-image: url('${thumbnails[thumbnail]}')`} />
	<div class="info">
		<SectionMeta item={{ section: item }} />
	</div>
</div>

<style lang="scss">
	@import '../../styles/index';

	.album-item {
		display: flex;
		flex-direction: column;
		cursor: pointer;
	}

	.cover {
		width: 100%;
		padding-top: 100%;
		border-radius: 16px;
		background-size: cover;
		background-position: center;
	}

	.info {
		margin: 12px 0 0 0;
	}
</style>
