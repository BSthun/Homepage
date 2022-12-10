<script lang="ts">
	import { onDestroy, onMount } from 'svelte'
	import { navigate } from 'svelte-navigator'

	export let album: Object
</script>

<div class="album-item" on:click={() => navigate('/photo/album/' + album.slug)}>
	<div class="image" style={`background-image: url(${album.thumbnail_url})`} />
	<div class="info">
		<h1>{album.name}</h1>
		<div class="r-1">
			<i class="las la-folder-open" />
			<h4>{album.section_count} sections</h4>
		</div>
		<div class="r-1">
			<i class="las la-images" />
			<h4>{album.photo_count} photos</h4>
		</div>
	</div>
	<div class="overlay" />
	<div class="image-overlay" />
</div>

<style lang="scss">
	@import 'src/styles/_index.scss';

	.album-item {
		width: 100%;
		height: 420px;
		border-radius: 16px;
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

		> h1 {
			margin: 0 0 6px 0;
		}
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
