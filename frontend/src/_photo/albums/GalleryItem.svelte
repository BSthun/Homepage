<script lang="ts">
	import moment from 'moment';
	import { trackLog } from '../../utils/api/track';
	import { getContext } from 'svelte';
	
	export let item: Object;
	export let index: number;
	
	const gallery = getContext('gallery');
	
	const onClick = () => {
		trackLog('gallery/expand', null, item.id);
		gallery.update((value) => ({
			...value,
			expand: index,
		}));
	};
</script>

<div class="gallery-item" on:click={onClick}>
	<img alt={item.title} class="img" src={item.root + item.thumbnail_path} />
	<div class="overlay">
		<h4>{item.title}</h4>
		<div class="r-1">
			<h5>{moment(item.exif.timestamp).calendar()}</h5>
		</div>
		<div class="r-2" style={item.ratio < 1 && "display: none;"}>
			<div class="r-1">
				<span class="material-symbols-outlined">camera</span>
				<h5>ƒ/{item.exif.aperture}</h5>
			</div>
			<div class="r-1">
				<span class="material-symbols-outlined">shutter_speed</span>
				<h5>{item.exif.exposure_time}s</h5>
			</div>
			<div class="r-1">
				<span class="material-symbols-outlined">crop</span>
				<h5>{item.exif.focal_length}mm</h5>
			</div>
			<div class="r-1">
				<span class="material-symbols-outlined">deblur</span>
				<h5>ISO {item.exif.iso_speed}</h5>
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
			img {
				filter: blur(4px) brightness(0.8);
			}
			
			.overlay {
				bottom: 25px;
			}
		}
	}
	
	.img {
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
		bottom: -64px;
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
		
		@include breakpoint('md', 'dn') {
			display: none;
		}
	}
</style>