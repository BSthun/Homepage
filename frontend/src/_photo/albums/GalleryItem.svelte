<script lang="ts">
	import moment from 'moment'
	import { getContext } from 'svelte'
	import { BlurhashImage } from 'svelte-blurhash'
	import { trackLog } from '../../utils/api/track'

	export let item: any
	export let index: number

	const gallery = getContext('gallery')
	const geometry = getContext('geometry')

	const onClick = () => {
		trackLog('gallery/expand', null, item.id)
		gallery.update((value) => ({
			...value,
			expand: index,
		}))
	}

	const onHover = (event: 'me' | 'ml' | 'ts' | 'te') => {
		trackLog('gallery/item/' + event, null, item.id)
	}

	$: calculated = item.exif.w == 0 ? null : $geometry.calculate(item.exif.w, item.exif.h)
</script>

<div
	class="gallery-item"
	on:click={onClick}
	on:mouseenter={() => onHover('me')}
	style={calculated !== null ? 'height: fit-content' : ''}
>
	<!-- TODO: Migrate all image to have width and height exif properties -->
	{#if calculated !== null}
		<div class="blurhash">
			<BlurhashImage
				hash={item.blurhash}
				src={item.root + item.thumbnail_path}
				width={calculated.width}
				height={calculated.height}
				fadeDuration={600}
			/>
		</div>
	{:else}
		<img class="img" alt={item.title} src={item.root + item.thumbnail_path} />
	{/if}

	<div class="overlay">
		<h4>{item.title}</h4>
		<div class="r-1">
			<h5>{moment(item.exif.t).calendar()}</h5>
		</div>
		<div class="r-2">
			<div class="r-1">
				<span class="material-symbols-outlined">camera</span>
				<h5>ƒ/{item.exif.apt}</h5>
			</div>
			<div class="r-1">
				<span class="material-symbols-outlined">shutter_speed</span>
				<h5>{item.exif.ss}s</h5>
			</div>
			<div class="r-1">
				<span class="material-symbols-outlined">crop</span>
				<h5>{item.exif.fl}mm</h5>
			</div>
			<div class="r-1">
				<span class="material-symbols-outlined">deblur</span>
				<h5>ISO {item.exif.iso}</h5>
			</div>
		</div>
	</div>
</div>

<style lang="scss">
	@import '../../styles/index';

	@mixin itemActive {
		.img,
		.blurhash {
			filter: blur(4px) brightness(0.8);
		}

		.overlay {
			bottom: 25px;
		}
	}

	.gallery-item {
		border-radius: 4px;
		overflow: hidden;
		position: relative;
		cursor: pointer;

		@include breakpoint('md', 'up') {
			height: 400px;
		}

		@include breakpoint('md', 'dn') {
			width: 100%;
		}

		&:hover {
			@media (pointer: fine) {
				@include itemActive;
			}
		}

		&:active {
			@media (hover: none) {
				@include itemActive;
			}
		}
	}

	.img,
	.blurhash {
		-webkit-touch-callout: none;
		transition: filter 300ms ease-in-out;

		@include breakpoint('md', 'up') {
			height: 100%;
		}

		@include breakpoint('md', 'dn') {
			width: 100%;
		}
	}

	.overlay {
		position: absolute;
		bottom: -96px;
		left: 25px;
		transition: bottom 300ms ease-in-out;
	}

	.r-1 {
		display: flex;
		align-items: center;
		margin-top: 4px;
		margin-right: 16px;

		.material-symbols-outlined {
			padding-right: 4px;
			font-size: 18px;
		}
	}

	.r-2 {
		display: flex;
		flex-wrap: wrap;
		white-space: nowrap;

		@include breakpoint('md', 'dn') {
			display: none;
		}
	}
</style>
