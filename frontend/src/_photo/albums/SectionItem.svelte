<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { navigate } from 'svelte-navigator';
	import SectionMeta from './SectionMeta.svelte';
	
	export let item: object;
	
	const thumbnails = item.thumbnail_url.split(',');
	let thumbnail = 0;
	let interval: NodeJS.Timer;
	
	$: console.log(thumbnails);
	
	const onClick = () => {
		navigate('/photo/section/' + item.id);
	};
	
	onMount(() => {
		interval = setInterval(() => {
			thumbnail = (thumbnail + 1) % thumbnails.length;
		}, 1000);
	});
	
	onDestroy(() => {
		clearInterval(interval);
	});
</script>

<div class="section-item" on:click={() => onClick()}>
	<div
		class="image"
		style={`background-image: url(${thumbnails[thumbnail]})`}
	></div>
	<div class="info">
		<SectionMeta item={{ section: item}} />
	</div>
	<div class="overlay"></div>
	<div class="image-overlay"></div>
</div>

<style lang="scss">
	.section-item {
		width: 100%;
		height: 420px;
		border-radius: 20px;
		overflow: hidden;
		position: relative;
		cursor: pointer;
		
		&:hover {
			.image-overlay {
				backdrop-filter: blur(4px) brightness(0.9);
			}
		}
	}
	
	.image {
		position: absolute;
		inset: 0;
		z-index: 0;
		background-size: cover;
		background-position: center;
	}
	
	.info {
		position: absolute;
		bottom: 0;
		z-index: 10;
		padding: 24px;
	}
	
	.overlay {
		position: absolute;
		background: linear-gradient(0deg, rgba(0, 0, 0, 0.7) 0%, rgba(0, 0, 0, 0) 100%);
		bottom: 0;
		left: 0;
		right: 0;
		height: 200px;
		z-index: 5;
	}
	
	.image-overlay {
		position: absolute;
		inset: 0;
		z-index: 1;
		transition: backdrop-filter 300ms ease-in-out;
	}
</style>