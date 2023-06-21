<script lang="ts">
	import moment from 'moment'
	import { getContext } from 'svelte'
	import { BlurhashImage } from 'svelte-blurhash'
	import { trackLog } from '../../utils/api/track'

	export let item: any
	export let index: number

	let windowWidth: number
	let divWidth: number

	const gallery = getContext('gallery')

	const onClick = () => {
		trackLog('gallery/expand', null, item.id)
		gallery.update((value) => ({
			...value,
			expand: index,
		}))
	}

	$: calculated = (() => {
		if (item.exif.w === 0) {
			return null
		}
		if (windowWidth < 768) {
			return {
				width: '100%',
				height: (divWidth / item.exif.w) * item.exif.h,
			}
		}

		const maxWidth = 1024 / 2
		const base = {
			width: (400 / item.exif.h) * item.exif.w,
			height: 400,
		}
		if (base.width > maxWidth) {
			base.height = (maxWidth / base.width) * base.height
			base.width = maxWidth
		}

		return base
	})()
</script>

<svelte:window bind:innerWidth={windowWidth} />

<div
	class="gallery-item"
	on:click={onClick}
	bind:offsetWidth={divWidth}
	style={item.exif.w !== 0 ? 'height: fit-content' : ''}
>
	<!-- TODO: Migrate all image to have width and height exif properties -->
	{#if item.exif.w !== 0}
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
			.img,
			.blurhash {
				filter: blur(4px) brightness(0.8);
			}

			.overlay {
				bottom: 25px;
			}
		}
	}

	.img,
	.blurhash {
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
