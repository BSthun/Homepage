<script lang="ts">
	import moment from 'moment';
	import { trackLog } from '../../utils/api/track';
	import { getContext } from 'svelte';
	import Zoom from 'svelte-zoom';
	import CircularProgress from '../../components/indicator/CircularProgress.svelte';
	
	const gallery = getContext('gallery');
	let loading = true;
	let plane = false;
	let image = new Image();
	image.onload = () => {
		loading = false;
	};
	
	const forward = (i) => {
		if (i < 0 && $gallery.expand != 0) {
			$gallery.expand -= 1;
		}
		if (i > 0 && $gallery.expand != $gallery.items.length - 1) {
			$gallery.expand += 1;
		}
		if (i > 0 && $gallery.expand == $gallery.items.length - 1 && $gallery.expand < $gallery.count) {
			$gallery.fetch();
		}
		trackLog('gallery/expand', null, $gallery.items[$gallery.expand].id);
	};
	
	const download = (id, url) => {
		trackLog('gallery/download', null, id);
		window.location.href = url + '?download=1';
	};
	
	$: item = $gallery.items[$gallery.expand];
	$: {
		if (item !== undefined) {
			loading = true;
			plane = false;
			image.src = item && (item.root + item.image_path);
		}
	}
	$: {
		if ($gallery.expand != null) {
			$gallery.loading = true;
			document.querySelector('body').style.overflow = 'hidden';
		} else {
			document.querySelector('body').style.overflow = 'scroll';
		}
	}
</script>

{#if ($gallery.expand != null)}
	<div class="gallery-detail">
		<div class="img-wrapper">
			{#if (loading === true)}
				<div class="center">
					<CircularProgress />
				</div>
			{:else }
				<Zoom alt={item.title} src={item.root + item.image_path} />
			{/if}
			<span class="material-symbols-outlined back" on:click={() => forward(-1)}>arrow_back</span>
			<span class="material-symbols-outlined forward" on:click={() => forward(1)}>arrow_forward</span>
		</div>
		<div class={plane ? "plane plane-toggled" : "plane"}>
			<div class="plane-close">
				<div
					class="plane-detail-toggle"
					on:click={() => { plane = !plane; trackLog('gallery/plane', null, plane); }}
				>
					<p>{plane ? "COLLAPSE" : "VIEW DETAIL / DOWNLOAD"}</p>
				</div>
				<span
					class="material-symbols-outlined"
					on:click={() => {$gallery.expand = null; trackLog('gallery/close', null, null);	}}
				>
					close
				</span>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">schedule</span>
				<div>
					<p>Timestamp</p>
					<p>{moment(item.exif.timestamp).format("MMM D YYYY, h:mm a")}</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">camera</span>
				<div>
					<p>Apearture</p>
					<p>ƒ/{item.exif.aperture}</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">shutter_speed</span>
				<div>
					<p>Shutter Speed</p>
					<p>{item.exif.exposure_time}s</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">crop</span>
				<div>
					<p>Focal Length</p>
					<p>{item.exif.focal_length}mm ({item.exif.focal_length_35mm}mm FF)</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">deblur</span>
				<div>
					<p>ISO Speed</p>
					<p>{item.exif.iso_speed}</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">light_mode</span>
				<div>
					<p>Brightness (RAW)</p>
					<p>{item.exif.brightness}</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">photo_camera</span>
				<div>
					<p>Camera</p>
					<p>{item.exif.camera_maker} {item.exif.camera_model}</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">lens</span>
				<div>
					<p>Lens</p>
					<p>{item.exif.lens_model}</p>
				</div>
			</div>
			<div class="plane-download" on:click={() => download(item.id, item.root + item.image_path)}>
				<span class="material-symbols-outlined">image</span>
				<p>Download Image</p>
			</div>
			<div class="plane-download" on:click={() => download(item.id, item.root + item.raw_path)}>
				<span class="material-symbols-outlined">add_photo_alternate</span>
				<p>Download Raw</p>
			</div>
		</div>
	</div>
{/if}

<style lang="scss">
	@import 'src/styles/_index.scss';
	
	.gallery-detail {
		position: fixed;
		inset: 0;
		z-index: 10000;
		background-color: #1E2118EA;
		display: flex;
		
		@include breakpoint('md', 'dn') {
			flex-direction: column;
		}
	}
	
	.img-wrapper {
		position: relative;
		flex: 1;
		overflow: hidden;
		
		@include breakpoint('md', 'up') {
			border-radius: 8px;
			margin: 12px 0 12px 12px;
		}
		
		@include breakpoint('md', 'dn') {
			margin-bottom: 48px;
		}
		
		.material-symbols-outlined {
			top: calc(50% - 12px);
			position: absolute;
			z-index: 200;
			color: $color-grey-800;
			background-color: $color-grey-100;
			@include tweak-clickable-icon($color-grey-300);
		}
		
		.back {
			left: 4px;
		}
		
		.forward {
			right: 4px;
		}
	}
	
	.plane {
		display: flex;
		flex-direction: column;
		background-color: white;
		color: $color-grey-900;
		transition: top 300ms ease-in-out;
		
		@include breakpoint('md', 'up') {
			width: 300px;
			border-radius: 8px;
			margin: 12px;
		}
		
		@include breakpoint('md', 'dn') {
			position: absolute;
			z-index: 201;
			top: calc(100% - 48px);
			width: 100%;
			height: 100%;
		}
	}
	
	.plane-toggled {
		top: 0;
	}
	
	.plane-close {
		padding: 8px;
		display: flex;
		justify-content: space-between;
		align-items: center;
		
		> span {
			@include tweak-clickable-icon($color-grey-300);
		}
	}
	
	.plane-detail-toggle {
		border-radius: 4px;
		border: #546E7A 1px solid;
		padding: 4px 8px;
		
		p {
			font-size: 12px;
		}
		
		@include breakpoint('md', 'up') {
			visibility: hidden;
		}
	}
	
	.plane-info {
		display: flex;
		align-items: center;
		padding: 12px;
		gap: 12px;
		
		> div {
			flex: 1;
			
			p:first-child {
				font-size: 12px;
				color: $color-bluegrey-700;
			}
			
			p:last-child {
				font-size: 16px;
				color: $color-grey-900;
			}
		}
	}
	
	.plane-download {
		margin: 12px 12px 0 12px;
		padding: 6px 12px;
		border: #546E7A 1px solid;
		border-radius: 4px;
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 6px;
		font-size: 16px;
		@include tweak-clickable;
	}
	
	.center {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100%;
	}
</style>